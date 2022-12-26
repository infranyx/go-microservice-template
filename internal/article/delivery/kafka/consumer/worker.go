package articleKafkaConsumer

import (
	"context"
	"encoding/json"

	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/infranyx/go-grpc-template/pkg/wrapper"
)

func (c *consumer) createEventWorker(
	workerChan chan bool,
) wrapper.HandlerFunc {
	return func(ctx context.Context, args ...interface{}) (interface{}, error) {
		// TODO : Graceful shutdown workers
		defer func() {
			workerChan <- true
		}()
		for {
			msg, err := c.createEventReader.Client.FetchMessage(ctx)
			if err != nil {
				return nil, err
			}

			logger.Zap.Sugar().Infof(
				"Kafka Worker recieve message at topic/partition/offset %v/%v/%v: %s = %s\n",
				msg.Topic,
				msg.Partition,
				msg.Offset,
				string(msg.Key),
				string(msg.Value),
			)

			aDto := new(articleDto.CreateArticleDto)
			if err := json.Unmarshal(msg.Value, &aDto); err != nil {
				continue
			}

			if err := c.createEventReader.Client.CommitMessages(ctx, msg); err != nil {
				return nil, err
			}
		}
	}
}

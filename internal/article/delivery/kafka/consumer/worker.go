package articleKafkaConsumer

import (
	"context"
	"encoding/json"
	"fmt"

	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/infranyx/go-grpc-template/pkg/wrapper"
)

func (ac *articleConsumer) createArticleWorker(
	ctx context.Context,
	c chan bool,
) wrapper.HandlerFunc {
	return func(tx context.Context, args ...interface{}) (interface{}, error) {
		defer func() {
			c <- true
		}()
		for {
			m, err := ac.createReader.Client.FetchMessage(ctx)
			if err != nil {
				return nil, err
			}

			logger.Zap.Sugar().Infof(
				"Kafka Worker recieve message at topic/partition/offset %v/%v/%v: %s = %s\n",
				m.Topic,
				m.Partition,
				m.Offset,
				string(m.Key),
				string(m.Value),
			)

			aDto := new(articleDto.CreateArticle)
			if err := json.Unmarshal(m.Value, &aDto); err != nil {
				continue
			}

			// run some usecase
			fmt.Println(aDto)

			if err := ac.createReader.Client.CommitMessages(ctx, m); err != nil {
				return nil, err
			}
		}
	}
}

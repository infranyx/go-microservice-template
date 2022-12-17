package articleKafkaConsumer

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
	errorUtils "github.com/infranyx/go-grpc-template/pkg/error/error_utils"
	"github.com/infranyx/go-grpc-template/pkg/logger"
)

func (ac *articleConsumer) createArticleWorker(
	ctx context.Context,
	wg *sync.WaitGroup,
	workerID int, // TODO: generate UUID
) {
	errorUtils.HandlerErrorWrapper(
		func() error {
			defer wg.Done()
			for {
				m, err := ac.createReader.Client.FetchMessage(ctx)
				if err != nil {
					continue
				}

				logger.Zap.Sugar().Infof(
					"Kafka Worker %v recieve message at topic/partition/offset %v/%v/%v: %s = %s\n",
					workerID,
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
					continue
				}
			}
		},
	)

}

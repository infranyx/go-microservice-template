package articleKafkaConsumer

import (
	"context"
	"sync"

	"github.com/infranyx/go-grpc-template/pkg/logger"
)

func (ac *articleConsumer) createArticleWorker(
	ctx context.Context,
	wg *sync.WaitGroup,
	workerID int, // TODO: generate UUID
) {
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

		// if err := json.Unmarshal(m.Value, &); err != nil {

		// 	continue
		// }

		// run some usecase

		if err := ac.createReader.Client.CommitMessages(ctx, m); err != nil {

			continue
		}
	}
}

package articleKafkaConsumer

import (
	"context"
	"sync"

	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	kafkaConsumer "github.com/infranyx/go-grpc-template/pkg/kafka/consumer"
	"github.com/infranyx/go-grpc-template/pkg/logger"
)

type articleConsumer struct {
	createReader *kafkaConsumer.Reader
}

func NewArticleConsumer(r *kafkaConsumer.Reader) articleDomain.ArticleConsumer {
	return &articleConsumer{createReader: r}
}

func (ac *articleConsumer) RunConsumers(ctx context.Context) {
	go ac.consumerCreateArticle(ctx, 2)
}

func (ac *articleConsumer) consumerCreateArticle(ctx context.Context, workersNum int) {
	r := ac.createReader.Client
	defer func() {
		if err := r.Close(); err != nil {
			logger.Zap.Sugar().Errorf("error closing create article consumer")
		}
	}()

	logger.Zap.Sugar().Infof("Starting consumer group: %v", r.Config().GroupID)

	wg := &sync.WaitGroup{}
	for i := 0; i <= workersNum; i++ {
		wg.Add(1)
		go ac.createArticleWorker(ctx, wg, i)
	}
	wg.Wait()
}

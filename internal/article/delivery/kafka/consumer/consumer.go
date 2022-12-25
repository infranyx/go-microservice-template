package articleKafkaConsumer

import (
	"context"

	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	kafkaConsumer "github.com/infranyx/go-grpc-template/pkg/kafka/consumer"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/infranyx/go-grpc-template/pkg/wrapper"
	wrapperErrorhandler "github.com/infranyx/go-grpc-template/pkg/wrapper/handlers/error_handler"
	wrapperRecoveryhandler "github.com/infranyx/go-grpc-template/pkg/wrapper/handlers/recovery_handler"
	wrapperSentryhandler "github.com/infranyx/go-grpc-template/pkg/wrapper/handlers/sentry_handler"
)

type consumer struct {
	createEventReader *kafkaConsumer.Reader
}

func NewConsumer(r *kafkaConsumer.Reader) articleDomain.ArticleConsumer {
	return &consumer{createEventReader: r}
}

func (c *consumer) RunConsumers(ctx context.Context) {
	go c.createEvent(ctx, 2)
}

func (c *consumer) createEvent(ctx context.Context, workersNum int) {
	r := c.createEventReader.Client
	defer func() {
		if err := r.Close(); err != nil {
			logger.Zap.Sugar().Errorf("error closing create article consumer")
		}
	}()

	logger.Zap.Sugar().Infof("Starting consumer group: %v", r.Config().GroupID)

	workerChan := make(chan bool)
	worker := wrapper.BuildChain(
		c.createEventWorker(workerChan),
		wrapperSentryhandler.SentryHandler,
		wrapperRecoveryhandler.RecoveryHandler,
		wrapperErrorhandler.ErrorHandler,
	)
	for i := 0; i <= workersNum; i++ {
		go worker(ctx, nil)
	}

	for {
		<-workerChan
		go worker(ctx, nil)
	}
}

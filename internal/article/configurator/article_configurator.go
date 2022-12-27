package articleConfigurator

import (
	"context"

	sampleExtServiceUseCase "github.com/infranyx/go-grpc-template/external/sample_ext_service/usecase"
	articleGrpcController "github.com/infranyx/go-grpc-template/internal/article/delivery/grpc"
	articleHttpController "github.com/infranyx/go-grpc-template/internal/article/delivery/http"
	articleKafkaConsumer "github.com/infranyx/go-grpc-template/internal/article/delivery/kafka/consumer"
	articleKafkaProducer "github.com/infranyx/go-grpc-template/internal/article/delivery/kafka/producer"
	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	articleJob "github.com/infranyx/go-grpc-template/internal/article/job"
	articleRepository "github.com/infranyx/go-grpc-template/internal/article/repository"
	articleUseCase "github.com/infranyx/go-grpc-template/internal/article/usecase"
	externalBridge "github.com/infranyx/go-grpc-template/pkg/external_bridge"
	infraContainer "github.com/infranyx/go-grpc-template/pkg/infra_container"
	articleV1 "github.com/infranyx/protobuf-template-go/golang-grpc-template/article/v1"
)

type configurator struct {
	ic *infraContainer.IContainer
	eb *externalBridge.ExternalBridge
}

func NewConfigurator(ic *infraContainer.IContainer, eb *externalBridge.ExternalBridge) articleDomain.Configurator {
	return &configurator{ic: ic, eb: eb}
}

func (c *configurator) Configure(ctx context.Context) error {
	seServiceUseCase := sampleExtServiceUseCase.NewSampleExtServiceUseCase(c.eb.SampleExtGrpcService)
	kafkaProducer := articleKafkaProducer.NewProducer(c.ic.KafkaWriter)
	repository := articleRepository.NewRepository(c.ic.Pg)
	useCase := articleUseCase.NewUseCase(repository, seServiceUseCase, kafkaProducer)

	// grpc
	grpcController := articleGrpcController.NewController(useCase)
	articleV1.RegisterArticleServiceServer(c.ic.GrpcServer.GetCurrentGrpcServer(), grpcController)

	// http
	httpRouterGp := c.ic.EchoServer.GetEchoInstance().Group(c.ic.EchoServer.GetBasePath())
	httpController := articleHttpController.NewController(useCase)
	articleHttpController.NewRouter(httpController).Register(httpRouterGp)

	// Consumers
	articleKafkaConsumer.NewConsumer(c.ic.KafkaReader).RunConsumers(ctx)

	// Jobs
	articleJob.NewJob(c.ic.Logger).StartJobs(ctx)

	return nil
}

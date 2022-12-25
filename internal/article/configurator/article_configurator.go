package articleConfigurator

import (
	"context"

	sampleExtServiceUseCase "github.com/infranyx/go-grpc-template/external/sample_ext_service/usecase"
	articleGrpc "github.com/infranyx/go-grpc-template/internal/article/delivery/grpc"
	articleHttp "github.com/infranyx/go-grpc-template/internal/article/delivery/http"
	articleKafkaConsumer "github.com/infranyx/go-grpc-template/internal/article/delivery/kafka/consumer"
	articleKafkaProducer "github.com/infranyx/go-grpc-template/internal/article/delivery/kafka/producer"
	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	articleJob "github.com/infranyx/go-grpc-template/internal/article/job"
	articleRepo "github.com/infranyx/go-grpc-template/internal/article/repository"
	articleUseCase "github.com/infranyx/go-grpc-template/internal/article/usecase"
	externalBridge "github.com/infranyx/go-grpc-template/pkg/external_bridge"
	infraContainer "github.com/infranyx/go-grpc-template/pkg/infra_container"

	articleV1 "github.com/infranyx/protobuf-template-go/golang-grpc-template/article/v1"
)

type articleConfigurator struct {
	ic *infraContainer.IContainer
	eb *externalBridge.ExternalBridge
}

func NewArticleConfigurator(ic *infraContainer.IContainer, eb *externalBridge.ExternalBridge) articleDomain.ArticleConfigurator {
	return &articleConfigurator{ic: ic, eb: eb}
}

func (c *articleConfigurator) ConfigureArticle(ctx context.Context) error {
	seServiceUseCase := sampleExtServiceUseCase.NewSampleExtServiceUseCase(c.eb.SampleExtGrpcService)
	kp := articleKafkaProducer.NewArticleProducer(c.ic.KafkaWriter)
	repo := articleRepo.NewArticleRepository(c.ic.Pg)
	uc := articleUseCase.NewArticleUseCase(repo, seServiceUseCase, kp)

	// grpc
	grpcController := articleGrpc.NewArticleGrpcController(uc)
	articleV1.RegisterArticleServiceServer(c.ic.GrpcServer.GetCurrentGrpcServer(), grpcController)

	// http
	httpRouterGp := c.ic.EchoServer.GetEchoInstance().Group(c.ic.EchoServer.GetBasePath())
	httpController := articleHttp.NewArticleHttpController(uc)
	articleHttp.NewArticleAPI(httpController).Register(httpRouterGp)

	// Consumers
	articleKafkaConsumer.NewArticleConsumer(c.ic.KafkaReader).RunConsumers(ctx)

	// Jobs
	articleJob.NewArticleJob().RunJobs(ctx)

	return nil
}

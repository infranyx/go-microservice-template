package articleConfigurator

import (
	"context"

	goTemplateUseCase "github.com/infranyx/go-grpc-template/external/go_template/usecase"
	articleGrpc "github.com/infranyx/go-grpc-template/internal/article/delivery/grpc"
	articleHttp "github.com/infranyx/go-grpc-template/internal/article/delivery/http"
	articleKafkaConsumer "github.com/infranyx/go-grpc-template/internal/article/delivery/kafka/consumer"
	articleKafkaProducer "github.com/infranyx/go-grpc-template/internal/article/delivery/kafka/producer"
	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	articleJob "github.com/infranyx/go-grpc-template/internal/article/job"
	articleRepo "github.com/infranyx/go-grpc-template/internal/article/repository"
	articleUseCase "github.com/infranyx/go-grpc-template/internal/article/usecase"
	clientContainer "github.com/infranyx/go-grpc-template/pkg/client_container"
	infraContainer "github.com/infranyx/go-grpc-template/pkg/infra_container"

	articleV1 "github.com/infranyx/protobuf-template-go/golang-grpc-template/article/v1"
)

type articleConfigurator struct {
	ic *infraContainer.IContainer
	cc *clientContainer.CContainer
}

func NewArticleConfigurator(ic *infraContainer.IContainer, cc *clientContainer.CContainer) articleDomain.ArticleConfigurator {
	return &articleConfigurator{ic: ic, cc: cc}
}

func (ac *articleConfigurator) ConfigureArticle(ctx context.Context) error {
	gtuc := goTemplateUseCase.NewGoTemplateUseCase(ac.cc.GoTemplateGrpc)
	rp := articleRepo.NewArticleRepository(ac.ic.Pg)
	ap := articleKafkaProducer.NewArticleProducer(ac.ic.KafkaWriter)
	uc := articleUseCase.NewArticleUseCase(rp, gtuc, ap)

	// grpc
	gc := articleGrpc.NewArticleGrpcController(uc)
	articleV1.RegisterArticleServiceServer(ac.ic.GrpcServer.GetCurrentGrpcServer(), gc)

	// http
	g := ac.ic.EchoServer.GetEchoInstance().Group(ac.ic.EchoServer.GetBasePath())
	hc := articleHttp.NewArticleHttpController(uc)
	articleHttp.NewArticleAPI(hc).Register(g)

	// Consumers
	articleKafkaConsumer.NewArticleConsumer(ac.ic.KafkaReader).RunConsumers(ctx)

	// Jobs
	articleJob.NewArticleJob().RunJobs(ctx)

	return nil
}

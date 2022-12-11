package articleConfigurator

import (
	"context"

	goTemplateUseCase "github.com/infranyx/go-grpc-template/external/go_template/usecase"
	articleGrpc "github.com/infranyx/go-grpc-template/internal/article/controllers/grpc"
	articleHttp "github.com/infranyx/go-grpc-template/internal/article/controllers/http"
	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
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
	uc := articleUseCase.NewArticleUseCase(rp, gtuc)

	// grpc
	gc := articleGrpc.NewArticleGrpcController(uc)
	articleV1.RegisterArticleServiceServer(ac.ic.GrpcServer.GetCurrentGrpcServer(), gc)

	// http
	g := ac.ic.EchoServer.GetEchoInstance().Group(ac.ic.EchoServer.GetBasePath())
	hc := articleHttp.NewArticleHttpController(uc)
	articleHttp.NewArticleAPI(hc).Register(g)

	return nil
}

package articleConfigurator

import (
	"context"
	articleGrpc "github.com/infranyx/go-grpc-template/internal/article/controllers/grpc"
	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	articleRepo "github.com/infranyx/go-grpc-template/internal/article/repository"
	articleUseCase "github.com/infranyx/go-grpc-template/internal/article/usecase"
	infraContainer "github.com/infranyx/go-grpc-template/pkg/infra_container"

	articleV1 "github.com/infranyx/protobuf-template-go/golang-grpc-template/article/v1"
)

type articleConfigurator struct {
	ic *infraContainer.IContainer
}

func NewArticleConfigurator(ic *infraContainer.IContainer) articleDomain.ArticleConfigurator {
	return &articleConfigurator{ic: ic}
}

func (ac *articleConfigurator) ConfigureArticle(ctx context.Context) error {
	rp := articleRepo.NewArticleRepository(ac.ic.Pg)
	uc := articleUseCase.NewArticleUseCase(rp)
	gc := articleGrpc.NewArticleGrpcController(uc)
	articleV1.RegisterArticleServiceServer(ac.ic.GrpcServer.GetCurrentGrpcServer(), gc)

	return nil
}

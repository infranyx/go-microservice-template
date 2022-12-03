package article_configurator

import (
	"context"

	article_grpc "github.com/infranyx/go-grpc-template/internal/article/controllers/grpc"
	article_repo "github.com/infranyx/go-grpc-template/internal/article/repository"
	article_usecase "github.com/infranyx/go-grpc-template/internal/article/usecase"
	infra "github.com/infranyx/go-grpc-template/shared/infra_container"
	articlev1 "github.com/infranyx/protobuf-template-go/golang-grpc-template/article/v1"
)

type articleControllerConfigurator struct {
	ic *infra.IContainer
}

func NewArticleConfigurator(ic *infra.IContainer) *articleControllerConfigurator {
	return &articleControllerConfigurator{ic: ic}
}

func (c *articleControllerConfigurator) ConfigureArticle(ctx context.Context) error {
	articleRepo := article_repo.NewArticleRepository(c.ic.Pg)
	articleUC := article_usecase.NewArticleUseCase(articleRepo)
	articleGrpcControllers := article_grpc.New(articleUC)
	articlev1.RegisterArticleServiceServer(c.ic.GrpcServer.GetCurrentGrpcServer(), articleGrpcControllers)

	return nil
}

package article_configurator

import (
	"context"

	article_grpc "github.com/infranyx/go-grpc-template/internal/article/controllers/grpc"
	article_repo "github.com/infranyx/go-grpc-template/internal/article/repository"
	article_usecase "github.com/infranyx/go-grpc-template/internal/article/usecase"
	"github.com/infranyx/go-grpc-template/pkg/grpc"
	"github.com/infranyx/go-grpc-template/shared/infrastructure"
	articlev1 "go.buf.build/grpc/go/infranyx/golang-grpc-template/article/v1"
)

type articleControllerConfigurator struct {
	grpcServer grpc.GrpcServer
	*infrastructure.InfrastructureConfiguration
}

func InitArticleConfigurator(ctx context.Context, infrastructureConfiguration *infrastructure.InfrastructureConfiguration, grpcServer grpc.GrpcServer) error {
	articleConfigurator := NewArticleConfigurator(infrastructureConfiguration, grpcServer)
	err := articleConfigurator.ConfigureArticle(ctx)
	return err
}

func NewArticleConfigurator(infrastructureConfiguration *infrastructure.InfrastructureConfiguration, grpcServer grpc.GrpcServer) *articleControllerConfigurator {
	return &articleControllerConfigurator{InfrastructureConfiguration: infrastructureConfiguration, grpcServer: grpcServer}
}

func (c *articleControllerConfigurator) ConfigureArticle(ctx context.Context) error {
	articleRepo := article_repo.NewArticleRepository(c.Pgx)
	articleUC := article_usecase.NewArticleUseCase(articleRepo)
	articleGrpcControllers := article_grpc.New(articleUC)
	articlev1.RegisterArticleServiceServer(c.grpcServer.GetCurrentGrpcServer(), articleGrpcControllers)

	return nil
}

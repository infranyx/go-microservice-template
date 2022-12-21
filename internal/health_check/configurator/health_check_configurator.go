package healthCheckConfigurator

import (
	"context"

	healthCheckHttp "github.com/infranyx/go-grpc-template/internal/health_check/delivery/http"
	healthCheckDomain "github.com/infranyx/go-grpc-template/internal/health_check/domain"
	healthCheckUseCase "github.com/infranyx/go-grpc-template/internal/health_check/usecase"
	infraContainer "github.com/infranyx/go-grpc-template/pkg/infra_container"
)

type healthCheckConfigurator struct {
	ic *infraContainer.IContainer
}

func NewHealthCheckConfigurator(ic *infraContainer.IContainer) healthCheckDomain.HealthCheckConfigurator {
	return &healthCheckConfigurator{ic: ic}
}

func (hc *healthCheckConfigurator) ConfigureHealthCheck(ctx context.Context) error {
	uc := healthCheckUseCase.NewHealthCheckUseCase()

	// grpc
	//gc := articleGrpc.NewArticleGrpcController(uc)
	//articleV1.RegisterArticleServiceServer(ac.ic.GrpcServer.GetCurrentGrpcServer(), gc)

	// http
	routerGroup := hc.ic.EchoServer.GetEchoInstance().Group(hc.ic.EchoServer.GetBasePath())
	healthCheckController := healthCheckHttp.NewHealthCheckHttpController(uc)
	healthCheckHttp.NewHealthCheckAPI(healthCheckController).Register(routerGroup)

	return nil
}

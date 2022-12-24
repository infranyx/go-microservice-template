package healthCheckConfigurator

import (
	"context"
	kafkaHealthCheckUseCase "github.com/infranyx/go-grpc-template/internal/health_check/usecase/kafka_health_check"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"

	healthCheckGrpc "github.com/infranyx/go-grpc-template/internal/health_check/delivery/grpc"
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
	uc := healthCheckUseCase.NewHealthCheckUseCase(hc.ic.Pg)
	kuc := kafkaHealthCheckUseCase.NewKafkaHealthCheck()
	// grpc
	gc := healthCheckGrpc.NewHealthCheckGrpcController(uc, kuc)
	grpcHealthV1.RegisterHealthServer(hc.ic.GrpcServer.GetCurrentGrpcServer(), gc)

	// http
	routerGroup := hc.ic.EchoServer.GetEchoInstance().Group(hc.ic.EchoServer.GetBasePath())
	healthCheckController := healthCheckHttp.NewHealthCheckHttpController(uc)
	healthCheckHttp.NewHealthCheckAPI(healthCheckController).Register(routerGroup)

	return nil
}

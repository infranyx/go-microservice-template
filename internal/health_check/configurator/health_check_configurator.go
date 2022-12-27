package healthCheckConfigurator

import (
	"context"
	kafkaHealthCheckUseCase "github.com/infranyx/go-grpc-template/internal/health_check/usecase/kafka_health_check"
	pgHealthCheckUseCase "github.com/infranyx/go-grpc-template/internal/health_check/usecase/pg_health_check"
	tmpDirHealthCheckUseCase "github.com/infranyx/go-grpc-template/internal/health_check/usecase/tmp_dir_health_check"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"

	healthCheckGrpc "github.com/infranyx/go-grpc-template/internal/health_check/delivery/grpc"
	healthCheckHttp "github.com/infranyx/go-grpc-template/internal/health_check/delivery/http"
	healthCheckDomain "github.com/infranyx/go-grpc-template/internal/health_check/domain"
	healthCheckUseCase "github.com/infranyx/go-grpc-template/internal/health_check/usecase"
	infraContainer "github.com/infranyx/go-grpc-template/pkg/infra_container"
)

type configurator struct {
	ic *infraContainer.IContainer
}

func NewConfigurator(ic *infraContainer.IContainer) healthCheckDomain.Configurator {
	return &configurator{ic: ic}
}

func (c *configurator) Configure(ctx context.Context) error {
	postgresHealthCheckUc := pgHealthCheckUseCase.NewUseCase(c.ic.Pg)
	kafkaHealthCheckUc := kafkaHealthCheckUseCase.NewUseCase()
	tmpDirHealthCheckUc := tmpDirHealthCheckUseCase.NewUseCase()

	healthCheckUc := healthCheckUseCase.NewUseCase(postgresHealthCheckUc, kafkaHealthCheckUc, tmpDirHealthCheckUc)

	// grpc
	grpcController := healthCheckGrpc.NewController(healthCheckUc, postgresHealthCheckUc, kafkaHealthCheckUc, tmpDirHealthCheckUc)
	grpcHealthV1.RegisterHealthServer(c.ic.GrpcServer.GetCurrentGrpcServer(), grpcController)

	// http
	httpRouterGp := c.ic.EchoServer.GetEchoInstance().Group(c.ic.EchoServer.GetBasePath())
	healthCheckController := healthCheckHttp.NewController(healthCheckUc)
	healthCheckHttp.NewRouter(healthCheckController).Register(httpRouterGp)

	return nil
}

package healthCheckFixture

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/test/bufconn"

	healthCheckGrpc "github.com/infranyx/go-grpc-template/internal/health_check/delivery/grpc"
	healthCheckHttp "github.com/infranyx/go-grpc-template/internal/health_check/delivery/http"
	healthCheckUseCase "github.com/infranyx/go-grpc-template/internal/health_check/usecase"
	kafkaHealthCheckUseCase "github.com/infranyx/go-grpc-template/internal/health_check/usecase/kafka_health_check"
	postgresHealthCheckUseCase "github.com/infranyx/go-grpc-template/internal/health_check/usecase/postgres_health_check"
	tmpDirHealthCheckUseCase "github.com/infranyx/go-grpc-template/internal/health_check/usecase/tmp_dir_health_check"
	"github.com/infranyx/go-grpc-template/pkg/logger"

	iContainer "github.com/infranyx/go-grpc-template/pkg/infra_container"
)

type IntegrationTestFixture struct {
	TearDown              func()
	Ctx                   context.Context
	Cancel                context.CancelFunc
	InfraContainer        *iContainer.IContainer
	HealthCheckGrpcClient grpcHealthV1.HealthClient
}

const BUFSIZE = 1024 * 1024

func NewIntegrationTestFixture() (*IntegrationTestFixture, error) {
	deadline := time.Now().Add(time.Duration(1 * time.Minute))
	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	ic, infraDown, err := iContainer.NewIC(ctx)
	if err != nil {
		cancel()
		return nil, err
	}

	postgresHealthCheckUc := postgresHealthCheckUseCase.NewUseCase(ic.Postgres)
	kafkaHealthCheckUc := kafkaHealthCheckUseCase.NewUseCase()
	tmpDirHealthCheckUc := tmpDirHealthCheckUseCase.NewUseCase()

	healthCheckUc := healthCheckUseCase.NewUseCase(postgresHealthCheckUc, kafkaHealthCheckUc, tmpDirHealthCheckUc)

	// http
	ic.EchoHttpServer.SetupDefaultMiddlewares()
	httpRouterGp := ic.EchoHttpServer.GetEchoInstance().Group(ic.EchoHttpServer.GetBasePath())
	httpController := healthCheckHttp.NewController(healthCheckUc)
	healthCheckHttp.NewRouter(httpController).Register(httpRouterGp)

	// grpc
	grpcController := healthCheckGrpc.NewController(healthCheckUc, postgresHealthCheckUc, kafkaHealthCheckUc, tmpDirHealthCheckUc)
	grpcHealthV1.RegisterHealthServer(ic.GrpcServer.GetCurrentGrpcServer(), grpcController)

	lis := bufconn.Listen(BUFSIZE)
	go func() {
		if err := ic.GrpcServer.GetCurrentGrpcServer().Serve(lis); err != nil {
			logger.Zap.Sugar().Fatalf("Server exited with error: %v", err)
		}
	}()

	grpcClientConn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		cancel()
		return nil, err
	}

	healthCheckGrpcClient := grpcHealthV1.NewHealthClient(grpcClientConn)

	return &IntegrationTestFixture{
		TearDown: func() {
			cancel()
			infraDown()
			_ = grpcClientConn.Close()
		},
		InfraContainer:        ic,
		Ctx:                   ctx,
		Cancel:                cancel,
		HealthCheckGrpcClient: healthCheckGrpcClient,
	}, nil
}

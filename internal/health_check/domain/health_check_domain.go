package healthCheckDomain

import (
	"context"
	healthCheckDto "github.com/infranyx/go-grpc-template/internal/health_check/dto"
	"github.com/labstack/echo/v4"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
)

type Configurator interface {
	Configure(ctx context.Context) error
}

type GrpcController interface {
	Check(ctx context.Context, request *grpcHealthV1.HealthCheckRequest) (*grpcHealthV1.HealthCheckResponse, error)
	Watch(request *grpcHealthV1.HealthCheckRequest, server grpcHealthV1.Health_WatchServer) error
}

type HttpController interface {
	Check(c echo.Context) error
}

type HealthCheckUseCase interface {
	Check() *healthCheckDto.HealthCheckResponseDto
}

type PostgresHealthCheckUseCase interface {
	Check() bool
}

type TmpDirHealthCheckUseCase interface {
	Check() bool
}

type KafkaHealthCheckUseCase interface {
	Check() bool
}

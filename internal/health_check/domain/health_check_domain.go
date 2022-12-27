package healthCheckDomain

import (
	"context"
	"github.com/labstack/echo/v4"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
)

type HealthCheckUnit struct {
	Unit string `json:"unit"`
	Up   bool   `json:"up"`
}

type HealthCheckResult struct {
	Status bool              `json:"status"`
	Units  []HealthCheckUnit `json:"units"`
}

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
	Check() *HealthCheckResult
}

type PgHealthCheckUseCase interface {
	Check() bool
}

type TmpDirHealthCheckUseCase interface {
	Check() bool
}

type KafkaHealthCheckUseCase interface {
	Check() bool
}

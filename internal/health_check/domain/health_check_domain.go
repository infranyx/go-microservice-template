package healthCheckDomain

import (
	"context"
	"github.com/labstack/echo/v4"
)

type HealthCheckUnit struct {
	Unit string `json:"unit"`
	Up   bool   `json:"up"`
}

type HealthCheckResult struct {
	Status bool              `json:"status"`
	Units  []HealthCheckUnit `json:"units"`
}

type HealthCheckConfigurator interface {
	ConfigureHealthCheck(ctx context.Context) error
}

type HealthCheckHttpController interface {
	Check(c echo.Context) error
}

type HealthCheckUseCase interface {
	Check(ctx context.Context) (*HealthCheckResult, error)
}

type PgHealthCheckUseCase interface {
	PingCheck() bool
}

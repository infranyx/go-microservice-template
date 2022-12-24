package healthCheckUseCase

import (
	"context"
	"github.com/infranyx/go-grpc-template/pkg/postgres"

	healthCheckDomain "github.com/infranyx/go-grpc-template/internal/health_check/domain"
	pgHealthCheckUseCase "github.com/infranyx/go-grpc-template/internal/health_check/usecase/pg_health_check"
)

type healthCheckUseCase struct {
	conn *postgres.Postgres
}

func NewHealthCheckUseCase(conn *postgres.Postgres) healthCheckDomain.HealthCheckUseCase {
	return &healthCheckUseCase{
		conn: conn,
	}
}

func (hu *healthCheckUseCase) Check(ctx context.Context) (*healthCheckDomain.HealthCheckResult, error) {
	healthCheckResult := healthCheckDomain.HealthCheckResult{
		Status: true,
		Units:  nil,
	}

	pgHealthCheck := pgHealthCheckUseCase.NewPgHealthCheck(hu.conn)
	pgUnit := healthCheckDomain.HealthCheckUnit{
		Unit: "postgres",
		Up:   pgHealthCheck.PingCheck(),
	}
	healthCheckResult.Units = append(healthCheckResult.Units, pgUnit)

	for _, v := range healthCheckResult.Units {
		if !v.Up {
			healthCheckResult.Status = false
			break
		}
	}

	return &healthCheckResult, nil
}

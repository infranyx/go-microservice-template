package healthCheckUseCase

import (
	"context"
	tmpDirHealthCheckUseCase "github.com/infranyx/go-grpc-template/internal/health_check/usecase/tmp_dir_health_check"
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
		Unit: "database",
		Up:   pgHealthCheck.PingCheck(),
	}
	healthCheckResult.Units = append(healthCheckResult.Units, pgUnit)

	tmpDirHealthCheck := tmpDirHealthCheckUseCase.NewPgHealthCheck()
	tmpDirUnit := healthCheckDomain.HealthCheckUnit{
		Unit: "writable-tmp-dir",
		Up:   tmpDirHealthCheck.PingCheck(),
	}
	healthCheckResult.Units = append(healthCheckResult.Units, tmpDirUnit)

	for _, v := range healthCheckResult.Units {
		if !v.Up {
			healthCheckResult.Status = false
			break
		}
	}

	return &healthCheckResult, nil
}

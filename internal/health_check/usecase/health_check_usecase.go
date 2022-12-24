package healthCheckUseCase

import (
	kafkaHealthCheckUseCase "github.com/infranyx/go-grpc-template/internal/health_check/usecase/kafka_health_check"
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

func (hu *healthCheckUseCase) Check() *healthCheckDomain.HealthCheckResult {
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

	tmpDirHealthCheck := tmpDirHealthCheckUseCase.NewTmpDirHealthCheck()
	tmpDirUnit := healthCheckDomain.HealthCheckUnit{
		Unit: "writable-tmp-dir",
		Up:   tmpDirHealthCheck.PingCheck(),
	}
	healthCheckResult.Units = append(healthCheckResult.Units, tmpDirUnit)

	kafkaHealthCheck := kafkaHealthCheckUseCase.NewKafkaHealthCheck()
	kafkaUnit := healthCheckDomain.HealthCheckUnit{
		Unit: "kafka",
		Up:   kafkaHealthCheck.PingCheck(),
	}
	healthCheckResult.Units = append(healthCheckResult.Units, kafkaUnit)

	for _, v := range healthCheckResult.Units {
		if !v.Up {
			healthCheckResult.Status = false
			break
		}
	}

	return &healthCheckResult
}

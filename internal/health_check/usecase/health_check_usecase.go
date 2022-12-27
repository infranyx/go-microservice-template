package healthCheckUseCase

import (
	healthCheckDomain "github.com/infranyx/go-grpc-template/internal/health_check/domain"
)

type useCase struct {
	postgresHealthCheckUc healthCheckDomain.PgHealthCheckUseCase
	kafkaHealthCheckUc    healthCheckDomain.KafkaHealthCheckUseCase
	tmpDirHealthCheckUc   healthCheckDomain.TmpDirHealthCheckUseCase
}

func NewUseCase(
	postgresHealthCheckUc healthCheckDomain.PgHealthCheckUseCase,
	kafkaHealthCheckUc healthCheckDomain.KafkaHealthCheckUseCase,
	tmpDirHealthCheckUc healthCheckDomain.TmpDirHealthCheckUseCase,
) healthCheckDomain.HealthCheckUseCase {
	return &useCase{
		postgresHealthCheckUc: postgresHealthCheckUc,
		kafkaHealthCheckUc:    kafkaHealthCheckUc,
		tmpDirHealthCheckUc:   tmpDirHealthCheckUc,
	}
}

func (uc *useCase) Check() *healthCheckDomain.HealthCheckResult {
	healthCheckResult := healthCheckDomain.HealthCheckResult{
		Status: true,
		Units:  nil,
	}

	pgUnit := healthCheckDomain.HealthCheckUnit{
		Unit: "postgres",
		Up:   uc.postgresHealthCheckUc.Check(),
	}
	healthCheckResult.Units = append(healthCheckResult.Units, pgUnit)

	kafkaUnit := healthCheckDomain.HealthCheckUnit{
		Unit: "kafka",
		Up:   uc.kafkaHealthCheckUc.Check(),
	}
	healthCheckResult.Units = append(healthCheckResult.Units, kafkaUnit)

	tmpDirUnit := healthCheckDomain.HealthCheckUnit{
		Unit: "writable-tmp-dir",
		Up:   uc.tmpDirHealthCheckUc.Check(),
	}
	healthCheckResult.Units = append(healthCheckResult.Units, tmpDirUnit)

	for _, v := range healthCheckResult.Units {
		if !v.Up {
			healthCheckResult.Status = false
			break
		}
	}

	return &healthCheckResult
}

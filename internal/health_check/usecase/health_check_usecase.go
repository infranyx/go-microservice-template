package healthCheckUseCase

import (
	healthCheckDomain "github.com/infranyx/go-microservice-template/internal/health_check/domain"
	healthCheckDto "github.com/infranyx/go-microservice-template/internal/health_check/dto"
)

type useCase struct {
	postgresHealthCheckUc healthCheckDomain.PostgresHealthCheckUseCase
	kafkaHealthCheckUc    healthCheckDomain.KafkaHealthCheckUseCase
	tmpDirHealthCheckUc   healthCheckDomain.TmpDirHealthCheckUseCase
}

func NewUseCase(
	postgresHealthCheckUc healthCheckDomain.PostgresHealthCheckUseCase,
	kafkaHealthCheckUc healthCheckDomain.KafkaHealthCheckUseCase,
	tmpDirHealthCheckUc healthCheckDomain.TmpDirHealthCheckUseCase,
) healthCheckDomain.HealthCheckUseCase {
	return &useCase{
		postgresHealthCheckUc: postgresHealthCheckUc,
		kafkaHealthCheckUc:    kafkaHealthCheckUc,
		tmpDirHealthCheckUc:   tmpDirHealthCheckUc,
	}
}

func (uc *useCase) Check() *healthCheckDto.HealthCheckResponseDto {
	healthCheckResult := healthCheckDto.HealthCheckResponseDto{
		Status: true,
		Units:  nil,
	}

	pgUnit := healthCheckDto.HealthCheckUnit{
		Unit: "postgres",
		Up:   uc.postgresHealthCheckUc.Check(),
	}
	healthCheckResult.Units = append(healthCheckResult.Units, pgUnit)

	kafkaUnit := healthCheckDto.HealthCheckUnit{
		Unit: "kafka",
		Up:   uc.kafkaHealthCheckUc.Check(),
	}
	healthCheckResult.Units = append(healthCheckResult.Units, kafkaUnit)

	tmpDirUnit := healthCheckDto.HealthCheckUnit{
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

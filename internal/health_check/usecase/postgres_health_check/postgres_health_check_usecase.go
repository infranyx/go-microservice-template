package postgresHealthCheckUseCase

import (
	healthCheckDomain "github.com/infranyx/go-microservice-template/internal/health_check/domain"
	"github.com/infranyx/go-microservice-template/pkg/postgres"
)

type useCase struct {
	postgres *postgres.Postgres
}

func NewUseCase(postgres *postgres.Postgres) healthCheckDomain.PostgresHealthCheckUseCase {
	return &useCase{
		postgres: postgres,
	}
}

func (uc *useCase) Check() bool {
	if err := uc.postgres.SqlxDB.DB.Ping(); err != nil {
		return false
	}
	return true
}

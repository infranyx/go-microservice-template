package pgHealthCheckUseCase

import (
	healthCheckDomain "github.com/infranyx/go-grpc-template/internal/health_check/domain"
	"github.com/infranyx/go-grpc-template/pkg/postgres"
)

type pgHealthCheck struct {
	conn *postgres.Postgres
}

func NewPgHealthCheck(Conn *postgres.Postgres) healthCheckDomain.PgHealthCheckUseCase {
	return &pgHealthCheck{
		conn: Conn,
	}
}

func (ph *pgHealthCheck) PingCheck() bool {
	if err := ph.conn.SqlxDB.DB.Ping(); err != nil {
		return false
	}
	return true
}

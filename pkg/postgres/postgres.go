package postgres

import (
	"context"

	"github.com/infranyx/go-grpc-template/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // pgx also supported
)

type PostgreSqlx struct {
	DB *sqlx.DB
}

func NewPostgreSqlx() (*PostgreSqlx, error) {
	conf := config.Conf
	DB, err := sqlx.ConnectContext(context.Background(), "postgres", conf.Postgres.Url)
	if err != nil {
		return nil, err
	}

	return &PostgreSqlx{
		DB: DB,
	}, nil
}

func (p *PostgreSqlx) Close() {
	p.DB.Close()
}

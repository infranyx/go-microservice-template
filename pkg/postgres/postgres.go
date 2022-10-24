package postgres

import (
	"context"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // pgx also supported
)

type PostgreSqlx struct {
	DB *sqlx.DB
}

func NewPostgreSqlx() (*PostgreSqlx, error) {
	DB, err := sqlx.ConnectContext(context.Background(), "postgres", os.Getenv("DATABASE_URL"))
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

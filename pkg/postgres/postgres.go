package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // pgx also supported
)

type PostgreSqlx struct {
	DB *sqlx.DB
}

func NewPostgreSqlx() (*PostgreSqlx, error) {
	fmt.Println(os.Getenv("PG_URL"))
	DB, err := sqlx.ConnectContext(context.Background(), "postgres", `postgres://postgres:postgrespw@localhost:5432/postgres?sslmode=disable`)
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

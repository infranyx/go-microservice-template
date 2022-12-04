package postgres

import (
	"context"
	"fmt"
	"github.com/infranyx/go-grpc-template/pkg/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type PgConf struct {
	Host    string
	Port    string
	User    string
	Pass    string
	DBName  string
	SslMode string
}

type Postgres struct {
	SqlxDB *sqlx.DB
}

func (db *Postgres) Close() {
	_ = db.SqlxDB.DB.Close()
	_ = db.SqlxDB.Close()
}

func NewPgConn(ctx context.Context, conf *PgConf) (*Postgres, error) {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Pass,
		conf.DBName,
		conf.SslMode,
	)

	fmt.Println(connString)

	db, err := sqlx.ConnectContext(ctx, "postgres", connString)
	if err != nil {
		// TODO : Log + Err
		return nil, err
	}

	db.SetMaxOpenConns(config.Conf.Postgres.MaxConn)                           // the defaultLogger is 0 (unlimited)
	db.SetMaxIdleConns(config.Conf.Postgres.MaxIdleConn)                       // defaultMaxIdleConn = 2
	db.SetConnMaxLifetime(time.Duration(config.Conf.Postgres.MaxLifeTimeConn)) // 0, connections are reused forever

	if err := db.Ping(); err != nil {
		// TODO : Log + Err
		fmt.Println("can not ping postgres")
		defer db.Close()
		return nil, err
	}

	return &Postgres{SqlxDB: db}, nil
}

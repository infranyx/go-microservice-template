package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/infranyx/go-grpc-template/pkg/config"

	_ "github.com/lib/pq"
)

type Postgres struct {
	SqlxDB *sqlx.DB
	// DB     *sql.DB
	// SquirrelBuilder squirrel.StatementBuilderType
	// GoquBuilder     *goqu.SelectDataset
}

var PgConn *Postgres

func init() {
	NewPgConn()
}

func NewPgConn() {
	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Conf.Postgres.Host,
		config.Conf.Postgres.Port,
		config.Conf.Postgres.User,
		config.Conf.Postgres.DBName,
		config.Conf.Postgres.Pass,
		config.Conf.Postgres.SslMode,
	)

	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		// TODO : Log + Err
		log.Fatal("can not connect to postgres", err)
	}

	db.SetMaxOpenConns(config.Conf.Postgres.MaxConn)                           // the defaultLogger is 0 (unlimited)
	db.SetMaxIdleConns(config.Conf.Postgres.MaxIdleConn)                       // defaultMaxIdleConn = 2
	db.SetConnMaxLifetime(time.Duration(config.Conf.Postgres.MaxLifeTimeConn)) // 0, connections are reused forever

	if err := db.Ping(); err != nil {
		// TODO : Log + Err
		fmt.Println("can not ping postgres")
		defer db.Close()
	}

	PgConn = &Postgres{SqlxDB: db}
}

func (db *Postgres) Close() {
	_ = db.SqlxDB.DB.Close()
	_ = db.SqlxDB.Close()
}

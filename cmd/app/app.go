package app

import (
	"fmt"
	"log"

	"github.com/infranyx/go-grpc-template/pkg/postgres"
)

func Run() {

	dbSQLX, err := postgres.NewPostgreSqlx()
	if err != nil {
		log.Fatalf("Could not initialize Database connection using sqlx %s", err)
	}
	defer dbSQLX.Close()

	fmt.Println("Hello from app")
}

package main

import (
	"github.com/infranyx/go-grpc-template/app"
	"github.com/infranyx/go-grpc-template/config"
	"github.com/infranyx/go-grpc-template/pkg/logger"
)

func main() {
	config.NewConfig()
	l := logger.NewLogger()

	application := app.NewServer(l)
	err := application.Run()
	if err != nil {
		l.Error(err)
	}
}

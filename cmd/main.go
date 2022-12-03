package main

import (
	"github.com/infranyx/go-grpc-template/app"
	"github.com/infranyx/go-grpc-template/config"
	"github.com/infranyx/go-grpc-template/pkg/logger"
)

func main() {
	config.NewConfig()
	l := logger.Zap

	application := app.New()
	err := application.Run()
	if err != nil {
		l.Sugar().Error(err)
	}
}

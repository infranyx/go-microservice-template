package main

import (
	"github.com/infranyx/go-grpc-template/app"
	"github.com/infranyx/go-grpc-template/pkg/logger"
)

func main() {
	application := app.New()
	err := application.Run()
	if err != nil {
		logger.Zap.Sugar().Error(err)
	}
}

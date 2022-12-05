package main

import (
	"github.com/infranyx/go-grpc-template/app"
	"github.com/infranyx/go-grpc-template/pkg/logger"
)

func main() {
	err := app.New().Run()
	if err != nil {
		logger.Zap.Sugar().Fatal(err)
	}
}

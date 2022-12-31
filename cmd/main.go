package main

import (
	"github.com/infranyx/go-microservice-template/app"
	"github.com/infranyx/go-microservice-template/pkg/logger"
)

func main() {
	err := app.New().Run()
	if err != nil {
		logger.Zap.Sugar().Fatal(err)
	}
}

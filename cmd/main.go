package main

import (
	"github.com/infranyx/go-grpc-template/app"
	"github.com/infranyx/go-grpc-template/config"
	"github.com/infranyx/go-grpc-template/pkg/logger"
)

func main() {
	config.NewConfig()
	logger := logger.NewLogger()

	// s := sentry.NewSentryClient()

	application := app.NewServer(logger)
	err := application.Run()
	if err != nil {
		logger.Error(err)
	}
}

package main

import (
	"fmt"

	"github.com/infranyx/go-grpc-template/app"
	"github.com/infranyx/go-grpc-template/config"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/infranyx/go-grpc-template/pkg/sentry"
)

func main() {
	config.NewConfig()
	logger := logger.NewLogger()
	logger.Infow("test",
		// Structured context as loosely typed key-value pairs.
		"url", 4,
		"attempt", 3)

	s := sentry.NewSentryClient()
	s.CaptureException(fmt.Errorf("error from golang"))

	application := app.NewServer(logger)
	err := application.Run()
	if err != nil {
		logger.Error(err)
	}
}

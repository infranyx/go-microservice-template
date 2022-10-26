package main

import (
	"fmt"

	"github.com/infranyx/go-grpc-template/app"
	"github.com/infranyx/go-grpc-template/config"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/infranyx/go-grpc-template/pkg/sentry"
)

func main() {
	conf := config.New()
	fmt.Println(conf.App.Port)

	l := logger.NewLogger()
	l.Infow("test",
		// Structured context as loosely typed key-value pairs.
		"url", 4,
		"attempt", 3)

	s := sentry.NewSentryClient()
	s.CaptureException(fmt.Errorf("error from golang"))

	fmt.Println("Hello world")
	app.Run()
}

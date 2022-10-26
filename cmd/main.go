package main

import (
	"fmt"

	"github.com/infranyx/go-grpc-template/cmd/app"
	"github.com/infranyx/go-grpc-template/config"
	"github.com/infranyx/go-grpc-template/pkg/logger"
)

func main() {
	config.Init()

	l := logger.NewLogger()
	l.Infow("test",
		// Structured context as loosely typed key-value pairs.
		"url", 4,
		"attempt", 3)

	fmt.Println("Hello world")
	app.Run()
}

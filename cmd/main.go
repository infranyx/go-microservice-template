package main

import (
	"fmt"

	"github.com/infranyx/go-grpc-template/cmd/app"
	"github.com/infranyx/go-grpc-template/config"
)

func main() {
	config.Init()

	fmt.Println("Hello world")
	app.Run()
}

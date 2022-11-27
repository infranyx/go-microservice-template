package main

import (
	"fmt"

	"github.com/infranyx/go-grpc-template/pkg/config"
)

func main() {

	fmt.Println(config.IsProdEnv())
}

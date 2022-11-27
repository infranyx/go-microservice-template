package main

import (
	"fmt"

	"github.com/infranyx/go-grpc-template/pkg/env"
)

func main() {

	url := env.New("APPLICATION_URL", []string{"foo", "jj"}).AsStringSlice(",")
	port := env.New("APPLICATION_PORT", 4000).AsInt()

	fmt.Println(url[0], port)
}

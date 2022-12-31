package externalBridge

import (
	"context"

	"github.com/infranyx/go-microservice-template/pkg/config"
	"github.com/infranyx/go-microservice-template/pkg/grpc"
)

type ExternalBridge struct {
	SampleExtGrpcService grpc.Client /* Like ETH Service */
}

func NewExternalBridge(ctx context.Context) (*ExternalBridge, func(), error) {
	var downFns []func()
	down := func() {
		for _, df := range downFns {
			df()
		}
	}
	sampleExtGrpcClient, _ := grpc.NewGrpcClient(
		ctx,
		&grpc.Config{Port: config.BaseConfig.SampleExtService.Port, Host: config.BaseConfig.SampleExtService.Host},
	)
	downFns = append(downFns, func() {
		_ = sampleExtGrpcClient.Close()
	})

	cc := &ExternalBridge{SampleExtGrpcService: sampleExtGrpcClient}

	return cc, down, nil
}

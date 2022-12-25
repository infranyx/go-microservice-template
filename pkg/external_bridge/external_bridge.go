package externalBridge

import (
	"context"

	"github.com/infranyx/go-grpc-template/pkg/config"
	"github.com/infranyx/go-grpc-template/pkg/grpc"
)

type ExternalBridge struct {
	SampleExtGrpcService grpc.GrpcClient /* Like ETH Service */
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
		&grpc.GrpcConfig{Port: config.BaseConfig.GoTemplateGrpcClient.Port, Host: config.BaseConfig.GoTemplateGrpcClient.Host},
	)
	downFns = append(downFns, func() {
		_ = sampleExtGrpcClient.Close()
	})

	cc := &ExternalBridge{SampleExtGrpcService: sampleExtGrpcClient}

	return cc, down, nil
}

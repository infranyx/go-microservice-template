package clientContainer

import (
	"context"

	"github.com/infranyx/go-grpc-template/pkg/config"
	"github.com/infranyx/go-grpc-template/pkg/grpc"
)

type CContainer struct {
	GoTemplateGrpc grpc.GrpcClient
}

func NewCC(ctx context.Context) (*CContainer, func(), error) {
	var downFns []func()
	down := func() {
		for _, df := range downFns {
			df()
		}
	}
	gtc, _ := grpc.NewGrpcClient(
		ctx,
		&grpc.GrpcConfig{Port: config.Conf.GoTemplateGrpcClient.Port, Host: config.Conf.GoTemplateGrpcClient.Host},
	)
	downFns = append(downFns, func() {
		gtc.Close()
	})

	cc := &CContainer{GoTemplateGrpc: gtc}

	return cc, down, nil
}

package grpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcClient struct {
	conn *grpc.ClientConn
}

type GrpcClient interface {
	GetGrpcConnection() *grpc.ClientConn
	Close() error
}

func NewGrpcClient(ctx context.Context, config *GrpcConfig) (GrpcClient, error) {
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%d", config.Host, config.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	return &grpcClient{conn: conn}, nil
}

func (g *grpcClient) GetGrpcConnection() *grpc.ClientConn {
	return g.conn
}

func (g *grpcClient) Close() error {
	return g.conn.Close()
}

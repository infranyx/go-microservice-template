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

type Client interface {
	GetGrpcConnection() *grpc.ClientConn
	Close() error
}

func NewGrpcClient(ctx context.Context, config *Config) (Client, error) {
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%d", config.Host, config.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClient{conn: conn}, nil
}

func (gc *grpcClient) GetGrpcConnection() *grpc.ClientConn {
	return gc.conn
}

func (gc *grpcClient) Close() error {
	return gc.conn.Close()
}

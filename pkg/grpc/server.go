package grpc

import (
	"context"
	"fmt"
	"net"
	"time"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcCtxTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	grpcInterceptors "github.com/infranyx/go-grpc-template/pkg/grpc/interceptors"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	googleGrpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

const (
	maxConnectionIdle = 5
	gRPCTimeout       = 15
	maxConnectionAge  = 5
	gRPCTime          = 10
)

type GrpcServer interface {
	RunGrpcServer(ctx context.Context, configGrpc func(grpcServer *googleGrpc.Server)) error
	GracefulShutdown()
	GetCurrentGrpcServer() *googleGrpc.Server
}

type GrpcConfig struct {
	Port        int
	Host        string
	Development bool
}

type grpcServer struct {
	server *googleGrpc.Server
	logger *logger.Logger
	config *GrpcConfig
}

func NewGrpcServer(conf *GrpcConfig, logger *logger.Logger) *grpcServer {
	s := googleGrpc.NewServer(
		googleGrpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: maxConnectionIdle * time.Minute,
			Timeout:           gRPCTimeout * time.Second,
			MaxConnectionAge:  maxConnectionAge * time.Minute,
			Time:              gRPCTime * time.Minute,
		}),

		googleGrpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
			grpcInterceptors.StreamServerInterceptor(),
		)),
		googleGrpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			grpcInterceptors.UnaryServerInterceptor(),
			grpcCtxTags.UnaryServerInterceptor(),
			grpcRecovery.UnaryServerInterceptor(),
		)),
	)

	return &grpcServer{server: s, logger: logger, config: conf}
}

func (s *grpcServer) RunGrpcServer(ctx context.Context, configGrpc func(grpcServer *googleGrpc.Server)) error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.Port))
	if err != nil {
		return err
	}
	if configGrpc != nil {
		configGrpc(s.server)
	}

	grpcPrometheus.Register(s.server)

	if s.config.Development {
		reflection.Register(s.server)
	}

	go func() {
		for {
			<-ctx.Done()
			s.logger.Infof("App is shutting down Grpc PORT: {%d}", s.config.Port)
			s.GracefulShutdown()
			return
		}
	}()

	s.logger.Infof("[grpcServer.RunGrpcServer] Writer gRPC server is listening on: %s:%d", s.config.Host, s.config.Port)

	err = s.server.Serve(l)
	if err != nil {
		s.logger.Error(fmt.Sprintf("[grpcServer_RunGrpcServer.Serve] grpc server serve error: %+v", err))
	}
	return err
}

func (s *grpcServer) GetCurrentGrpcServer() *googleGrpc.Server {
	return s.server
}

func (s *grpcServer) GracefulShutdown() {
	s.server.Stop()
	s.server.GracefulStop()
}

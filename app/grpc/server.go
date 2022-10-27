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
	"github.com/infranyx/go-grpc-template/config"
	const_app_env "github.com/infranyx/go-grpc-template/constant/app_env"
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

type grpcServer struct {
	server *googleGrpc.Server
	logger *logger.Logger
}

func NewGrpcServer(logger *logger.Logger) *grpcServer {
	s := googleGrpc.NewServer(
		googleGrpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: maxConnectionIdle * time.Minute,
			Timeout:           gRPCTimeout * time.Second,
			MaxConnectionAge:  maxConnectionAge * time.Minute,
			Time:              gRPCTime * time.Minute,
		}),

		googleGrpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
		// grpcError.StreamServerInterceptor(),
		)),
		googleGrpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			grpcCtxTags.UnaryServerInterceptor(),
			grpcRecovery.UnaryServerInterceptor(),
		)),
	)

	return &grpcServer{server: s, logger: logger}
}

func (s *grpcServer) RunGrpcServer(ctx context.Context, configGrpc func(grpcServer *googleGrpc.Server)) error {
	conf := config.Conf
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Grpc.Port))
	if err != nil {
		return err
	}
	if configGrpc != nil {
		configGrpc(s.server)
	}

	grpcPrometheus.Register(s.server)

	if conf.App.AppEnv == const_app_env.DEV {
		reflection.Register(s.server)
	}

	go func() {
		for {
			<-ctx.Done()
			s.logger.Infof("App is shutting down Grpc PORT: {%d}", conf.Grpc.Port)
			s.GracefulShutdown()
			return
		}
	}()

	s.logger.Infof("[grpcServer.RunGrpcServer] Writer gRPC server is listening on: %s:%d", conf.Grpc.Host, conf.Grpc.Port)

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

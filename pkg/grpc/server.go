package grpc

import (
	"context"
	"fmt"

	"net"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcCtxTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcErrorInterceptor "github.com/infranyx/go-grpc-template/pkg/grpc/interceptors/error_interceptor"
	grpcLoggerInterceptor "github.com/infranyx/go-grpc-template/pkg/grpc/interceptors/logger_interceptor"
	grpcSentryInterceptor "github.com/infranyx/go-grpc-template/pkg/grpc/interceptors/sentry_interceptor"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	sentryUtils "github.com/infranyx/go-grpc-template/pkg/sentry/sentry_utils"
	googleGrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	config *GrpcConfig
}

func NewGrpcServer(conf *GrpcConfig) GrpcServer {
	gso := &sentryUtils.Options{
		Repanic: true,
	}

	s := googleGrpc.NewServer(
		googleGrpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			grpcSentryInterceptor.UnaryServerInterceptor(gso),
			grpcErrorInterceptor.UnaryServerInterceptor(),
			grpcLoggerInterceptor.UnaryServerInterceptor(),
			grpcCtxTags.UnaryServerInterceptor(),
			grpcRecovery.UnaryServerInterceptor(),
		)),
		googleGrpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
			grpcSentryInterceptor.StreamServerInterceptor(gso),
			grpcErrorInterceptor.StreamServerInterceptor(),
		)),
	)

	return &grpcServer{server: s, config: conf}
}

func (s *grpcServer) RunGrpcServer(ctx context.Context, configGrpc func(grpcServer *googleGrpc.Server)) error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.Port))
	if err != nil {
		return err
	}
	if configGrpc != nil {
		configGrpc(s.server)
	}

	if s.config.Development {
		reflection.Register(s.server)
	}

	go func() {
		<-ctx.Done()
		logger.Zap.Sugar().Infof("App is shutting down Grpc PORT: {%d}", s.config.Port)
		s.GracefulShutdown()
	}()

	logger.Zap.Sugar().Infof("[grpcServer.RunGrpcServer] Writer gRPC server is listening on: %s:%d", s.config.Host, s.config.Port)

	err = s.server.Serve(l)
	if err != nil {
		logger.Zap.Sugar().Error(fmt.Sprintf("[grpcServer_RunGrpcServer.Serve] grpc server serve error: %+v", err))
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

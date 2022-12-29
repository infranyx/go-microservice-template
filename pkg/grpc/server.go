package grpc

import (
	"context"
	"fmt"
	"net"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcCtxTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	googleGrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpcErrorInterceptor "github.com/infranyx/go-grpc-template/pkg/grpc/interceptors/error_interceptor"
	grpcLoggerInterceptor "github.com/infranyx/go-grpc-template/pkg/grpc/interceptors/logger_interceptor"
	grpcSentryInterceptor "github.com/infranyx/go-grpc-template/pkg/grpc/interceptors/sentry_interceptor"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	sentryUtils "github.com/infranyx/go-grpc-template/pkg/sentry/sentry_utils"
)

type Server interface {
	RunGrpcServer(ctx context.Context, configGrpc func(grpcServer *googleGrpc.Server)) error
	GracefulShutdown()
	GetCurrentGrpcServer() *googleGrpc.Server
}

type Config struct {
	Port        int
	Host        string
	Development bool
}

type grpcServer struct {
	server *googleGrpc.Server
	config *Config
}

func NewGrpcServer(conf *Config) Server {
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

func (gs *grpcServer) RunGrpcServer(ctx context.Context, configGrpc func(grpcServer *googleGrpc.Server)) error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", gs.config.Port))
	if err != nil {
		return err
	}

	if configGrpc != nil {
		configGrpc(gs.server)
	}

	if gs.config.Development {
		reflection.Register(gs.server)
	}

	go func() {
		<-ctx.Done()
		logger.Zap.Sugar().Infof("App is shutting down Grpc PORT: {%d}", gs.config.Port)
		gs.GracefulShutdown()
	}()

	logger.Zap.Sugar().Infof("[grpcServer.RunGrpcServer] Writer gRPC server is listening on: %s:%d", gs.config.Host, gs.config.Port)

	err = gs.server.Serve(l)
	if err != nil {
		logger.Zap.Sugar().Error(fmt.Sprintf("[grpcServer_RunGrpcServer.Serve] grpc server serve error: %+v", err))
	}

	return err
}

func (gs *grpcServer) GetCurrentGrpcServer() *googleGrpc.Server {
	return gs.server
}

func (gs *grpcServer) GracefulShutdown() {
	gs.server.Stop()
	gs.server.GracefulStop()
}

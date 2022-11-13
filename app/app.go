package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/infranyx/go-grpc-template/config"
	const_app_env "github.com/infranyx/go-grpc-template/constant/app_env"
	article_configurator "github.com/infranyx/go-grpc-template/internal/article/configurator"
	"github.com/infranyx/go-grpc-template/pkg/grpc"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/infranyx/go-grpc-template/shared/infrastructure"
)

type Server struct {
	logger *logger.Logger
}

// Server constructor
func NewServer(logger *logger.Logger) *Server {
	return &Server{logger: logger}
}

func (s *Server) Run() error {
	var serverError error

	conf := config.Conf
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ic := infrastructure.NewInfrastructureConfigurator(s.logger, conf)
	infrastructureConfigurations, infraCleanup, err := ic.ConfigInfrastructures(ctx)
	if err != nil {
		return err
	}
	defer infraCleanup()

	grpcServerConfig := &grpc.GrpcConfig{
		Port:        conf.Grpc.Port,
		Host:        conf.Grpc.Host,
		Development: false,
	}
	if conf.App.AppEnv == const_app_env.DEV {
		grpcServerConfig.Development = true
	}

	grpcServer := grpc.NewGrpcServer(grpcServerConfig, s.logger)

	articleConfigurator := article_configurator.NewArticleControllerConfigurator(infrastructureConfigurations, grpcServer)
	articleConfigurator.ConfigureArticleController(ctx)

	go func() {
		if err := grpcServer.RunGrpcServer(ctx, nil); err != nil {
			s.logger.Errorf("(s.RunGrpcServer) err: {%v}", err)
			serverError = err
			cancel()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		s.logger.Errorf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		s.logger.Errorf("ctx.Done: %v", done)
	}

	s.logger.Info("Server Exited Properly")
	return serverError
}

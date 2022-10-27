package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/infranyx/go-grpc-template/config"
	const_app_env "github.com/infranyx/go-grpc-template/constant/app_env"
	"github.com/infranyx/go-grpc-template/pkg/grpc"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/infranyx/go-grpc-template/pkg/postgres"
)

type Server struct {
	logger *logger.Logger
}

// Server constructor
func NewServer(logger *logger.Logger) *Server {
	return &Server{logger: logger}
}

func (s *Server) Run() error {
	conf := config.Conf
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup DB
	postgresSQLX, err := postgres.NewPostgreSqlx()
	if err != nil {
		s.logger.Fatalf("Could not initialize Database connection using sqlx %s", err)
	}
	defer postgresSQLX.Close()

	var grpcServerConfig *grpc.GrpcConfig
	if conf.App.AppEnv == const_app_env.DEV {
		grpcServerConfig = &grpc.GrpcConfig{
			Port:        conf.Grpc.Port,
			Host:        conf.Grpc.Host,
			Development: true,
		}
	} else {
		grpcServerConfig = &grpc.GrpcConfig{
			Port:        conf.Grpc.Port,
			Host:        conf.Grpc.Host,
			Development: false,
		}
	}

	grpcServer := grpc.NewGrpcServer(grpcServerConfig, s.logger)
	grpcErr := grpcServer.RunGrpcServer(ctx, nil)
	if grpcErr != nil {
		s.logger.Error(grpcErr)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		s.logger.Errorf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		s.logger.Errorf("ctx.Done: %v", done)
	}

	s.logger.Info("Server Exited Properly")
	return nil
}

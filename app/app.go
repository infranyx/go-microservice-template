package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/infranyx/go-grpc-template/app/grpc"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup DB
	postgresSQLX, err := postgres.NewPostgreSqlx()
	if err != nil {
		log.Fatalf("Could not initialize Database connection using sqlx %s", err)
	}
	defer postgresSQLX.Close()

	grpcServer := grpc.NewGrpcServer(s.logger)
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

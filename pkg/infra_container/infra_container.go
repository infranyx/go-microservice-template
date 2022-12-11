package infraContainer

import (
	"context"
	"fmt"

	"github.com/infranyx/go-grpc-template/pkg/config"
	"github.com/infranyx/go-grpc-template/pkg/grpc"
	httpEcho "github.com/infranyx/go-grpc-template/pkg/http/echo"
	"github.com/infranyx/go-grpc-template/pkg/kafka"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/infranyx/go-grpc-template/pkg/postgres"
	"go.uber.org/zap"
)

type IContainer struct {
	GrpcServer grpc.GrpcServer // grpc.GrpcServer : Interface
	EchoServer httpEcho.EchoHttpServer
	Logger     *zap.Logger
	Cfg        *config.Config
	Pg         *postgres.Postgres
	Kafka      *kafka.Kafka
}

func NewIC(ctx context.Context) (*IContainer, func(), error) {
	var downFns []func()
	down := func() {
		for _, df := range downFns {
			df()
		}
	}

	grpcServerConfig := &grpc.GrpcConfig{
		Port:        config.Conf.Grpc.Port,
		Host:        config.Conf.Grpc.Host,
		Development: config.IsDevEnv(),
	}
	grpcServer := grpc.NewGrpcServer(grpcServerConfig)
	downFns = append(downFns, func() {
		grpcServer.GracefulShutdown()
	})

	echoServerConfig := &httpEcho.EchoHttpConfig{
		Port:        config.Conf.Http.Port,
		Development: config.IsDevEnv(),
	}
	echoServer := httpEcho.NewEchoHttpServer(echoServerConfig)
	downFns = append(downFns, func() {
		echoServer.GracefulShutdown(ctx)
	})

	pg, err := postgres.NewPgConn(ctx, &postgres.PgConf{
		Host:    config.Conf.Postgres.Host,
		Port:    config.Conf.Postgres.Port,
		User:    config.Conf.Postgres.User,
		Pass:    config.Conf.Postgres.Pass,
		DBName:  config.Conf.Postgres.DBName,
		SslMode: config.Conf.Postgres.SslMode,
	})
	if err != nil {
		return nil, down, fmt.Errorf("could not initialize database connection using sqlx %s", err)
	}
	downFns = append(downFns, func() {
		pg.Close()
	})

	ic := &IContainer{Cfg: config.Conf, Logger: logger.Zap, GrpcServer: grpcServer, EchoServer: echoServer, Pg: pg}

	return ic, down, nil
}

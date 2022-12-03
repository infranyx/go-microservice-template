package infra_container

import (
	"context"
	"fmt"

	"github.com/infranyx/go-grpc-template/config"
	"github.com/infranyx/go-grpc-template/pkg/grpc"
	"github.com/infranyx/go-grpc-template/pkg/kafka"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/infranyx/go-grpc-template/pkg/postgres"
	"go.uber.org/zap"
)

var IC *IContainer

type IContainer struct {
	GrpcServer grpc.GrpcServer // Interface
	Logger     *zap.Logger
	Cfg        *config.Config
	Pg         *postgres.Postgres
	Kafka      *kafka.Kafka
}

func NewIC(ctx context.Context) (*IContainer, func(), error) {
	cfg := config.Conf
	l := logger.Zap
	ic := &IContainer{Cfg: cfg, Logger: l}
	cleanup := []func(){}
	cleanupfn := func() {
		for _, c := range cleanup {
			c()
		}
	}

	grpcServerConfig := &grpc.GrpcConfig{
		Port:        ic.Cfg.Grpc.Port,
		Host:        ic.Cfg.Grpc.Host,
		Development: config.IsDevelopment(),
	}
	grpcServer := grpc.NewGrpcServer(grpcServerConfig)
	cleanup = append(cleanup, func() {
		grpcServer.GracefulShutdown()
	})
	ic.GrpcServer = grpcServer

	// Setup DB
	pg, err := postgres.NewPostgreSqlx(ctx, &postgres.Config{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		User:     cfg.Postgres.User,
		DBName:   cfg.Postgres.DBName,
		SSLMode:  cfg.Postgres.SSLMode,
		Password: cfg.Postgres.Password,
	})
	if err != nil {
		return nil, cleanupfn, fmt.Errorf("could not initialize database connection using sqlx %s", err)
	}
	cleanup = append(cleanup, func() {
		_ = pg.Sqlx.Close()
	})
	ic.Pg = pg

	// Setup Kafka
	// kafkaConn, err := kafka.NewKafkaConn(ctx, &kafka.Config{
	// 	Network: ic.cfg.Kafka.Network,
	// 	Address: ic.cfg.Kafka.Address,
	// })
	// if err != nil {
	// 	return nil, cleanupfn, err
	// }
	// cleanup = append(cleanup, func() {
	// 	kafkaConn.Conn.Close()
	// })
	// // TODO
	// // kafkaConn.CreateTopic()
	// ic.Kafka = kafkaConn

	return ic, cleanupfn, nil
}

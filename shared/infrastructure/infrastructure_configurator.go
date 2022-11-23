package infrastructure

import (
	"context"
	"fmt"

	"github.com/infranyx/go-grpc-template/config"
	"github.com/infranyx/go-grpc-template/pkg/kafka"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/infranyx/go-grpc-template/pkg/postgres"
)

type InfrastructureConfiguration struct {
	Logger *logger.Logger
	Cfg    *config.Config
	Pgx    *postgres.Postgres
	Kafka  *kafka.Kafka
}

type InfrastructureConfigurator interface {
	ConfigureInfrastructure() error
}

type infrastructureConfigurator struct {
	logger *logger.Logger
	cfg    *config.Config
}

func NewInfrastructureConfigurator(logger *logger.Logger, cfg *config.Config) *infrastructureConfigurator {
	return &infrastructureConfigurator{logger: logger, cfg: cfg}
}

func (ic *infrastructureConfigurator) ConfigInfrastructures(ctx context.Context) (*InfrastructureConfiguration, func(), error) {
	infrastructure := &InfrastructureConfiguration{Cfg: ic.cfg, Logger: ic.logger}

	cleanup := []func(){}
	cleanupfn := func() {
		for _, c := range cleanup {
			c()
		}
	}

	// Setup DB
	pgx, err := postgres.NewPostgreSqlx(ctx, &postgres.Config{
		Host:     ic.cfg.Postgres.Host,
		Port:     ic.cfg.Postgres.Port,
		User:     ic.cfg.Postgres.User,
		DBName:   ic.cfg.Postgres.DBName,
		SSLMode:  ic.cfg.Postgres.SSLMode,
		Password: ic.cfg.Postgres.Password,
	})
	if err != nil {
		return nil, cleanupfn, fmt.Errorf("could not initialize database connection using sqlx %s", err)
	}
	cleanup = append(cleanup, func() {
		pgx.Sqlx.Close()
	})
	infrastructure.Pgx = pgx

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
	// infrastructure.Kafka = kafkaConn

	return infrastructure, cleanupfn, nil
}

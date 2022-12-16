package infraContainer

import (
	"context"
	"fmt"

	"github.com/infranyx/go-grpc-template/pkg/config"
	"github.com/infranyx/go-grpc-template/pkg/grpc"
	httpEcho "github.com/infranyx/go-grpc-template/pkg/http/echo"
	kafkaConsumer "github.com/infranyx/go-grpc-template/pkg/kafka/consumer"
	kafkaProducer "github.com/infranyx/go-grpc-template/pkg/kafka/producer"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/infranyx/go-grpc-template/pkg/postgres"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type IContainer struct {
	GrpcServer  grpc.GrpcServer
	EchoServer  httpEcho.EchoHttpServer
	Logger      *zap.Logger
	Cfg         *config.Config
	Pg          *postgres.Postgres
	KafkaWriter *kafkaProducer.Writer
	KafkaReader *kafkaConsumer.Reader
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
		BasePath:    "/api/v1",
	}
	echoServer := httpEcho.NewEchoHttpServer(echoServerConfig)
	echoServer.SetupDefaultMiddlewares()
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

	kwc := &kafkaProducer.WriterConf{
		Brokers:      config.Conf.Kafka.ClientBrokers,
		Topic:        config.Conf.Kafka.Topic,
		RequiredAcks: kafka.RequireAll,
	}
	kw := kafkaProducer.NewKafkaWriter(kwc)
	downFns = append(downFns, func() {
		kw.Client.Close()
	})

	krc := &kafkaConsumer.ReaderConf{
		Brokers: config.Conf.Kafka.ClientBrokers,
		Topic:   config.Conf.Kafka.Topic,
		GroupID: config.Conf.Kafka.ClientGroupId,
	}
	kr := kafkaConsumer.NewKafkaReader(krc)
	downFns = append(downFns, func() {
		kr.Client.Close()
	})

	ic := &IContainer{Cfg: config.Conf, Logger: logger.Zap, GrpcServer: grpcServer, EchoServer: echoServer, Pg: pg, KafkaWriter: kw, KafkaReader: kr}

	return ic, down, nil
}

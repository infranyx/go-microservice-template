package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Config struct {
	Network string
	Address string
}

type Kafka struct {
	Conn *kafka.Conn
}

func NewKafkaConn(ctx context.Context, conf *Config) (*Kafka, error) {
	conn, err := kafka.DialContext(ctx, conf.Network, conf.Address)
	if err != nil {
		return nil, err
	}

	return &Kafka{
		Conn: conn,
	}, nil
}

package kafka

import (
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/segmentio/kafka-go"
)

type KafkaReader struct {
	Client *kafka.Reader
}

func NewKafkaReader(rc kafka.ReaderConfig) *KafkaReader {
	return &KafkaReader{
		Client: kafka.NewReader(rc),
	}
}

func NewKafkaReaderConfig() kafka.ReaderConfig {
	return kafka.ReaderConfig{
		// Brokers:   []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		// GroupID:                groupID,
		// Topic:                  topic,
		QueueCapacity:          queueCapacity,
		MinBytes:               minBytes,
		MaxBytes:               maxBytes,
		HeartbeatInterval:      heartbeatInterval,
		CommitInterval:         commitInterval,
		PartitionWatchInterval: partitionWatchInterval,
		Logger:                 kafka.LoggerFunc(logger.Zap.Sugar().Debugf),
		ErrorLogger:            kafka.LoggerFunc(logger.Zap.Sugar().Errorf),
		MaxAttempts:            maxAttempts,
		Dialer: &kafka.Dialer{
			Timeout: dialTimeout,
		},
	}
}

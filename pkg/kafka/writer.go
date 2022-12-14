package kafka

import (
	"time"

	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/segmentio/kafka-go"
)

const (
	minBytes               = 10e3 // 10KB
	maxBytes               = 10e6 // 10MB
	queueCapacity          = 100
	heartbeatInterval      = 3 * time.Second
	commitInterval         = 0
	partitionWatchInterval = 5 * time.Second
	maxAttempts            = 3
	dialTimeout            = 3 * time.Minute

	writerReadTimeout  = 10 * time.Second
	writerWriteTimeout = 10 * time.Second
	writerRequiredAcks = -1
	writerMaxAttempts  = 3
)

type KafkaWriter struct {
	Client *kafka.Writer
}

func NewKafkaWriter(wc kafka.WriterConfig) *KafkaWriter {
	return &KafkaWriter{
		Client: kafka.NewWriter(wc),
	}
}

func NewKafkaWriterConfig() kafka.WriterConfig {
	return kafka.WriterConfig{
		// Brokers:      []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		// Topic:        "topic-A",
		QueueCapacity: queueCapacity,
		Balancer:      &kafka.LeastBytes{},
		RequiredAcks:  writerRequiredAcks,
		MaxAttempts:   writerMaxAttempts,
		Logger:        kafka.LoggerFunc(logger.Zap.Sugar().Debugf),
		ErrorLogger:   kafka.LoggerFunc(logger.Zap.Sugar().Errorf),
		ReadTimeout:   writerReadTimeout,
		WriteTimeout:  writerWriteTimeout,
	}
}

package kafkaProducer

import (
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
)

type Writer struct {
	Client *kafka.Writer
}

type WriterConfig struct {
	Brokers      []string
	Topic        string
	RequiredAcks kafka.RequiredAcks
}

func NewKafkaWriter(cfg *WriterConfig) *Writer {
	kafkaWriterConfig := &kafka.Writer{
		Addr:         kafka.TCP(cfg.Brokers...),
		Topic:        cfg.Topic,
		RequiredAcks: cfg.RequiredAcks,
		Balancer:     &kafka.LeastBytes{},
		Compression:  compress.Snappy,
		Logger:       kafka.LoggerFunc(logger.Zap.Sugar().Infof),
		ErrorLogger:  kafka.LoggerFunc(logger.Zap.Sugar().Errorf),
	}
	return &Writer{
		Client: kafkaWriterConfig,
	}
}

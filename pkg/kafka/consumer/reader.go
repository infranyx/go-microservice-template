package kafkaConsumer

import (
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/segmentio/kafka-go"
)

type Reader struct {
	Client *kafka.Reader
}

type ReaderConf struct {
	Brokers []string
	GroupID string
	Topic   string
}

func NewKafkaReader(cfg *ReaderConf) *Reader {
	rc := kafka.ReaderConfig{
		Brokers:     cfg.Brokers,
		GroupID:     cfg.GroupID,
		Topic:       cfg.Topic,
		Logger:      kafka.LoggerFunc(logger.Zap.Sugar().Infof),
		ErrorLogger: kafka.LoggerFunc(logger.Zap.Sugar().Errorf),
	}
	return &Reader{
		Client: kafka.NewReader(rc),
	}
}

package kafka

import (
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/segmentio/kafka-go/compress"
	"time"

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

type Writer struct {
	Client *kafka.Writer
}

type WriterConf struct {
	Brokers []string
	Topic   string
}

func NewKafkaWriter(cfg *WriterConf) *Writer {
	w := &kafka.Writer{
		Addr:         kafka.TCP(cfg.Brokers...),
		Topic:        cfg.Topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: writerRequiredAcks,
		MaxAttempts:  writerMaxAttempts,
		Logger:       kafka.LoggerFunc(logger.Zap.Sugar().Infof),
		ErrorLogger:  kafka.LoggerFunc(logger.Zap.Sugar().Errorf),
		Compression:  compress.Snappy,
		ReadTimeout:  writerReadTimeout,
		WriteTimeout: writerWriteTimeout,
	}
	return &Writer{
		Client: w,
	}
}

//func NewKafkaWriterConfig() *KafkaWriter {
//	w := &kafka.Writer{
//		Addr:         kafka.TCP(pcg.Brokers...),
//		Topic:        topic,
//		Balancer:     &kafka.LeastBytes{},
//		RequiredAcks: writerRequiredAcks,
//		MaxAttempts:  writerMaxAttempts,
//		Logger:       kafka.LoggerFunc(pcg.log.Debugf),
//		ErrorLogger:  kafka.LoggerFunc(pcg.log.Errorf),
//		Compression:  compress.Snappy,
//		ReadTimeout:  writerReadTimeout,
//		WriteTimeout: writerWriteTimeout,
//	}
//	return w
//	//return kafka.WriterConfig{
//	//	// Brokers:      []string{"localhost:9092", "localhost:9093", "localhost:9094"},
//	//	// Topic:        "topic-A",
//	//	QueueCapacity: queueCapacity,
//	//	Balancer:      &kafka.LeastBytes{},
//	//	RequiredAcks:  writerRequiredAcks,
//	//	MaxAttempts:   writerMaxAttempts,
//	//	Logger:        kafka.LoggerFunc(logger.Zap.Sugar().Debugf),
//	//	ErrorLogger:   kafka.LoggerFunc(logger.Zap.Sugar().Errorf),
//	//	ReadTimeout:   writerReadTimeout,
//	//	WriteTimeout:  writerWriteTimeout,
//	//}
//}

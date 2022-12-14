package articleKafka

import (
	"context"

	"github.com/infranyx/go-grpc-template/pkg/kafka"
	kafkaGo "github.com/segmentio/kafka-go"
)

type articleProducer struct {
	createWriter *kafka.KafkaWriter
}

func NewArticleProducer(w *kafka.KafkaWriter) *articleProducer {
	return &articleProducer{createWriter: w}
}

func (p *articleProducer) PublishCreate(ctx context.Context, msgs ...kafkaGo.Message) error {
	return p.createWriter.Client.WriteMessages(ctx, msgs...)
}

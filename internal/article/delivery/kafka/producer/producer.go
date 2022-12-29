package articleKafkaProducer

import (
	"context"

	"github.com/segmentio/kafka-go"

	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	kafkaProducer "github.com/infranyx/go-grpc-template/pkg/kafka/producer"
)

type producer struct {
	createWriter *kafkaProducer.Writer
}

func NewProducer(w *kafkaProducer.Writer) articleDomain.KafkaProducer {
	return &producer{createWriter: w}
}

func (p *producer) PublishCreateEvent(ctx context.Context, messages ...kafka.Message) error {
	return p.createWriter.Client.WriteMessages(ctx, messages...)
}

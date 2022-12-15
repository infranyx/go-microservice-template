package articleKafkaProducer

import (
	"context"

	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	"github.com/infranyx/go-grpc-template/pkg/kafka"
	kafkaGo "github.com/segmentio/kafka-go"
)

type articleProducer struct {
	createWriter *kafka.Writer
}

func NewArticleProducer(w *kafka.Writer) articleDomain.ArticleProducer {
	return &articleProducer{createWriter: w}
}

func (p *articleProducer) PublishCreate(ctx context.Context, msgs ...kafkaGo.Message) error {
	return p.createWriter.Client.WriteMessages(ctx, msgs...)
}

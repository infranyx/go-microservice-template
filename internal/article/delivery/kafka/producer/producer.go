package articleKafkaProducer

import (
	"context"

	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	kafkaProducer "github.com/infranyx/go-grpc-template/pkg/kafka/producer"
	kafkaGo "github.com/segmentio/kafka-go"
)

type articleProducer struct {
	createWriter *kafkaProducer.Writer
}

func NewArticleProducer(w *kafkaProducer.Writer) articleDomain.ArticleProducer {
	return &articleProducer{createWriter: w}
}

func (p *articleProducer) PublishCreate(ctx context.Context, msgs ...kafkaGo.Message) error {
	return p.createWriter.Client.WriteMessages(ctx, msgs...)
}

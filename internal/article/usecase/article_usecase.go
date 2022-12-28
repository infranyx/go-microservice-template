package articleUseCase

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"

	sampleExtServiceDomain "github.com/infranyx/go-grpc-template/external/sample_ext_service/domain"
	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
)

type useCase struct {
	repository              articleDomain.Repository
	kafkaProducer           articleDomain.KafkaProducer
	sampleExtServiceUseCase sampleExtServiceDomain.SampleExtServiceUseCase
}

func NewUseCase(
	repository articleDomain.Repository,
	sampleExtServiceUseCase sampleExtServiceDomain.SampleExtServiceUseCase,
	kafkaProducer articleDomain.KafkaProducer,
) articleDomain.UseCase {
	return &useCase{
		repository:              repository,
		kafkaProducer:           kafkaProducer,
		sampleExtServiceUseCase: sampleExtServiceUseCase,
	}
}

func (uc *useCase) CreateArticle(ctx context.Context, req *articleDto.CreateArticleRequestDto) (*articleDto.CreateArticleResponseDto, error) {
	article, err := uc.repository.CreateArticle(ctx, req)
	if err != nil {
		return nil, err
	}

	// TODO : if err => return Marshal_Err_Exception
	jsonArticle, _ := json.Marshal(article)

	// if it has go keyword and if we pass the request context to it, it will terminate after request lifecycle.
	_ = uc.kafkaProducer.PublishCreateEvent(context.Background(), kafka.Message{
		Key:   []byte("Article"),
		Value: jsonArticle,
	})

	return article, err
}

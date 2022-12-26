package articleUseCase

import (
	"context"
	"encoding/json"

	sampleExtServiceDomain "github.com/infranyx/go-grpc-template/external/sample_ext_service/domain"
	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
	"github.com/segmentio/kafka-go"
)

type useCase struct {
	repository              articleDomain.ArticleRepository
	kafkaProducer           articleDomain.ArticleProducer
	sampleExtServiceUseCase sampleExtServiceDomain.SampleExtServiceUseCase
}

func NewUseCase(
	repository articleDomain.ArticleRepository,
	esu sampleExtServiceDomain.SampleExtServiceUseCase,
	kp articleDomain.ArticleProducer,
) articleDomain.ArticleUseCase {
	return &useCase{
		repository:              repository,
		kafkaProducer:           kp,
		sampleExtServiceUseCase: esu,
	}
}

func (uc *useCase) Create(ctx context.Context, req *articleDto.CreateArticleRequestDto) (*articleDto.CreateArticleResponseDto, error) {
	article, err := uc.repository.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	// TODO : if err => return Marshal_Err_Exception
	jsonArticle, _ := json.Marshal(article)

	//it has go keyword so if we pass the request context to it, it will terminate after request lifecycle.
	go uc.kafkaProducer.PublishCreateEvent(context.Background(), kafka.Message{
		Key:   []byte("Article"),
		Value: jsonArticle,
	})
	return article, err
}

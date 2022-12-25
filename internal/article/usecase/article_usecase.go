package articleUseCase

import (
	"context"
	"encoding/json"

	sampleExtServiceDomain "github.com/infranyx/go-grpc-template/external/sample_ext_service/domain"
	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
	"github.com/segmentio/kafka-go"
)

type articleUseCase struct {
	articleRepo             articleDomain.ArticleRepository
	sampleExtServiceUseCase sampleExtServiceDomain.SampleExtServiceUseCase
	articleProducer         articleDomain.ArticleProducer
}

func NewArticleUseCase(
	articleRepo articleDomain.ArticleRepository,
	esu sampleExtServiceDomain.SampleExtServiceUseCase,
	ap articleDomain.ArticleProducer,
) articleDomain.ArticleUseCase {
	return &articleUseCase{
		articleRepo:             articleRepo,
		sampleExtServiceUseCase: esu,
		articleProducer:         ap,
	}
}

func (au *articleUseCase) Create(ctx context.Context, article *articleDto.CreateArticle) (*articleDomain.Article, error) {
	c := context.Background()
	result, err := au.articleRepo.Create(ctx, article)
	if err != nil {
		return nil, err
	}
	b, _ := json.Marshal(result)

	//it has go keyword so if we pass the request context to it, it will terminate after request lifecycle.
	go au.articleProducer.PublishCreate(c, kafka.Message{
		Key:   []byte("Article"),
		Value: b,
	})
	return result, err
}

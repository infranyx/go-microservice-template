package articleDomain

import (
	"context"

	"github.com/google/uuid"

	articleV1 "github.com/infranyx/protobuf-template-go/golang_template/article/v1"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"

	articleDto "github.com/infranyx/go-microservice-template/internal/article/dto"
)

type Article struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"desc"`
}

type Configurator interface {
	Configure(ctx context.Context) error
}

type UseCase interface {
	CreateArticle(ctx context.Context, article *articleDto.CreateArticleRequestDto) (*articleDto.CreateArticleResponseDto, error)
}

type Repository interface {
	CreateArticle(ctx context.Context, article *articleDto.CreateArticleRequestDto) (*articleDto.CreateArticleResponseDto, error)
}

type GrpcController interface {
	CreateArticle(ctx context.Context, req *articleV1.CreateArticleRequest) (*articleV1.CreateArticleResponse, error)
	GetArticleById(ctx context.Context, req *articleV1.GetArticleByIdRequest) (*articleV1.GetArticleByIdResponse, error)
}

type HttpController interface {
	CreateArticle(c echo.Context) error
}

type Job interface {
	StartJobs(ctx context.Context)
}

type KafkaProducer interface {
	PublishCreateEvent(ctx context.Context, messages ...kafka.Message) error
}

type KafkaConsumer interface {
	RunConsumers(ctx context.Context)
}

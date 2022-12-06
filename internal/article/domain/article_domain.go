package articleDomain

import (
	"context"
	"github.com/google/uuid"
	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
	articleV1 "github.com/infranyx/protobuf-template-go/golang-grpc-template/article/v1"
)

type Article struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}

type ArticleConfigurator interface {
	ConfigureArticle(ctx context.Context) error
}

type ArticleGrpcController interface {
	CreateArticle(ctx context.Context, req *articleV1.CreateArticleRequest) (*articleV1.CreateArticleResponse, error)
	GetArticleById(ctx context.Context, req *articleV1.GetArticleByIdRequest) (*articleV1.GetArticleByIdResponse, error)
}

type ArticleUseCase interface {
	Create(ctx context.Context, article *articleDto.CreateArticle) (*Article, error)
}

type ArticleRepository interface {
	Create(ctx context.Context, article *articleDto.CreateArticle) (*Article, error)
}

//type ArticleHttpController interface {
//	Create(response http.ResponseWriter, request *http.Request)
//}

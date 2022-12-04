package articleDomain

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
)

type Article struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}

type ArticleController interface {
	Create(response http.ResponseWriter, request *http.Request)
}

type ArticleUseCase interface {
	Create(ctx context.Context, article *articleDto.CreateArticle) (*Article, error)
}

type ArticleRepository interface {
	Create(ctx context.Context, article *articleDto.CreateArticle) (*Article, error)
}

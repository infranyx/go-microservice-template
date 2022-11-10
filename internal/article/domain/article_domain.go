package article_domain

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	article_dto "github.com/infranyx/go-grpc-template/internal/article/dto"
)

// Item represents a Item for all sub domains
type Article struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}

// ArticleService is a contract of http adapter layer
type ArticleController interface {
	Create(response http.ResponseWriter, request *http.Request)
}

// ArticleUseCase is a contract of business rule layer
type ArticleUseCase interface {
	Create(ctx context.Context, article *article_dto.CreateArticle) (*Article, error)
}

// ArticleRepository is a contract of database connection adapter layer
type ArticleRepository interface {
	Create(ctx context.Context, article *article_dto.CreateArticle) (*Article, error)
}

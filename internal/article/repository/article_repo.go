package article_repo

import (
	"context"
	"fmt"

	article_domain "github.com/infranyx/go-grpc-template/internal/article/domain"
	article_dto "github.com/infranyx/go-grpc-template/internal/article/dto"
	"github.com/infranyx/go-grpc-template/pkg/postgres"
)

type articleRepository struct {
	Conn *postgres.Postgres
}

func NewArticleRepository(Conn *postgres.Postgres) article_domain.ArticleRepository {
	return &articleRepository{Conn}
}

func (r *articleRepository) Create(ctx context.Context, article *article_dto.CreateArticle) (*article_domain.Article, error) {
	query := `INSERT INTO articles (name, description) VALUES ($1, $2) RETURNING id, name, description`

	var result article_domain.Article
	x, err := r.Conn.Sqlx.QueryContext(ctx, query, article.Name, article.Description)

	if err != nil {
		return &article_domain.Article{}, fmt.Errorf("error inserting article record")
	}

	for x.Next() {
		err = x.Scan(&result.ID, &result.Name, &result.Description)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

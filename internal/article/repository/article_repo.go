package articleRepo

import (
	"context"
	"fmt"

	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
	"github.com/infranyx/go-grpc-template/pkg/postgres"
)

type articleRepository struct {
	Conn *postgres.Postgres
}

func NewArticleRepository(Conn *postgres.Postgres) articleDomain.ArticleRepository {
	return &articleRepository{Conn}
}

func (ar *articleRepository) Create(ctx context.Context, article *articleDto.CreateArticle) (*articleDomain.Article, error) {
	query := `INSERT INTO articles (name, description) VALUES ($1, $2) RETURNING id, name, description`

	var result articleDomain.Article
	x, err := ar.Conn.SqlxDB.QueryContext(ctx, query, article.Name, article.Description)

	if err != nil {
		return &articleDomain.Article{}, fmt.Errorf("error inserting article record")
	}

	for x.Next() {
		err = x.Scan(&result.ID, &result.Name, &result.Description)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

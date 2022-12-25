package articleRepository

import (
	"context"
	"fmt"

	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
	"github.com/infranyx/go-grpc-template/pkg/postgres"
)

type repository struct {
	conn *postgres.Postgres
}

func NewRepository(Conn *postgres.Postgres) articleDomain.ArticleRepository {
	return &repository{Conn}
}

func (r *repository) Create(ctx context.Context, entity *articleDto.CreateArticle) (*articleDomain.Article, error) {
	query := `INSERT INTO articles (name, description) VALUES ($1, $2) RETURNING id, name, description`

	res, err := r.conn.SqlxDB.QueryContext(ctx, query, entity.Name, entity.Description)
	if err != nil {
		return nil, fmt.Errorf("error inserting article record")
	}

	result := new(articleDomain.Article)
	for res.Next() {
		err = res.Scan(&result.ID, &result.Name, &result.Description)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

package articleRepository

import (
	"context"
	"fmt"

	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
	"github.com/infranyx/go-grpc-template/pkg/postgres"
)

type repository struct {
	postgres *postgres.Postgres
}

func NewRepository(conn *postgres.Postgres) articleDomain.Repository {
	return &repository{postgres: conn}
}

func (rp *repository) CreateArticle(
	ctx context.Context,
	entity *articleDto.CreateArticleRequestDto,
) (*articleDto.CreateArticleResponseDto, error) {
	query := `INSERT INTO articles (name, description) VALUES ($1, $2) RETURNING id, name, description`

	result, err := rp.postgres.SqlxDB.QueryContext(ctx, query, entity.Name, entity.Description)
	if err != nil {
		return nil, fmt.Errorf("error inserting article record")
	}

	article := new(articleDto.CreateArticleResponseDto)
	for result.Next() {
		err = result.Scan(&article.ID, &article.Name, &article.Description)
		if err != nil {
			return nil, err
		}
	}

	return article, nil
}

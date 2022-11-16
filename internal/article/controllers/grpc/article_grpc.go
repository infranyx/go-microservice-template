package article_grpc

import (
	"context"

	article_domain "github.com/infranyx/go-grpc-template/internal/article/domain"
	article_dto "github.com/infranyx/go-grpc-template/internal/article/dto"
	articlev1 "go.buf.build/grpc/go/infranyx/golang-grpc-template/article/v1"
)

type ArticleController struct {
	articleUC article_domain.ArticleUseCase
}

func New(usecase article_domain.ArticleUseCase) *ArticleController {
	return &ArticleController{
		articleUC: usecase,
	}
}

func (c *ArticleController) CreateArticle(ctx context.Context, req *articlev1.CreateArticleRequest) (*articlev1.CreateArticleResponse, error) {
	article, err := c.articleUC.Create(ctx, &article_dto.CreateArticle{
		Name:        req.Name,
		Description: req.Desc,
	})
	if err != nil {
		return nil, err
	}
	return &articlev1.CreateArticleResponse{
		Id:   article.ID.String(),
		Name: article.Name,
		Desc: article.Description,
	}, nil
}

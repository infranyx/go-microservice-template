package article_grpc

import (
	"context"

	article_domain "github.com/infranyx/go-grpc-template/internal/article/domain"
	article_dto "github.com/infranyx/go-grpc-template/internal/article/dto"
	article_exception "github.com/infranyx/go-grpc-template/internal/article/exception"
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
	articleDto := &article_dto.CreateArticle{
		Name:        req.Name,
		Description: req.Desc,
	}
	err := articleDto.ValidateCreateArticleDto()
	if err != nil {
		return nil, article_exception.NewCreateArticleValidationErr(err)
	}
	article, err := c.articleUC.Create(ctx, articleDto)
	if err != nil {
		return nil, err
	}
	return &articlev1.CreateArticleResponse{
		Id:   article.ID.String(),
		Name: article.Name,
		Desc: article.Description,
	}, nil
}

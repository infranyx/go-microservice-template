package articleGrpc

import (
	"context"

	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
	articleValidationException "github.com/infranyx/go-grpc-template/internal/article/exception"
	articleV1 "github.com/infranyx/protobuf-template-go/golang-grpc-template/article/v1"
)

type ArticleController struct {
	articleUC articleDomain.ArticleUseCase
}

func New(uc articleDomain.ArticleUseCase) *ArticleController {
	return &ArticleController{
		articleUC: uc,
	}
}

func (c *ArticleController) CreateArticle(ctx context.Context, req *articleV1.CreateArticleRequest) (*articleV1.CreateArticleResponse, error) {
	articleDto := &articleDto.CreateArticle{
		Name:        req.Name,
		Description: req.Desc,
	}
	err := articleDto.ValidateCreateArticleDto()
	if err != nil {
		return nil, articleValidationException.NewCreateArticleValidationErr(err)
	}
	article, err := c.articleUC.Create(ctx, articleDto)
	if err != nil {
		return nil, err
	}
	return &articleV1.CreateArticleResponse{
		Id:   article.ID.String(),
		Name: article.Name,
		Desc: article.Description,
	}, nil
}

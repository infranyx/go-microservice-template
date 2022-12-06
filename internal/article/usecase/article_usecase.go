package articleUseCase

import (
	"context"

	goTemplateDomain "github.com/infranyx/go-grpc-template/external/go_template/domain"
	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
)

type articleUseCase struct {
	articleRepo       articleDomain.ArticleRepository
	goTemplateUseCase goTemplateDomain.GoTemplateUseCase
}

func NewArticleUseCase(articleRepo articleDomain.ArticleRepository, gtu goTemplateDomain.GoTemplateUseCase) articleDomain.ArticleUseCase {
	return &articleUseCase{
		articleRepo:       articleRepo,
		goTemplateUseCase: gtu,
	}
}

func (au *articleUseCase) Create(ctx context.Context, article *articleDto.CreateArticle) (*articleDomain.Article, error) {
	result, err := au.articleRepo.Create(ctx, article)
	return result, err
}

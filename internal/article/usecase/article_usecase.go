package articleUseCase

import (
	"context"

	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
)

type articleUseCase struct {
	articleRepo articleDomain.ArticleRepository
}

func NewArticleUseCase(articleRepo articleDomain.ArticleRepository) articleDomain.ArticleUseCase {
	return &articleUseCase{
		articleRepo: articleRepo,
	}
}

func (u *articleUseCase) Create(ctx context.Context, article *articleDto.CreateArticle) (*articleDomain.Article, error) {
	result, err := u.articleRepo.Create(ctx, article)
	return result, err
}

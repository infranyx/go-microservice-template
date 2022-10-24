package article_usecases

import (
	"context"

	article_domain "github.com/infranyx/go-grpc-template/internal/article/domain"
	article_dto "github.com/infranyx/go-grpc-template/internal/article/dto"
)

type articleUseCase struct {
	articleRepo article_domain.ArticleRepository
}

func NewArticleUseCase(articleRepo article_domain.ArticleRepository) article_domain.ArticleUseCase {
	return &articleUseCase{
		articleRepo: articleRepo,
	}
}

func (u *articleUseCase) Create(ctx context.Context, article *article_dto.CreateArticle) (*article_domain.Article, error) {
	result, err := u.articleRepo.Create(ctx, article)
	return result, err
}

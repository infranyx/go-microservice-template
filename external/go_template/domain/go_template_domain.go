package goTemplateDomain

import (
	"context"
	articleV1 "github.com/infranyx/protobuf-template-go/golang-grpc-template/article/v1"
)

type GoTemplateUseCase interface {
	CreateArticle(ctx context.Context, in *articleV1.CreateArticleRequest) (*articleV1.CreateArticleResponse, error)
}

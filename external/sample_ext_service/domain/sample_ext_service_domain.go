package sampleExtServiceDomain

import (
	"context"

	articleV1 "github.com/infranyx/protobuf-template-go/golang-grpc-template/article/v1"
)

type SampleExtServiceUseCase interface {
	CreateArticle(ctx context.Context, req *articleV1.CreateArticleRequest) (*articleV1.CreateArticleResponse, error)
}

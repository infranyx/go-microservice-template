package goTemplateUseCase

import (
	"context"

	goTemplateDomain "github.com/infranyx/go-grpc-template/external/go_template/domain"
	grpcError "github.com/infranyx/go-grpc-template/pkg/error/grpc"
	"github.com/infranyx/go-grpc-template/pkg/grpc"
	articleV1 "github.com/infranyx/protobuf-template-go/golang-grpc-template/article/v1"
)

type goTemplateUseCase struct {
	grpcClient grpc.GrpcClient
}

func NewGoTemplateUseCase(grpcClient grpc.GrpcClient) goTemplateDomain.GoTemplateUseCase {
	return &goTemplateUseCase{
		grpcClient: grpcClient,
	}
}

func (gtu *goTemplateUseCase) CreateArticle(ctx context.Context, in *articleV1.CreateArticleRequest) (*articleV1.CreateArticleResponse, error) {
	articleGrpcClient := articleV1.NewArticleServiceClient(gtu.grpcClient.GetGrpcConnection())
	res, err := articleGrpcClient.CreateArticle(ctx, in)
	if err != nil {
		return nil, grpcError.ParseExternalGrpcErr(err)
	}
	return res, nil
}

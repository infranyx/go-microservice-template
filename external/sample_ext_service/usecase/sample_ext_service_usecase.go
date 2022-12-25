package sampleExtServiceUseCase

import (
	"context"

	sampleExtServiceDomain "github.com/infranyx/go-grpc-template/external/sample_ext_service/domain"
	grpcError "github.com/infranyx/go-grpc-template/pkg/error/grpc"
	"github.com/infranyx/go-grpc-template/pkg/grpc"
	articleV1 "github.com/infranyx/protobuf-template-go/golang-grpc-template/article/v1"
)

type sampleExtServiceUseCase struct {
	grpcClient grpc.GrpcClient
}

func NewSampleExtServiceUseCase(grpcClient grpc.GrpcClient) sampleExtServiceDomain.SampleExtServiceUseCase {
	return &sampleExtServiceUseCase{
		grpcClient: grpcClient,
	}
}

func (esu *sampleExtServiceUseCase) CreateArticle(ctx context.Context, req *articleV1.CreateArticleRequest) (*articleV1.CreateArticleResponse, error) {
	articleGrpcClient := articleV1.NewArticleServiceClient(esu.grpcClient.GetGrpcConnection())

	res, err := articleGrpcClient.CreateArticle(ctx, req)
	if err != nil {
		return nil, grpcError.ParseExternalGrpcErr(err)
	}

	return res, nil
}

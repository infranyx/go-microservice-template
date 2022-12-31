package sampleExtServiceUseCase

import (
	"context"

	articleV1 "github.com/infranyx/protobuf-template-go/golang_template/article/v1"

	sampleExtServiceDomain "github.com/infranyx/go-microservice-template/external/sample_ext_service/domain"
	grpcError "github.com/infranyx/go-microservice-template/pkg/error/grpc"
	"github.com/infranyx/go-microservice-template/pkg/grpc"
)

type sampleExtServiceUseCase struct {
	grpcClient grpc.Client
}

func NewSampleExtServiceUseCase(grpcClient grpc.Client) sampleExtServiceDomain.SampleExtServiceUseCase {
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

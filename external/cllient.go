package main

import (
	"context"
	"fmt"

	"github.com/infranyx/go-grpc-template/pkg/grpc"
	grpcErrors "github.com/infranyx/go-grpc-template/shared/error/grpc"
	articlev1 "github.com/infranyx/protobuf-template-go/golang-grpc-template/article/v1"
)

func main() {
	// TODO : Update external client samples.
	ctx := context.Background()
	c, _ := grpc.NewGrpcClient(&grpc.GrpcConfig{Port: 3000, Host: "0.0.0.0"})

	articleGrpcClient := articlev1.NewArticleServiceClient(c.GetGrpcConnection())
	res, err := articleGrpcClient.CreateArticle(ctx, &articlev1.CreateArticleRequest{
		Name: "te",
		Desc: "re",
	})
	fmt.Println(res)
	// fmt.Printf("%+v\n", err)

	e := grpcErrors.ParseExternalGrpcErr(err)
	fmt.Println(e.GetDetails())
	//fmt.Println(*e.GetDetails()[0])
}

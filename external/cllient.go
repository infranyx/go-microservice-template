package main

import (
	"context"
	"fmt"

	grpcErrors "github.com/infranyx/go-grpc-template/pkg/error/grpc"
	"github.com/infranyx/go-grpc-template/pkg/grpc"
	articleV1 "github.com/infranyx/protobuf-template-go/golang-grpc-template/article/v1"
)

func main() {
	// TODO : Update external client samples.
	ctx := context.Background()
	c, _ := grpc.NewGrpcClient(&grpc.GrpcConfig{Port: 3000, Host: "0.0.0.0"})

	articleGrpcClient := articleV1.NewArticleServiceClient(c.GetGrpcConnection())
	res, err := articleGrpcClient.CreateArticle(ctx, &articleV1.CreateArticleRequest{
		Name: "te",
		Desc: "re",
	})
	fmt.Println(res)
	// fmt.Printf("%+v\n", err)

	e := grpcErrors.ParseExternalGrpcErr(err)
	fmt.Println(e.GetDetails())
	//fmt.Println(*e.GetDetails()[0])
}

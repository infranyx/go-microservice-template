package job

import (
	"context"
	"fmt"

	"github.com/infranyx/go-grpc-template/pkg/wrapper"
)

func (aj *articleJob) logArticleWorker() wrapper.HandlerFunc {
	return func(ctx context.Context, args ...interface{}) (interface{}, error) {
		fmt.Println("article log job")
		return nil, nil
	}
}

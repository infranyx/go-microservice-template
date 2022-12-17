package job

import (
	"context"
	"fmt"

	errorUtils "github.com/infranyx/go-grpc-template/pkg/error/error_utils"
)

func (aj *articleJob) logArticleWorker(
	ctx context.Context,
) {
	errorUtils.HandlerErrorWrapper(
		func() error {
			fmt.Println("article log job")
			return nil
		},
	)

}

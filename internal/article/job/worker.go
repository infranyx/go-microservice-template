package articleJob

import (
	"context"

	"github.com/infranyx/go-grpc-template/pkg/wrapper"
)

func (j *job) logArticleWorker() wrapper.HandlerFunc {
	return func(ctx context.Context, args ...interface{}) (interface{}, error) {
		j.logger.Info("article log job")
		return nil, nil
	}
}

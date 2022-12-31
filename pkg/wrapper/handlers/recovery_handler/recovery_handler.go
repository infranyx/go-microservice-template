package wrapperRecoveryhandler

import (
	"context"

	"go.uber.org/zap"

	"github.com/infranyx/go-microservice-template/pkg/logger"
	"github.com/infranyx/go-microservice-template/pkg/wrapper"
)

var RecoveryHandler = func(f wrapper.HandlerFunc) wrapper.HandlerFunc {
	return func(ctx context.Context, args ...interface{}) (interface{}, error) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					logger.Zap.Sugar().Errorf("%v", r)
					return
				}
				logger.Zap.Error(err.Error(), zap.Error(err))
			}
		}()

		return f(ctx, args)
	}
}

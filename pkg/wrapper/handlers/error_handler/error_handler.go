package wrapperErrorhandler

import (
	"context"

	"github.com/getsentry/sentry-go"
	loggerConst "github.com/infranyx/go-grpc-template/pkg/constant/logger"
	customError "github.com/infranyx/go-grpc-template/pkg/error/custom_error"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/infranyx/go-grpc-template/pkg/wrapper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ErrorHandler = func(f wrapper.HandlerFunc) wrapper.HandlerFunc {
	return func(ctx context.Context, args interface{}) (interface{}, error) {
		res, err := f(ctx, args)
		if err != nil {
			hub := sentry.GetHubFromContext(ctx)
			logFields := []zapcore.Field{
				zap.String(loggerConst.TYPE, loggerConst.WORKER),
			}
			sentryContext := make(map[string]interface{})

			if ce := customError.AsCustomError(err); ce != nil {
				sentryContext[loggerConst.CODE] = ce.Code()
				sentryContext[loggerConst.DETAILS] = ce.Details()

				logFields = append(logFields,
					zap.Int(loggerConst.CODE, ce.Code()),
					zap.Any(loggerConst.DETAILS, ce.Details()),
				)
			}

			if hub != nil {
				sentryContext[loggerConst.TYPE] = loggerConst.WORKER
				hub.ConfigureScope(func(scope *sentry.Scope) {
					scope.SetLevel(sentry.LevelError)
					scope.SetContext("systemErr", sentryContext)
				})
				hub.CaptureException(err)
			}

			logger.Zap.Error(
				err.Error(),
				logFields...,
			)
		}

		return res, err
	}
}

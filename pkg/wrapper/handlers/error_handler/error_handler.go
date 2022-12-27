package wrapperErrorhandler

import (
	"context"

	"github.com/getsentry/sentry-go"
	loggerConstant "github.com/infranyx/go-grpc-template/pkg/constant/logger"
	customError "github.com/infranyx/go-grpc-template/pkg/error/custom_error"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/infranyx/go-grpc-template/pkg/wrapper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ErrorHandler = func(f wrapper.HandlerFunc) wrapper.HandlerFunc {
	return func(ctx context.Context, args ...interface{}) (interface{}, error) {
		res, err := f(ctx, args)
		if err != nil {
			hub := sentry.GetHubFromContext(ctx)
			logFields := []zapcore.Field{
				zap.String(loggerConstant.TYPE, loggerConstant.WORKER),
			}
			sentryContext := make(map[string]interface{})

			if ce := customError.AsCustomError(err); ce != nil {
				sentryContext[loggerConstant.CODE] = ce.Code()
				sentryContext[loggerConstant.DETAILS] = ce.Details()

				logFields = append(logFields,
					zap.Int(loggerConstant.CODE, ce.Code()),
					zap.Any(loggerConstant.DETAILS, ce.Details()),
				)
			}

			if hub != nil {
				sentryContext[loggerConstant.TYPE] = loggerConstant.WORKER
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

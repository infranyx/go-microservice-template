package wrapper

import (
	"github.com/getsentry/sentry-go"
	"github.com/infranyx/go-grpc-template/pkg/config"
	loggerConst "github.com/infranyx/go-grpc-template/pkg/constant/logger"
	customError "github.com/infranyx/go-grpc-template/pkg/error/custom_error"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	sentryUtils "github.com/infranyx/go-grpc-template/pkg/sentry/sentry_utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/context"
)

type middleware func(HandlerFunc) HandlerFunc
type HandlerFunc func(ctx context.Context, args interface{}) (interface{}, error)

func BuildChain(f HandlerFunc, m ...middleware) HandlerFunc {
	if len(m) == 0 {
		return f
	}
	return m[0](BuildChain(f, m[1:cap(m)]...))
}

func (hf HandlerFunc) ToCronJobFunc(ctx context.Context, args interface{}) func() {
	return func() {
		hf(ctx, args)
	}
}

var SentryMiddleware = func(f HandlerFunc) HandlerFunc {
	return func(ctx context.Context, args interface{}) (interface{}, error) {
		opts := &sentryUtils.Options{
			Repanic: true,
		}
		hub := sentry.GetHubFromContext(ctx)
		if hub == nil {
			hub = sentry.CurrentHub().Clone()
			ctx = sentry.SetHubOnContext(ctx, hub)
		}
		hub.Scope().SetExtra("args", args)
		hub.Scope().SetTag("application", config.Conf.App.AppName)
		hub.Scope().SetTag("AppEnv", config.Conf.App.AppEnv)
		defer sentryUtils.RecoverWithSentry(hub, ctx, opts)

		return f(ctx, args)
	}
}

var RecoveryMiddleware = func(f HandlerFunc) HandlerFunc {
	return func(ctx context.Context, args interface{}) (interface{}, error) {
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

var ErrorHandlerMiddleware = func(f HandlerFunc) HandlerFunc {
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

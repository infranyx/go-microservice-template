package wrapperSentryhandler

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/infranyx/go-grpc-template/pkg/config"
	sentryUtils "github.com/infranyx/go-grpc-template/pkg/sentry/sentry_utils"
	"github.com/infranyx/go-grpc-template/pkg/wrapper"
)

var SentryHandler = func(f wrapper.HandlerFunc) wrapper.HandlerFunc {
	return func(ctx context.Context, args ...interface{}) (interface{}, error) {
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

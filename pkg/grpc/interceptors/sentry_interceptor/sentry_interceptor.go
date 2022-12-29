package grpcSentryInterceptor

import (
	"context"

	"github.com/getsentry/sentry-go"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"

	"github.com/infranyx/go-grpc-template/pkg/config"
	sentryUtils "github.com/infranyx/go-grpc-template/pkg/sentry/sentry_utils"
)

func UnaryServerInterceptor(opts *sentryUtils.Options) grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		hub := sentry.GetHubFromContext(ctx)
		if hub == nil {
			hub = sentry.CurrentHub().Clone()
			ctx = sentry.SetHubOnContext(ctx, hub)
		}
		hub.Scope().SetExtra("request", req)
		hub.Scope().SetTransaction(info.FullMethod)
		hub.Scope().SetTag("application", config.BaseConfig.App.AppName)
		hub.Scope().SetTag("AppEnv", config.BaseConfig.App.AppEnv)

		defer sentryUtils.RecoverWithSentry(hub, ctx, opts)

		resp, err := handler(ctx, req)
		return resp, err
	}
}

func StreamServerInterceptor(opts *sentryUtils.Options) grpc.StreamServerInterceptor {
	return func(srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler) error {

		ctx := ss.Context()

		stream := grpcMiddleware.WrapServerStream(ss)
		stream.WrappedContext = ctx

		hub := sentry.GetHubFromContext(ctx)
		if hub == nil {
			hub = sentry.CurrentHub().Clone()
			ctx = sentry.SetHubOnContext(ctx, hub)
		}
		hub.Scope().SetTransaction(info.FullMethod)
		hub.Scope().SetTag("application", config.BaseConfig.App.AppName)
		hub.Scope().SetTag("AppEnv", config.BaseConfig.App.AppEnv)

		defer sentryUtils.RecoverWithSentry(hub, ctx, opts)

		err := handler(srv, stream)

		return err
	}
}

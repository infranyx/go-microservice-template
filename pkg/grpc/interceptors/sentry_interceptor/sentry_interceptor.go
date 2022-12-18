package grpcSentryInterceptor

import (
	"context"
	"github.com/infranyx/go-grpc-template/pkg/config"
	"time"

	"github.com/getsentry/sentry-go"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

type Options struct {
	Repanic         bool
	WaitForDelivery bool
	Timeout         time.Duration
}

func recoverWithSentry(hub *sentry.Hub, ctx context.Context, o *Options) {
	if err := recover(); err != nil {
		eventID := hub.RecoverWithContext(ctx, err)
		if eventID != nil && o.WaitForDelivery {
			hub.Flush(o.Timeout)
		}

		if o.Repanic {
			panic(err)
		}
	}
}

func UnaryServerInterceptor(opts *Options) grpc.UnaryServerInterceptor {
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
		// TODO : set application tag like echo
		hub.Scope().SetTag("application", "go-grpc-template")
		hub.Scope().SetTag("AppEnv", config.Conf.App.AppEnv)

		defer recoverWithSentry(hub, ctx, opts)

		resp, err := handler(ctx, req)
		return resp, err
	}
}

func StreamServerInterceptor(opts *Options) grpc.StreamServerInterceptor {
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
		// TODO : set application tag like echo
		hub.Scope().SetTag("application", "go-grpc-template")
		hub.Scope().SetTag("AppEnv", config.Conf.App.AppEnv)

		defer recoverWithSentry(hub, ctx, opts)

		err := handler(srv, stream)

		return err
	}
}

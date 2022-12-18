package grpcErrorInterceptor

import (
	"context"
	"github.com/getsentry/sentry-go"
	grpcTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	loggerConst "github.com/infranyx/go-grpc-template/pkg/constant/logger"
	errorUtils "github.com/infranyx/go-grpc-template/pkg/error/error_utils"
	grpcErrors "github.com/infranyx/go-grpc-template/pkg/error/grpc"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor returns a problem-detail error to client
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		resp, err := handler(ctx, req)

		if err != nil {
			grpcErr := grpcErrors.ParseError(err)

			hub := sentry.GetHubFromContext(ctx)
			if hub != nil {
				hub.ConfigureScope(func(scope *sentry.Scope) {
					scope.SetLevel(sentry.LevelError)
					scope.SetContext("grpcErr", map[string]interface{}{
						loggerConst.TYPE:        loggerConst.GRPC,
						loggerConst.TITILE:      grpcErr.GetTitle(),
						loggerConst.CODE:        grpcErr.GetCode(),
						loggerConst.STATUS:      grpcErr.GetStatus().String(),
						loggerConst.TIME:        grpcErr.GetTimestamp(),
						loggerConst.DETAILS:     grpcErr.GetDetails(),
						loggerConst.STACK_TRACE: errorUtils.RootStackTrace(err),
					})
					tags := grpcTags.Extract(ctx)
					for k, v := range tags.Values() {
						scope.SetTag(k, v.(string))
					}
				})
				hub.CaptureException(err)
			}

			logger.Zap.Error(
				err.Error(),
				zap.String(loggerConst.TYPE, loggerConst.GRPC),
				zap.String(loggerConst.TITILE, grpcErr.GetTitle()),
				zap.Int(loggerConst.CODE, grpcErr.GetCode()),
				zap.String(loggerConst.STATUS, grpcErr.GetStatus().String()),
				zap.Time(loggerConst.TIME, grpcErr.GetTimestamp()),
				zap.Any(loggerConst.DETAILS, grpcErr.GetDetails()),
				zap.String(loggerConst.STACK_TRACE, errorUtils.RootStackTrace(err)),
			)
			return nil, grpcErr.ToGrpcResponseErr()
		}

		return resp, err
	}
}

// StreamServerInterceptor returns a problem-detail error to client.
func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		err := handler(srv, ss)
		if err != nil {
			grpcErr := grpcErrors.ParseError(err)
			return grpcErr.ToGrpcResponseErr()
		}
		return err
	}
}

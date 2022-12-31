package grpcErrorInterceptor

import (
	"context"

	"github.com/getsentry/sentry-go"
	grpcTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	loggerConstant "github.com/infranyx/go-microservice-template/pkg/constant/logger"
	errorUtils "github.com/infranyx/go-microservice-template/pkg/error/error_utils"
	grpcErrors "github.com/infranyx/go-microservice-template/pkg/error/grpc"
	"github.com/infranyx/go-microservice-template/pkg/logger"
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
						loggerConstant.TYPE:        loggerConstant.GRPC,
						loggerConstant.TITILE:      grpcErr.GetTitle(),
						loggerConstant.CODE:        grpcErr.GetCode(),
						loggerConstant.STATUS:      grpcErr.GetStatus().String(),
						loggerConstant.TIME:        grpcErr.GetTimestamp(),
						loggerConstant.DETAILS:     grpcErr.GetDetails(),
						loggerConstant.STACK_TRACE: errorUtils.RootStackTrace(err),
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
				zap.String(loggerConstant.TYPE, loggerConstant.GRPC),
				zap.String(loggerConstant.TITILE, grpcErr.GetTitle()),
				zap.Int(loggerConstant.CODE, grpcErr.GetCode()),
				zap.String(loggerConstant.STATUS, grpcErr.GetStatus().String()),
				zap.Time(loggerConstant.TIME, grpcErr.GetTimestamp()),
				zap.Any(loggerConstant.DETAILS, grpcErr.GetDetails()),
				zap.String(loggerConstant.STACK_TRACE, errorUtils.RootStackTrace(err)),
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

package error_interceptors

import (
	"context"

	loggerConst "github.com/infranyx/go-grpc-template/constant/logger"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	grpcErrors "github.com/infranyx/go-grpc-template/shared/error/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// UnaryServerInterceptor returns a problem-detail error to client
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		l := logger.Zap
		resp, err := handler(ctx, req)
		if err != nil {
			grpcErr := grpcErrors.ParseError(err)
			l.Error(
				err.Error(),
				zap.String(loggerConst.TYPE, loggerConst.GRPC),
				zap.String(loggerConst.TITILE, grpcErr.GetTitle()),
				zap.Int(loggerConst.CODE, grpcErr.GetCode()),
				zap.String(loggerConst.STATUS, codes.Code(grpcErr.GetStatus()).String()),
				zap.Time(loggerConst.TIME, grpcErr.GetTimestamp()),
				zap.Any(loggerConst.DETAILS, grpcErr.GetDetails()),
				zap.String(loggerConst.STACK_TRACE, grpcErr.GetStackTrace()),
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
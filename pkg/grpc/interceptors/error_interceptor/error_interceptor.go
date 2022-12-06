package errorInterceptors

import (
	"context"
	errorUtils "github.com/infranyx/go-grpc-template/pkg/error/error_utils"

	grpcLoggerConst "github.com/infranyx/go-grpc-template/pkg/constant/grpc_logger"
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
			logger.Zap.Error(
				err.Error(),
				zap.String(grpcLoggerConst.TYPE, grpcLoggerConst.GRPC),
				zap.String(grpcLoggerConst.TITILE, grpcErr.GetTitle()),
				zap.Int(grpcLoggerConst.CODE, grpcErr.GetCode()),
				zap.String(grpcLoggerConst.STATUS, grpcErr.GetStatus().String()),
				zap.Time(grpcLoggerConst.TIME, grpcErr.GetTimestamp()),
				zap.Any(grpcLoggerConst.DETAILS, grpcErr.GetDetails()),
				zap.String(grpcLoggerConst.STACK_TRACE, errorUtils.RootStackTrace(err)),
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

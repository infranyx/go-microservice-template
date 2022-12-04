package logger_interceptors

import (
	"context"
	grpcLoggerConst "github.com/infranyx/go-grpc-template/pkg/constant/grpc_logger"
	"time"

	"github.com/infranyx/go-grpc-template/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor returns a problem-detail error to client
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()
		l := logger.Zap

		l.Sugar().Info(
			zap.String(grpcLoggerConst.TYPE, grpcLoggerConst.GRPC),
			zap.Any(grpcLoggerConst.REQUEST, req),
			zap.Time(grpcLoggerConst.TIME, startTime),
		)

		resp, err := handler(ctx, req)

		return resp, err
	}
}

// // StreamServerInterceptor returns a problem-detail error to client.
// func StreamLoggerInterceptor() grpc.StreamServerInterceptor {
// 	return func(
// 		srv interface{},
// 		ss grpc.ServerStream,
// 		info *grpc.StreamServerInfo,
// 		handler grpc.StreamHandler,
// 	) error {
// 		err := handler(srv, ss)
// 		if err != nil {
// 			grpcErr := grpcErrors.ParseError(err)
// 			return grpcErr.ToGrpcResponseErr()
// 		}
// 		return err
// 	}
// }

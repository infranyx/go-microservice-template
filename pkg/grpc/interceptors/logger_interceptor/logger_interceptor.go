package grpcLoggerInterceptor

import (
	"context"
	"time"

	loggerConst "github.com/infranyx/go-grpc-template/pkg/constant/logger"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor returns a problem-detail error to client
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()

		logger.Zap.Info(
			"Incoming Request",
			zap.String(loggerConst.TYPE, loggerConst.GRPC),
			zap.Any(loggerConst.REQUEST, req),
			zap.Time(loggerConst.TIME, startTime),
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

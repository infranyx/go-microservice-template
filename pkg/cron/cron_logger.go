package cronJob

import "github.com/infranyx/go-grpc-template/pkg/logger"

type cronLogger struct{}

func NewCronLogger() *cronLogger {
	return &cronLogger{}
}

func (cl *cronLogger) Info(msg string, keysAndValues ...interface{}) {
	logger.Zap.Sugar().Infow(msg, keysAndValues)
}

func (cl *cronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	logger.Zap.Sugar().Errorw(msg, keysAndValues)
}

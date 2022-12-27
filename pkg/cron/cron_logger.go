package cronJob

import "github.com/infranyx/go-grpc-template/pkg/logger"

type cronLogger struct{}

func NewLogger() *cronLogger {
	return &cronLogger{}
}

func (l *cronLogger) Info(msg string, keysAndValues ...interface{}) {
	logger.Zap.Sugar().Infow(msg, keysAndValues)
}

func (l *cronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	logger.Zap.Sugar().Errorw(msg, keysAndValues)
}

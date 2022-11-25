package logger

import (
	"github.com/infranyx/go-grpc-template/pkg/logger/contracts"
	"github.com/infranyx/go-grpc-template/pkg/logger/zap"
)

var (
	Defaultlogger *Logger
)

// interface wrapper
type Logger struct {
	contracts.Logger
}

func NewLogger() *Logger {
	Defaultlogger = &Logger{
		zap.NewZapLogger(&contracts.LogConfig{
			LogLevel: "debug",
			LogType:  contracts.Zap,
		}),
	}

	return Defaultlogger
}

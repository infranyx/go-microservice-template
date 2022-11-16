package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger() *Logger {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Encoding:    "json",
		OutputPaths: []string{"stdout", "tmp/logs/main.log"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.LowercaseLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	logger := zap.Must(cfg.Build())
	defer logger.Sync()
	sugar := logger.Sugar()
	return &Logger{
		sugar,
	}
}

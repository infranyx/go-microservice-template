package logger

import (
	"os"

	"github.com/infranyx/go-grpc-template/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Zap *zap.Logger
)

func init() {
	Zap = initLogger()
}

// Init logger
func initLogger() *zap.Logger {
	logLevel := zapcore.DebugLevel

	var logWriter zapcore.WriteSyncer

	var encoderCfg zapcore.EncoderConfig
	var encoder zapcore.Encoder

	if config.IsProduction() {
		encoderCfg = zap.NewProductionEncoderConfig()
		encoderCfg.NameKey = "[SERVICE]"
		encoderCfg.TimeKey = "[TIME]"
		encoderCfg.LevelKey = "[LEVEL]"
		encoderCfg.FunctionKey = "[CALLER]"
		encoderCfg.CallerKey = "[LINE]"
		encoderCfg.MessageKey = "[MESSAGE]"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
		encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
		encoderCfg.EncodeName = zapcore.FullNameEncoder
		encoderCfg.EncodeDuration = zapcore.StringDurationEncoder
		encoder = zapcore.NewJSONEncoder(encoderCfg)
		logFile, _ := os.OpenFile("tmp/logs/main.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		logWriter = zapcore.AddSync(logFile)
	} else {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
		encoderCfg.NameKey = "[SERVICE]"
		encoderCfg.TimeKey = "[TIME]"
		encoderCfg.LevelKey = "[LEVEL]"
		encoderCfg.FunctionKey = "[CALLER]"
		encoderCfg.CallerKey = "[LINE]"
		encoderCfg.MessageKey = "[MESSAGE]"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderCfg.EncodeName = zapcore.FullNameEncoder
		encoderCfg.EncodeDuration = zapcore.StringDurationEncoder
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderCfg.EncodeCaller = zapcore.FullCallerEncoder
		encoderCfg.ConsoleSeparator = " | "
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
		logWriter = zapcore.AddSync(os.Stdout)
	}

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller())

	return logger
}

// func (l *zapLogger) GrpcClientInterceptorLogger(method string, req, reply interface{}, time time.Duration, metaData map[string][]string, err error) {
// 	l.logger.Info(
// 		constants.GRPC,
// 		zap.String(constants.METHOD, method),
// 		zap.Any(constants.REQUEST, req),
// 		zap.Any(constants.REPLY, reply),
// 		zap.Duration(constants.TIME, time),
// 		zap.Any(constants.METADATA, metaData),
// 		zap.Error(err),
// 	)
// }

// func mapToFields(fields map[string]interface{}) []zap.Field {
// 	var zapFields []zap.Field
// 	for k, v := range fields {
// 		zapFields = append(zapFields, zap.Any(k, v))
// 	}

// 	return zapFields
// }

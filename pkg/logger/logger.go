package logger

import (
	"github.com/infranyx/go-grpc-template/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var Zap *zap.Logger

func init() {
	Zap = newLogger()
}

func newLogger() *zap.Logger {
	var logWriter zapcore.WriteSyncer
	var encoderCfg zapcore.EncoderConfig
	var encoder zapcore.Encoder

	if config.IsProdEnv() {
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

		_, callerDir, _, ok := runtime.Caller(0)
		if !ok {
			log.Fatal("Error generating env dir")
		}
		tmpLogDir := filepath.Join(filepath.Dir(callerDir), "../..", "tmp/logs")

		if _, err := os.Stat(tmpLogDir); os.IsNotExist(err) {
			_ = os.MkdirAll(tmpLogDir, os.ModePerm)
		}

		logFile, _ := os.OpenFile(filepath.Join(tmpLogDir, "main.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
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

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(zapcore.DebugLevel))
	return zap.New(core, zap.AddCaller())
}

package common

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func MustInitLogger() *zap.SugaredLogger {
	switch env := os.Getenv("APP_ENV"); env {
	case "DEV":
		return zap.Must(zap.NewDevelopment()).Sugar()
	case "PROD":
		return zap.Must(zap.NewProduction()).Sugar()
	default:
		return NewDefaultLogger().Sugar()
	}
}

func NewDefaultLogger() *zap.Logger {
	return zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "@timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
			EncodeDuration: zapcore.NanosDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}),
		zapcore.AddSync(os.Stdout),
		zap.NewAtomicLevelAt(zapcore.DebugLevel),
	), zap.AddCaller(), zap.AddCallerSkip(1))
}

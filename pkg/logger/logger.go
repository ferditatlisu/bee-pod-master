package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const timeFormat string = "2006-01-02T15:04:05.999Z"

var zapLogger, _ = NewLogger()

type LoggerFactory struct {
	Logger *zap.Logger
}

func Logger() *zap.Logger {
	return zapLogger
}

func GetLogger(ctx *gin.Context) *zap.Logger {
	loggerInterface, _ := ctx.Get("logger")
	loggerFactory := LoggerFactory{Logger: loggerInterface.(*zap.Logger)}
	return loggerFactory.Logger
}

func NewLogger() (*zap.Logger, error) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "sourceLocation",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout(timeFormat),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	return config.Build(zap.AddStacktrace(zap.FatalLevel))
}

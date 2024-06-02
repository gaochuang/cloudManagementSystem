package log

import (
	"context"
	"go.uber.org/zap"
	"strings"
)

var (
	Logger LoggerInterface
)

type Configuration struct {
	LogType    string
	LogFileDir string
	LogLevel   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}

type LoggerInterface interface {
	LogDebug(msg string, fields ...zap.Field)
	LogDebugWithCtx(ctx context.Context, msg string, fields ...zap.Field)
	LogInfo(msg string, fields ...zap.Field)
	LogInfoWithCtx(ctx context.Context, msg string, fields ...zap.Field)
	LogError(msg string, fields ...zap.Field)
	LogErrorWithCtx(ctx context.Context, msg string, fields ...zap.Field)
	LogWarn(msg string, fields ...zap.Field)
	LogWarnWithCtx(ctx context.Context, msg string, fields ...zap.Field)
	AddLoggerCtx(ctx context.Context, key, value interface{}) context.Context
	GetLoggerFromCtx(ctx context.Context) *zap.Logger
}

func SetupLoggers(logType, logDir, logLevel string) (err error) {
	var level string
	switch strings.ToLower(logLevel) {
	case "error":
		level = "error"
	case "warn":
		level = "warn"
	default:
		level = "info"
	}
	c := &Configuration{
		LogType:    logType,
		LogFileDir: logDir,
		LogLevel:   level,
		MaxSize:    200,
		MaxAge:     10,
		MaxBackups: 2,
		Compress:   true,
	}
	Logger, err = NewZapLogger(c)
	return
}

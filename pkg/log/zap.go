package log

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"strings"
	"time"
)

const (
	STDERR = "stderr"
	FILE   = "file"
)

type zapLogger struct {
	logger    *zap.Logger
	writer    io.Writer
	verbosity int
}

func NewZapLogger(c *Configuration) (LoggerInterface, error) {
	if c == nil {
		return nil, errors.New("input c is null")
	}
	var write io.Writer
	switch strings.ToLower(c.LogType) {
	//日志将输出到标准错误输出
	case STDERR:
		write = os.Stderr
	case FILE:
		write = &lumberjack.Logger{
			Filename:   c.LogFileDir, //日志文件路径
			MaxSize:    c.MaxSize,    //每个日志文件的最大尺寸(MB)
			MaxAge:     c.MaxAge,     //保留旧文件最大天数
			MaxBackups: c.MaxBackups, //保留旧文件最大数量
			Compress:   c.Compress,   //是否压缩旧文件
		}
	default:
		write = os.Stdout //如果LogType是其他值或未指定，日志将输出到标准输出。
	}

	config := zapcore.EncoderConfig{
		TimeKey:        "time",
		MessageKey:     "message",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "trace",
		EncodeTime:     formatEncodeTime,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	var logLevel zapcore.Level

	if err := logLevel.UnmarshalText([]byte(c.LogLevel)); err != nil {
		return nil, err
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(config),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(write)),
		logLevel)

	var cores []zapcore.Core

	cores = append(cores, core)
	logger := zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.AddCallerSkip(2))

	return &zapLogger{
		logger:    logger,
		writer:    write,
		verbosity: 0,
	}, nil
}

func (zl *zapLogger) LogDebug(msg string, fields ...zap.Field) {
	zl.logger.Debug(msg, fields...)
}

func (zl *zapLogger) LogDebugWithCtx(ctx context.Context, msg string, fields ...zap.Field) {
	if val := ctx.Value("X-Trace-Id"); val != nil {
		if trace, ok := val.(string); ok {
			zl.LogDebug(msg, append(fields, zap.String("traceId", trace))...)
		}
	} else {
		zl.LogDebug(msg, fields...)
	}
}

func (zl *zapLogger) LogInfo(msg string, fields ...zap.Field) {
	zl.logger.Info(msg, fields...)
}

func (zl *zapLogger) LogInfoWithCtx(ctx context.Context, msg string, fields ...zap.Field) {
	if traceID, ok := ctx.Value("X-Trace-Id").(string); ok {
		zl.LogInfo(msg, append(fields, zap.String("traceId", traceID))...)
	} else {
		zl.LogInfo(msg, fields...)
	}
}

func (zl *zapLogger) LogError(msg string, fields ...zap.Field) {
	zl.logger.Error(msg, fields...)
}

func (zl *zapLogger) LogErrorWithCtx(ctx context.Context, msg string, fields ...zap.Field) {
	if val := ctx.Value("X-Trace-Id"); val != nil {
		if trace, ok := val.(string); ok {
			zl.LogError(msg, append(fields, zap.String("traceId", trace))...)
		}
	} else {
		zl.LogError(msg, fields...)
	}
}

func (zl *zapLogger) LogWarn(msg string, fields ...zap.Field) {
	zl.logger.Warn(msg, fields...)
}

func (zl *zapLogger) LogWarnWithCtx(ctx context.Context, msg string, fields ...zap.Field) {
	if val := ctx.Value("X-Trace-Id"); val != nil {
		if trace, ok := val.(string); ok {
			zl.LogWarn(msg, append(fields, zap.String("traceId", trace))...)
		}
	} else {
		zl.LogWarn(msg, fields...)
	}

}

func (zl *zapLogger) GetLoggerFromCtx(ctx context.Context) *zap.Logger {
	log, ok := ctx.Value("X-Trace-Id").(*zap.Logger)
	if ok {
		return log.WithOptions(zap.AddCallerSkip(-1))
	}
	return zl.logger
}

func (zl *zapLogger) AddLoggerCtx(ctx context.Context, key, value interface{}) context.Context {
	ctx = context.WithValue(ctx, key, value)
	return ctx
}

func formatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%d/%02d/%02d-%02d:%02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.UnixMilli()))
}

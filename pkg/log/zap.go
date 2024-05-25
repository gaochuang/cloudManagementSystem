package log

import (
	"fmt"
	"github.com/gaochuang/cloudManagementSystem/common"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

var (
	level zapcore.Level
)

func Zap() (logger *zap.Logger) {
	ok, _ := DirExit(common.CONFIG.Zap.Director)
	if !ok {
		fmt.Printf("create %v directory\n", common.CONFIG.Zap.Director)
		err := os.Mkdir(common.CONFIG.Zap.Director, os.ModePerm)
		if err != nil {
			panic(fmt.Errorf("create %s failed, reason: +%v", common.CONFIG.Zap.Director, err))
		}
	}

	switch common.CONFIG.Zap.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "panic":
		level = zap.PanicLevel
	case "dpanic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	if zap.DebugLevel == level || zap.ErrorLevel == level {
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}

	return logger
}

func getEncoderCore() (core zapcore.Core) {
	writer, err := getWriteSyncer()
	if err != nil {
		fmt.Printf("get write syncer failed: %v \n", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

func getWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(path.Join(common.CONFIG.Zap.Director, "%Y-%m-%d.log"),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithLinkName(common.CONFIG.Zap.LinkName),
	)

	if common.CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

func getEncoder() zapcore.Encoder {
	if common.CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
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
	switch {
	case common.CONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder":
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case common.CONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder":
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case common.CONFIG.Zap.EncodeLevel == "CapitalLevelEncoder":
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case common.CONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder":
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

func formatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%d/%02d/%02d-%02d:%02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.UnixMilli()))
}

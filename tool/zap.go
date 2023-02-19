package tool

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"kubernetes_management_system/common"
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
		rotatelogs.WithMaxAge(7*24*time.Hour),               // 文件最大保存时间
		rotatelogs.WithRotationTime(8*time.Hour),            //日志切割时间间隔
		rotatelogs.WithLinkName(common.CONFIG.Zap.LinkName), //// 生成软链，指向最新日志文件
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
	switch {
	case common.CONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case common.CONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case common.CONFIG.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case common.CONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

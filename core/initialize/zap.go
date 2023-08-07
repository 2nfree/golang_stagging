package initialize

import (
	"go.uber.org/zap/zapcore"
	"golang-stagging/core"
	"os"
)

type ZapCoreType int

const (
	ZAP ZapCoreType = iota + 1
	GORM
)

// GetZapCores 获取日志配置
func GetZapCores(types ZapCoreType) zapcore.Core {
	switch types {
	case ZAP:
		return createCore(core.Config.Log.Path, true, core.Config.Log.Level)
	case GORM:
		return createCore(core.Config.Database.Path, core.Config.Database.ConsoleLog, core.Config.Database.LogMode)
	default:
		return zapcore.NewCore(EncoderConfig(true), zapcore.AddSync(os.Stdout), GetZapLogLevel(core.Config.Log.Level))
	}
}

func createCore(path string, isConsole bool, logMode string) zapcore.Core {
	if path == "" {
		return zapcore.NewCore(EncoderConfig(true), zapcore.AddSync(os.Stdout), GetZapLogLevel(logMode))
	} else if path != "" && !isConsole {
		return zapcore.NewCore(EncoderConfig(false), Writer(path), GetZapLogLevel(logMode))
	} else {
		return zapcore.NewTee(
			zapcore.NewCore(EncoderConfig(false), Writer(path), GetZapLogLevel(logMode)),
			zapcore.NewCore(EncoderConfig(true), zapcore.AddSync(os.Stdout), GetZapLogLevel(logMode)),
		)
	}
}

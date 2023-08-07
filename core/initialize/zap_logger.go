package initialize

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang-stagging/core"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

// InitZapLogger 初始化应用日志
func InitZapLogger() {
	core.Logger = zap.New(GetZapCores(ZAP), zap.Development(), zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	core.SugaredLogger = core.Logger.Sugar()
}

// EncoderConfig 日志格式
func EncoderConfig(isConsole bool) zapcore.Encoder {
	var encoderLevel zapcore.LevelEncoder
	if core.Config.Log.JSONFormat && !isConsole {
		encoderLevel = zapcore.CapitalLevelEncoder
	} else {
		encoderLevel = zapcore.CapitalColorLevelEncoder
	}
	zapEncode := zapcore.EncoderConfig{
		MessageKey:          "Message",
		LevelKey:            "Level",
		TimeKey:             "Timestamp",
		NameKey:             "Name",
		CallerKey:           "Caller",
		FunctionKey:         "Function",
		StacktraceKey:       "Stacktrace",
		SkipLineEnding:      false,
		LineEnding:          zapcore.DefaultLineEnding,
		EncodeLevel:         encoderLevel,
		EncodeTime:          encodeTime,
		EncodeDuration:      zapcore.SecondsDurationEncoder,
		EncodeCaller:        zapcore.ShortCallerEncoder,
		EncodeName:          zapcore.FullNameEncoder,
		NewReflectedEncoder: nil,
		ConsoleSeparator:    " ",
	}
	if core.Config.Log.JSONFormat && !isConsole {
		return zapcore.NewJSONEncoder(zapEncode)
	} else {
		return zapcore.NewConsoleEncoder(zapEncode)
	}
}

// GetZapLogLevel 日志等级
func GetZapLogLevel(logLevel string) zapcore.LevelEnabler {
	switch logLevel {
	case "DEBUG":
		return zapcore.DebugLevel
	case "INFO":
		return zapcore.InfoLevel
	case "WARN":
		return zapcore.WarnLevel
	case "ERROR":
		return zapcore.ErrorLevel
	default:
		return zapcore.DebugLevel
	}
}

// Writer 日志分片写入
func Writer(filename string) zapcore.WriteSyncer {
	var lumberJackLogger *lumberjack.Logger
	lumberJackLogger = &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    core.Config.Log.MaxSize,
		MaxBackups: core.Config.Log.MaxBackup,
		MaxAge:     core.Config.Log.MaxAge,
		Compress:   core.Config.Log.Compress,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 时间格式
func encodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}

package initialize

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang-stagging/core"
	"gorm.io/gorm/logger"
	"time"
)

// GormLogger gorm日志打印，实现接口
type GormLogger struct {
	*zap.Logger
	logger.Config
}

// InitGormLogger 初始化gorm日志
func InitGormLogger() *GormLogger {
	return &GormLogger{
		Logger: zap.New(GetZapCores(GORM), zap.Development(), zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)),
		Config: logger.Config{
			SlowThreshold:             20 * time.Second,  // 慢 SQL 阈值
			LogLevel:                  getGormLogLevel(), // Log level
			IgnoreRecordNotFoundError: true,              // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,             // 彩色打印
			ParameterizedQueries:      true,              // 打印sql参数
		},
	}
}

// 获取日志等级
func getGormLogLevel() logger.LogLevel {
	switch core.Config.Database.LogMode {
	case "INFO":
		return logger.Info
	case "WARN":
		return logger.Warn
	case "ERROR":
		return logger.Error
	case "NONE":
		return logger.Silent
	default:
		return logger.Info
	}
}

func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

func (l *GormLogger) Info(c context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.Logger.Info(msg, zap.Any("traceID", c.Value("traceID")), zap.Any("data", data))
	}
}

func (l *GormLogger) Warn(c context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.Logger.Warn(msg, zap.Any("traceID", c.Value("traceID")), zap.Any("data", data))
	}
}

func (l *GormLogger) Error(c context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.Logger.Error(msg, zap.Any("traceID", c.Value("traceID")), zap.Any("data", data))
	}
}

func (l *GormLogger) Trace(c context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.Logger.Error("",
				zap.Any("traceID", c.Value("traceID")),
				zap.Float64("runTime", float64(elapsed.Nanoseconds())/1e6),
				zap.String("row", "-"),
				zap.String("sql", sql),
				zap.Any("error", err),
			)
		} else {
			l.Logger.Error("",
				zap.Any("traceID", c.Value("traceID")),
				zap.Float64("runTime", float64(elapsed.Nanoseconds())/1e6),
				zap.Int64("row", rows),
				zap.String("sql", sql),
				zap.Any("error", err),
			)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.Logger.Warn("",
				zap.Any("traceID", c.Value("traceID")),
				zap.Float64("runTime", float64(elapsed.Nanoseconds())/1e6),
				zap.String("row", "-"),
				zap.String("sql", sql),
				zap.Any("slowLog", slowLog),
			)
		} else {
			l.Logger.Warn("",
				zap.Any("traceID", c.Value("traceID")),
				zap.Float64("runTime", float64(elapsed.Nanoseconds())/1e6),
				zap.Int64("row", rows),
				zap.String("sql", sql),
				zap.Any("slowLog", slowLog),
			)
		}
	case l.LogLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			l.Logger.Info("",
				zap.Any("traceID", c.Value("traceID")),
				zap.Float64("runTime", float64(elapsed.Nanoseconds())/1e6),
				zap.String("row", "-"),
				zap.String("sql", sql),
			)
		} else {
			l.Logger.Info("",
				zap.Any("traceID", c.Value("traceID")),
				zap.Float64("runTime", float64(elapsed.Nanoseconds())/1e6),
				zap.Int64("row", rows),
				zap.String("sql", sql),
			)
		}
	}
}

func (l *GormLogger) ParamsFilter(c context.Context, sql string, params ...interface{}) (string, []interface{}) {
	if l.Config.ParameterizedQueries {
		return sql, nil
	}
	return sql, params
}

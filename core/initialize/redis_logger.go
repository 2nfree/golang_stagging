package initialize

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

// RDBLogger redis client使用的logger
type RDBLogger struct {
	*zap.Logger
}

func (l *RDBLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	l.Logger.Warn(fmt.Sprintf(format, v), zap.Any("traceID", ctx.Value("traceID")))
}

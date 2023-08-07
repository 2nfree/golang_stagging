package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang-stagging/core"
	"time"
)

// LogMiddle 日志中间件，用于记录api请求
func LogMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		//traceID := c.GetHeader("traceID")
		//if traceID == "" {
		//	traceID = uuid.New().String()
		//}
		start := time.Now()
		c.Next()
		cost := time.Now().Sub(start)
		core.Logger.Info(
			"",
			zap.String("requestInfo", fmt.Sprintf("%v %v %v %v %v",
				c.Request.Method,
				c.Request.URL.Path,
				c.Request.Proto,
				c.Writer.Status(),
				cost,
			)),
			//zap.String("traceID", traceID),
			zap.String("clientIP", c.ClientIP()),
		)
	}
}

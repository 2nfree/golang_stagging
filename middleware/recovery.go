package middleware

import (
	"github.com/gin-gonic/gin"
	"golang-stagging/core"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

// RecoveryMiddle panic recovery记录中间件
func RecoveryMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					core.SugaredLogger.Errorf("| %v\n%v%v", c.Request.URL.Path, err, string(httpRequest))
					c.Error(err.(error))
					c.Abort()
					return
				}
				core.SugaredLogger.Errorf("[Recovery from panic] panic error: %v\n%v", err, string(httpRequest))
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

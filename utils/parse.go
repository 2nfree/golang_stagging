package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

// ParseDuration 将string转化为时间期间, 1h=一小时, 1d=一天等
func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")
		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}
	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}

// IsHttps 判断是否为https
func IsHttps(c *gin.Context) bool {
	if c.GetHeader("X-Forwarded-Proto") == "https" || c.Request.TLS != nil {
		return true
	}
	return false
}

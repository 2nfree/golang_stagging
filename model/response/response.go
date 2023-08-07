package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(code, Response{
		code,
		data,
		msg,
	})
}

func Forbidden(msg string, c *gin.Context) {
	c.JSON(http.StatusForbidden, Response{
		http.StatusForbidden,
		gin.H{"reload": true},
		msg,
	})
	c.Abort()
}

package router

import (
	"github.com/gin-gonic/gin"
	"golang-stagging/api"
)

type SystemRouter struct{}

func (router *SystemRouter) InitSystemRouter(routerGroup *gin.RouterGroup) {
	group := routerGroup.Group("system")
	group.GET("info", api.Groups.GetSystemInfo)
}

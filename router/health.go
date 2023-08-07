package router

import (
	"github.com/gin-gonic/gin"
	"golang-stagging/api"
)

type HealthCheckRouter struct{}

func (router *HealthCheckRouter) InitHealthCheckRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("health", api.Groups.HealthCheck)
}

package router

import (
	"github.com/gin-gonic/gin"
	"golang-stagging/middleware"
)

// Group 路由组
type Group struct {
	HealthCheckRouter
	SystemRouter
}

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.CorsMiddle(), middleware.LogMiddle(), middleware.RecoveryMiddle())
	routerGroup := new(Group)
	group := r.Group("/api/v1")
	routerGroup.InitHealthCheckRouter(group)
	routerGroup.InitSystemRouter(group)
	return r
}

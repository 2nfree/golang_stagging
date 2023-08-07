package api

import (
	"github.com/gin-gonic/gin"
	"golang-stagging/model/response"
)

type HealthCheckApi struct{}

// HealthCheck
// @Tags      System
// @Summary   健康检查
// @Security
// @Produce   application/json
// @Success   200   nil  ok
// @Router    /health [get]
func (api *HealthCheckApi) HealthCheck(c *gin.Context) {
	response.Result(200, nil, "ok", c)
}

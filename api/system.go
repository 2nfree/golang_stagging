package api

import (
	"github.com/gin-gonic/gin"
	"golang-stagging/core"
	"golang-stagging/model/response"
	"net/http"
)

type SystemApi struct{}

// GetSystemInfo
// @Tags      System
// @Summary   获取当前所在操作系统的资源消耗
// @Security
// @Produce   application/json
// @Success   200   {object}  response.Response{data=response.SystemResourceInfo,msg=string}  "当前所在操作系统的资源消耗"
// @Router    /system/info [get]
func (api *SystemApi) GetSystemInfo(c *gin.Context) {
	resourceInfo, osType, err := systemService.GetResourceInfo()
	systemResourceInfo := response.SystemResourceInfo{
		OSType:       osType,
		ResourceInfo: resourceInfo,
	}
	if err != nil {
		core.SugaredLogger.Errorf("get system resources info failed with error: %v", err)
		response.Result(http.StatusBadRequest, systemResourceInfo, "获取失败", c)
		return
	}
	response.Result(http.StatusOK, systemResourceInfo, "获取成功", c)
}

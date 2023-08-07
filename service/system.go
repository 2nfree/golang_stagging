package service

import "golang-stagging/utils"

type SystemService struct{}

// GetResourceInfo 获取当前所在操作系统的资源消耗
func (service *SystemService) GetResourceInfo() (info utils.SystemInfo, osType string, err error) {
	if utils.IsContainer() {
		info, err = utils.GetContainerInfo()
		osType = "container"
		if err != nil {
			return
		}
		return
	} else {
		info, err = utils.GetHardwareInfo()
		osType = "hardware"
		if err != nil {
			return
		}
		return
	}
}

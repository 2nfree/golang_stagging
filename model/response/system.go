package response

import "golang-stagging/utils"

type SystemResourceInfo struct {
	OSType       string           `json:"os_type"`
	ResourceInfo utils.SystemInfo `json:"system_info"`
}

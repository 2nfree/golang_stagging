package api

import "golang-stagging/service"

// Group Api组
type Group struct {
	HealthCheckApi
	SystemApi
}

var (
	Groups        = new(Group)
	systemService = service.Groups.SystemService
)

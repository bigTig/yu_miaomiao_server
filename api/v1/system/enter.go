package system

import "yuyu/service"

type ApiGroup struct {
	BaseApi
	JwtApi
}

var (
	userService   = service.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService    = service.ServiceGroupApp.SystemServiceGroup.JwtService
	AdvertService = service.ServiceGroupApp.SystemServiceGroup.AdvertService
)

package system

import "yuyu/service"

type ApiGroup struct {
	BaseApi
	JwtApi
	HealthNewsApi
	FastCateApi
}

var (
	userService       = service.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService        = service.ServiceGroupApp.SystemServiceGroup.JwtService
	advertService     = service.ServiceGroupApp.SystemServiceGroup.AdvertService
	categoryService   = service.ServiceGroupApp.SystemServiceGroup.CategoryService
	brandService      = service.ServiceGroupApp.SystemServiceGroup.BrandService
	healthNewsService = service.ServiceGroupApp.SystemServiceGroup.HealthNewsService
	fastCateService   = service.ServiceGroupApp.SystemServiceGroup.FastCateService
)

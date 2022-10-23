package system

import (
	"github.com/gin-gonic/gin"
	v1 "yuyu/api/v1"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi

	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("wxLogin", baseApi.WxLogin)
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.GET("advertList", baseApi.AdvertList)
		baseRouter.POST("insertAdvert", baseApi.InsertAdvert)
		baseRouter.PUT("updateAdvert", baseApi.UpdateAdvert)
		baseRouter.DELETE("deleteAdvert/:id", baseApi.DeleteAdvert)
	}

	return baseRouter
}

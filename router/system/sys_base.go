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
		baseRouter.POST("advertList", baseApi.AdvertList)
		baseRouter.POST("categoryList", baseApi.CategoryList)
		baseRouter.POST("brandList", baseApi.BrandList)
	}

	return baseRouter
}

// InitBaseAuthRouter
// 初始化需要权限的基础接口
func (s *BaseRouter) InitBaseAuthRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	baseAuthRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi

	{
		baseAuthRouter.POST("insertAdvert", baseApi.InsertAdvert)
		baseAuthRouter.PUT("updateAdvert", baseApi.UpdateAdvert)
		baseAuthRouter.DELETE("deleteAdvert/:id", baseApi.DeleteAdvert)
		baseAuthRouter.POST("insertCategory", baseApi.InsertCategory)
		baseAuthRouter.PUT("updateCategory", baseApi.UpdateCategory)
		baseAuthRouter.DELETE("deleteCategory/:id", baseApi.DeleteCategory)
		baseAuthRouter.POST("insertBrand", baseApi.InsertBrand)
		baseAuthRouter.PUT("updateBrand", baseApi.UpdateBrand)
		baseAuthRouter.DELETE("deleteBrand/:id", baseApi.DeleteBrand)
	}

	return baseAuthRouter
}

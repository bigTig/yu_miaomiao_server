package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "yuyu/docs"
	"yuyu/global"
	"yuyu/middleware"
	"yuyu/router"
)

// Routers 初始化总路由
func Routers() *gin.Engine {
	Router := gin.Default()

	systemRouter := router.RouterGroupApp.System

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GvaLog.Info("register swagger handler")

	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")

	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}

	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		systemRouter.InitJwtRouter(PrivateGroup)      // jwt 相关路由
		systemRouter.InitUserRouter(PrivateGroup)     // 注册用户路由
		systemRouter.InitBaseAuthRouter(PrivateGroup) // 需要权限的基础接口
	}

	global.GvaLog.Info("router register success")

	return Router
}

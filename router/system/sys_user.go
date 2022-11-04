package system

import (
	"github.com/gin-gonic/gin"
	v1 "yuyu/api/v1"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi

	{
		userRouter.PUT("setUserInfo", baseApi.SetUserInfo)        // 设置用户信息
		userRouter.GET("getUserInfo", baseApi.GetUserInfo)        // 获取用户信息
		userRouter.POST("logout", baseApi.Logout)                 // 退出登录
		userRouter.POST("changePassword", baseApi.ChangePassword) // 修改密码
	}
}

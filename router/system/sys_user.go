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
		userRouter.PUT("setUserInfo", baseApi.SetUserInfo) // 设置用户信息
	}
}

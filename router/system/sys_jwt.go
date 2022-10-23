package system

import (
	"github.com/gin-gonic/gin"
	v1 "yuyu/api/v1"
)

type JwtRouter struct{}

func (j *JwtRouter) InitJwtRouter(Router *gin.RouterGroup) {
	jwtRouter := Router.Group("jwt")
	jwtApi := v1.ApiGroupApp.SystemApiGroup.JwtApi
	{
		jwtRouter.POST("jsonInBlacklist", jwtApi.JsonInBlacklist) // jwt 加入黑名单
	}
}

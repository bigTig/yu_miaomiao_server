package system

import (
	"github.com/gin-gonic/gin"
	v1 "yuyu/api/v1"
)

type HealthNewsRouter struct{}

func (h *HealthNewsRouter) InitHealthNewsRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	healthNewsRouter := Router.Group("health")
	healthNewsApi := v1.ApiGroupApp.SystemApiGroup.HealthNewsApi
	{
		healthNewsRouter.POST("healthNewsList", healthNewsApi.HealthNewsList)     // 获取新闻资讯列表
		healthNewsRouter.POST("insertHealthNew", healthNewsApi.InsertHealthNew)   // 新增新闻资讯
		healthNewsRouter.PUT("updateHealthNew", healthNewsApi.UpdateHealthNew)    // 更新新闻资讯
		healthNewsRouter.DELETE("deleteHealthNew", healthNewsApi.DeleteHealthNew) // 删除新闻资讯
		healthNewsRouter.GET("healthNewDetail", healthNewsApi.HealthNewDetail)    // 获取新闻资讯详情
	}

	return healthNewsRouter
}

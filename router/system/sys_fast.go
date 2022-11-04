package system

import (
	"github.com/gin-gonic/gin"
	v1 "yuyu/api/v1"
)

type FastRouter struct{}

func (f *FastRouter) InitFastRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	fastRouter := Router.Group("fast")
	fastApi := v1.ApiGroupApp.SystemApiGroup.FastApi
	{
		fastRouter.GET("fastCateList", fastApi.FastCateList)            // 获取禁食分类列表
		fastRouter.POST("insertFastCate", fastApi.InsertFastCate)       // 添加禁食分类
		fastRouter.PUT("updateFastCate", fastApi.UpdateFastCate)        // 更新禁食分类
		fastRouter.DELETE("deleteFastCate/:id", fastApi.DeleteFastCate) // 删除禁食分类
		fastRouter.POST("fastList", fastApi.FastList)                   // 禁食列表
		fastRouter.POST("insertFast", fastApi.InsertFast)               // 添加禁食
		fastRouter.PUT("updateFast", fastApi.UpdateFast)                // 更新禁食
		fastRouter.DELETE("deleteFast/:id", fastApi.DeleteFast)         // 删除禁食
		fastRouter.GET("fastDetailById", fastApi.FastDetailById)        // 根据id获取禁食详情
	}

	return fastRouter
}

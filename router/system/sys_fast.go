package system

import (
	"github.com/gin-gonic/gin"
	v1 "yuyu/api/v1"
)

type FastRouter struct{}

func (f *FastRouter) InitFastRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	fastRouter := Router.Group("fast")
	fastApi := v1.ApiGroupApp.SystemApiGroup.FastCateApi
	{
		fastRouter.GET("fastCateList", fastApi.FastCateList)        // 获取禁食分类列表
		fastRouter.POST("insertFastCate", fastApi.InsertFastCate)   // 添加禁食分类
		fastRouter.PUT("updateFastCate", fastApi.UpdateFastCate)    // 更新禁食分类
		fastRouter.DELETE("deleteFastCate", fastApi.DeleteFastCate) // 删除禁食分类
	}

	return fastRouter
}

package system

import (
	"github.com/gin-gonic/gin"
	v1 "yuyu/api/v1"
)

type AddressRouter struct{}

func (a *AddressRouter) InitAddressRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	addrRouter := Router.Group("address")
	addrApi := v1.ApiGroupApp.SystemApiGroup.AddressApi
	{
		addrRouter.GET("addressList", addrApi.AddressList)             // 收货地址列表
		addrRouter.POST("insertAddress", addrApi.InsertAddress)        // 添加收货地址
		addrRouter.PUT("updateAddress", addrApi.UpdateAddress)         // 更新收货地址
		addrRouter.DELETE("deleteAddress/:id", addrApi.DeleteAddress)  // 删除收货地址
		addrRouter.GET("addressDetailById", addrApi.AddressDetailById) // 地址详情
	}

	return addrRouter
}

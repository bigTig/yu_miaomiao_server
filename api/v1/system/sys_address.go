package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"yuyu/global"
	"yuyu/model/common/response"
	systemReq "yuyu/model/system/request"
	"yuyu/utils"
)

type AddressApi struct{}

// AddressList
// @Tags      Address
// @Summary   获取收货地址列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=response.PageResult{list=[]system.SysAddress,}} ""
// @Router    /address/addressList [get]
func (address *AddressApi) AddressList(c *gin.Context) {
	uuid := utils.GetUserUuid(c)
	list, total, err := addressService.AddressList(uuid)

	if err != nil {
		global.GvaLog.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:  list,
		Total: total,
	}, "收货地址获取成功", c)
}

// InsertAddress
// @Tags      Address
// @Summary   添加收货地址
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.InsertAddressParams true  " "
// @Success   200  {object}  response.Response{msg=string} ""
// @Router    /address/insertAddress [post]
func (address *AddressApi) InsertAddress(c *gin.Context) {
	var addr systemReq.InsertAddressParams
	err := c.ShouldBindJSON(&addr)
	if err != nil {
		global.GvaLog.Error("参数json格式", zap.Error(err))
		response.FailWithBadRequest("参数json格式", c)
		return
	}
	err = utils.Verify(addr, utils.InsertAddressVerify)
	if err != nil {
		global.GvaLog.Error(err.Error(), zap.Error(err))
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	uuid := utils.GetUserUuid(c)
	err = addressService.InsertAddress(addr, uuid)
	if err != nil {
		global.GvaLog.Error("添加失败!", zap.Error(err))
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	response.OkWithMessage("添加成功", c)
}

// UpdateAddress
// @Tags      Address
// @Summary   更新收货地址
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.UpdateAddressParams true  " "
// @Success   200  {object}  response.Response{data=bool, msg=string} ""
// @Router    /address/updateAddress [put]
func (address *AddressApi) UpdateAddress(c *gin.Context) {
	var addr systemReq.UpdateAddressParams
	err := c.ShouldBindJSON(&addr)
	if err != nil {
		response.FailWithBadRequest(err.Error()+"参数json格式", c)
		return
	}
	err = utils.Verify(addr, utils.UpdateAddressVerify)
	if err != nil {
		response.FailWithBadRequest(err.Error(), c)
		return
	}
	uuid := utils.GetUserUuid(c)
	err = addressService.UpdateAddress(&addr, uuid)

	if err != nil {
		global.GvaLog.Error("更新失败!", zap.Error(err))
		response.FailWithInternalServerError("更新失败"+err.Error(), c)
		return
	}

	response.OkWithMessage("地址更新成功", c)
}

// DeleteAddress
// @Tags      Address
// @Summary   删除收货地址
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id  path string true  " "
// @Success   200  {object}  response.Response{data=bool} ""
// @Router    /address/deleteAddress/:id [delete]
func (address *AddressApi) DeleteAddress(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		response.FailWithBadRequest("id 不能为空", c)
		return
	}

	uuid := utils.GetUserUuid(c)
	err := addressService.DeleteAddress(id, uuid)

	if err != nil {
		global.GvaLog.Error("删除失败!", zap.Error(err))
		response.FailWithInternalServerError("删除失败, "+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// AddressDetailById
// @Tags      Address
// @Summary   获取收货地址详情
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id  query string true  " "
// @Success   200  {object}  response.Response{data=system.SysAddress}
// @Router    /address/addressDetailById [get]
func (address *AddressApi) AddressDetailById(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		response.FailWithBadRequest("id 不能为空", c)
		return
	}

	addr, err := addressService.AddressDetailById(id)
	if err != nil {
		global.GvaLog.Error(err.Error()+" 获取失败!", zap.Error(err))
		response.FailWithInternalServerError(err.Error()+" 获取失败!", c)
		return
	}
	response.OkWithDetailed(addr, "获取成功", c)
}

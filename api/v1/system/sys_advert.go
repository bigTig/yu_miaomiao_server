package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"yuyu/global"
	"yuyu/model/common/request"
	"yuyu/model/common/response"
	systemReq "yuyu/model/system/request"
	"yuyu/utils"
)

// AdvertList
// @Tags      Base
// @Summary   获取轮播图
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @params data body request.PageInfo true  "页码, 每页大小"
// @Success   200  {object}  response.Response{data=response.PageResult{list=[]system.SysAdvert,}} ""
// @Router    /base/advertList [post]
func (b *BaseApi) AdvertList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithBadRequest(err.Error()+"参数为json格式", c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	list, total, err := advertService.AdvertList(pageInfo)

	if err != nil {
		global.GvaLog.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "轮播图获取成功", c)
}

// InsertAdvert
// @Tags      Base
// @Summary   添加轮播图
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.InsertAdvert true  " "
// @Success   200  {object}  response.Response{msg=string} ""
// @Router    /base/insertAdvert [post]
func (b *BaseApi) InsertAdvert(c *gin.Context) {

	var adv systemReq.InsertAdvert
	err := c.ShouldBindJSON(&adv)
	if err != nil {
		response.FailWithBadRequest("参数json格式", c)
		return
	}
	err = utils.Verify(adv, utils.InsertAdvertVerify)
	if err != nil {
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	err = advertService.InsertAdvert(&adv)
	if err != nil {
		global.GvaLog.Error("添加轮播图失败!", zap.Error(err))
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	response.OkWithMessage("添加轮播图成功", c)
}

// UpdateAdvert
// @Tags      Base
// @Summary   更新轮播图
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.UpdateAdvert true  " "
// @Success   200  {object}  response.Response{data=bool, msg=string} ""
// @Router    /base/updateAdvert [put]
func (b *BaseApi) UpdateAdvert(c *gin.Context) {
	var adv systemReq.UpdateAdvert
	err := c.ShouldBindJSON(&adv)
	if err != nil {
		response.FailWithBadRequest("参数json格式", c)
		return
	}
	err = utils.Verify(adv, utils.UpdateAdvertVerify)
	if err != nil {
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	err = advertService.UpdateAdvert(&adv)

	if err != nil {
		global.GvaLog.Error("更新失败!", zap.Error(err))
		response.FailWithInternalServerError("更新失败"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新轮播图成功", c)
}

// DeleteAdvert
// @Tags      Base
// @Summary   删除轮播图
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id  path string true  " "
// @Success   200  {object}  response.Response{data=bool, msg=string} ""
// @Router    /base/deleteAdvert/:id [delete]
func (b *BaseApi) DeleteAdvert(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		response.FailWithBadRequest("id 不能为空", c)
		return
	}
	err := advertService.DeleteAdvert(id)

	if err != nil {
		global.GvaLog.Error("删除失败!", zap.Error(err))
		response.FailWithInternalServerError("删除失败, "+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

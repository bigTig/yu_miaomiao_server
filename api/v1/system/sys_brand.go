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

// BrandList
// @Tags      Base
// @Summary   获取品牌列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @param data body request.PageInfo true  "页码, 每页大小"
// @Success   200  {object}  response.Response{data=response.PageResult{list=[]system.SysAdvert,}} ""
// @Router    /base/brandtList [post]
func (b *BaseApi) BrandList(c *gin.Context) {
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

	list, total, err := brandService.BrandList(pageInfo)

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
	}, "品牌获取成功", c)
}

// InsertBrand
// @Tags      Base
// @Summary   添加品牌
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.InsertBrandReq true  " "
// @Success   200  {object}  response.Response{msg=string} ""
// @Router    /base/insertBrand [post]
func (b *BaseApi) InsertBrand(c *gin.Context) {
	var brand systemReq.InsertBrandReq
	err := c.ShouldBindJSON(&brand)
	if err != nil {
		response.FailWithBadRequest(err.Error()+"参数json格式", c)
		return
	}
	err = utils.Verify(brand, utils.InsertBrandVerify)
	if err != nil {
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	err = brandService.InsertBrand(&brand)
	if err != nil {
		global.GvaLog.Error("添加品牌失败!", zap.Error(err))
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	response.OkWithMessage("添加品牌成功", c)
}

// UpdateBrand
// @Tags      Base
// @Summary   更新品牌
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.UpdateBrandReq true  " "
// @Success   200  {object}  response.Response{data=bool, msg=string} ""
// @Router    /base/updateBrand [put]
func (b *BaseApi) UpdateBrand(c *gin.Context) {
	var brand systemReq.UpdateBrandReq
	err := c.ShouldBindJSON(&brand)
	if err != nil {
		response.FailWithBadRequest(err.Error()+"参数json格式", c)
		return
	}
	err = utils.Verify(brand, utils.UpdateBrandVerify)
	if err != nil {
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	err = brandService.UpdateBrand(&brand)

	if err != nil {
		global.GvaLog.Error("更新失败!", zap.Error(err))
		response.FailWithInternalServerError("更新失败"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新品牌成功", c)
}

// DeleteBrand
// @Tags      Base
// @Summary   删除品牌
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id  path string true  " "
// @Success   200  {object}  response.Response{data=bool, msg=string} ""
// @Router    /base/deleteBrand/:id [delete]
func (b *BaseApi) DeleteBrand(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		response.FailWithBadRequest("id 不能为空", c)
		return
	}
	err := brandService.DeleteBrand(id)

	if err != nil {
		global.GvaLog.Error("删除失败!", zap.Error(err))
		response.FailWithInternalServerError("删除失败, "+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

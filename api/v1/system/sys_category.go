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

// CategoryList
// @Tags      Base
// @Summary   获取分类列表数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @param data body request.PageInfo true  "页码, 每页大小"
// @Success   200  {object}  response.Response{data=response.PageResult{list=[]system.SysCategory,}} ""
// @Router    /base/categoryList [post]
func (b *BaseApi) CategoryList(c *gin.Context) {
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

	list, total, err := categoryService.CategoryList(pageInfo)

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
	}, "类目获取成功", c)
}

// InsertCategory
// @Tags      Base
// @Summary   添加类目
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.InsertCateReq true  " "
// @Success   200  {object}  response.Response{msg=string} ""
// @Router    /base/insertCategory [post]
func (b *BaseApi) InsertCategory(c *gin.Context) {
	var cate systemReq.InsertCateReq
	err := c.ShouldBindJSON(&cate)
	if err != nil {
		global.GvaLog.Error(err.Error() + " 参数json格式")
		response.FailWithBadRequest(err.Error()+" 参数json格式", c)
		return
	}
	err = utils.Verify(cate, utils.InsertCateVerify)
	if err != nil {
		global.GvaLog.Error(err.Error())
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	err = categoryService.InsertCategory(&cate)
	if err != nil {
		global.GvaLog.Error("添加类目失败失败!", zap.Error(err))
		response.FailWithBadRequest(err.Error()+" 添加类目失败失败", c)
		return
	}

	response.OkWithMessage("添加类目成功", c)
}

// UpdateCategory
// @Tags      Base
// @Summary   更新类目
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.UpdateCateReq true  " "
// @Success   200  {object}  response.Response{data=bool, msg=string} ""
// @Router    /base/updateCategory [put]
func (b *BaseApi) UpdateCategory(c *gin.Context) {
	var cate systemReq.UpdateCateReq
	err := c.ShouldBindJSON(&cate)
	if err != nil {
		global.GvaLog.Error(err.Error() + "参数json格式")
		response.FailWithBadRequest("参数json格式", c)
		return
	}
	err = utils.Verify(cate, utils.UpdateCateVerify)
	if err != nil {
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	err = categoryService.UpdateCategory(&cate)

	if err != nil {
		global.GvaLog.Error("更新失败!", zap.Error(err))
		response.FailWithInternalServerError("更新失败"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新类目成功", c)
}

// DeleteCategory
// @Tags      Base
// @Summary   删除类目
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id  path string true  " "
// @Success   200  {object}  response.Response{data=bool, msg=string} ""
// @Router    /base/deleteCategory/:id [delete]
func (b *BaseApi) DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		global.GvaLog.Error("id 不能为空")
		response.FailWithBadRequest("id 不能为空", c)
		return
	}
	err := categoryService.DeleteCategory(id)

	if err != nil {
		global.GvaLog.Error("删除失败!", zap.Error(err))
		response.FailWithInternalServerError("删除失败, "+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

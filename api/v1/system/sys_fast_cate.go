package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"yuyu/global"
	"yuyu/model/common/response"
	systemReq "yuyu/model/system/request"
	"yuyu/utils"
)

type FastCateApi struct{}

// FastCateList
// @Tags      Fast
// @Summary   获取禁食分类列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=response.PageResult{list=[]system.SysHealthNews,}} ""
// @Router    /fast/fastCateList [get]
func (f *FastCateApi) FastCateList(c *gin.Context) {

	list, total, err := fastCateService.FastCateList()

	if err != nil {
		global.GvaLog.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:  list,
		Total: total,
	}, "禁食分类获取成功", c)
}

// InsertFastCate
// @Tags      Fast
// @Summary   添加禁食分类
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.InsertFastCateReq true  " "
// @Success   200  {object}  response.Response{msg=string} ""
// @Router    /fast/insertFastCate [post]
func (f *FastCateApi) InsertFastCate(c *gin.Context) {
	var cate systemReq.InsertFastCateReq
	err := c.ShouldBindJSON(&cate)
	if err != nil {
		global.GvaLog.Error("参数json格式", zap.Error(err))
		response.FailWithBadRequest("参数json格式", c)
		return
	}
	err = utils.Verify(cate, utils.InsertFastCateVerify)
	if err != nil {
		global.GvaLog.Error(err.Error(), zap.Error(err))
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	err = fastCateService.InsertFastCate(&cate)
	if err != nil {
		global.GvaLog.Error("添加失败!", zap.Error(err))
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	response.OkWithMessage("添加成功", c)
}

// UpdateFastCate
// @Tags      Fast
// @Summary   更新禁食分类
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.UpdateFastCateReq true  " "
// @Success   200  {object}  response.Response{data=bool, msg=string} ""
// @Router    /fast/updateFastCate [put]
func (f *FastCateApi) UpdateFastCate(c *gin.Context) {
	var cate systemReq.UpdateFastCateReq
	err := c.ShouldBindJSON(&cate)
	if err != nil {
		response.FailWithBadRequest(err.Error()+"参数json格式", c)
		return
	}
	err = utils.Verify(cate, utils.UpdateFastCateVerify)
	if err != nil {
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	err = fastCateService.UpdateFastCate(&cate)

	if err != nil {
		global.GvaLog.Error("更新失败!", zap.Error(err))
		response.FailWithInternalServerError("更新失败"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// DeleteFastCate
// @Tags      Fast
// @Summary   删除禁食分类
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id  path string true  " "
// @Success   200  {object}  response.Response{data=bool, msg=string} ""
// @Router    /fast/deleteFastCate/:id [delete]
func (f *FastCateApi) DeleteFastCate(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		response.FailWithBadRequest("id 不能为空", c)
		return
	}
	err := fastCateService.DeleteFastCate(id)

	if err != nil {
		global.GvaLog.Error("删除失败!", zap.Error(err))
		response.FailWithInternalServerError("删除失败, "+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

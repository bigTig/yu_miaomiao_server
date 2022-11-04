package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"yuyu/global"
	"yuyu/model/common/response"
	systemReq "yuyu/model/system/request"
	"yuyu/utils"
)

type FastApi struct{}

// FastCateList
// @Tags      Fast
// @Summary   获取禁食分类列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=response.PageResult{list=[]system.SysFastCate,}} ""
// @Router    /fast/fastCateList [get]
func (f *FastApi) FastCateList(c *gin.Context) {

	list, total, err := fastService.FastCateList()

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
func (f *FastApi) InsertFastCate(c *gin.Context) {
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

	err = fastService.InsertFastCate(&cate)
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
func (f *FastApi) UpdateFastCate(c *gin.Context) {
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

	err = fastService.UpdateFastCate(&cate)

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
func (f *FastApi) DeleteFastCate(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		response.FailWithBadRequest("id 不能为空", c)
		return
	}
	err := fastService.DeleteFastCate(id)

	if err != nil {
		global.GvaLog.Error("删除失败!", zap.Error(err))
		response.FailWithInternalServerError("删除失败, "+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// FastList
// @Tags      Fast
// @Summary   获取禁食列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=response.PageResult{list=[]system.SysFast,}} ""
// @Router    /fast/fastList [post]
func (f *FastApi) FastList(c *gin.Context) {
	var pageInfo systemReq.FastListReq
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithBadRequest(err.Error()+"参数为json格式", c)
		return
	}
	err = utils.Verify(pageInfo, utils.FastListVerify)
	if err != nil {
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	list, total, err := fastService.FastList(&pageInfo)

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
	}, "获取成功", c)
}

// FastDetailById
// @Tags      Fast
// @Summary   根据id获取禁食详情
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id  query string true  " "
// @Success   200  {object}  response.Response{data=system.SysFast} ""
// @Router    /fast/fastDetailById [get]
func (f *FastApi) FastDetailById(c *gin.Context) {
	id := c.Query("id")
	detail, err := fastService.FastDetailById(id)

	if err != nil {
		global.GvaLog.Error("获取失败!", zap.Error(err))
		response.FailWithInternalServerError(err.Error()+",获取失败", c)
		return
	}

	response.OkWithDetailed(detail, "获取成功", c)
}

// InsertFast
// @Tags      Fast
// @Summary   添加禁食
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.InsertFastReq true  " "
// @Success   200  {object}  response.Response{msg=string} ""
// @Router    /fast/insertFast [post]
func (f *FastApi) InsertFast(c *gin.Context) {
	var fast systemReq.InsertFastReq
	err := c.ShouldBindJSON(&fast)
	if err != nil {
		global.GvaLog.Error("参数json格式", zap.Error(err))
		response.FailWithBadRequest("参数json格式", c)
		return
	}
	err = utils.Verify(fast, utils.InsertFastVerify)
	if err != nil {
		global.GvaLog.Error(err.Error(), zap.Error(err))
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	err = fastService.InsertFast(&fast)
	if err != nil {
		global.GvaLog.Error("添加失败!", zap.Error(err))
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	response.OkWithMessage("添加成功", c)
}

// UpdateFast
// @Tags      Fast
// @Summary   更新禁食
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.UpdateFastReq true  " "
// @Success   200  {object}  response.Response{data=bool, msg=string} ""
// @Router    /fast/updateFast [put]
func (f *FastApi) UpdateFast(c *gin.Context) {
	var fast systemReq.UpdateFastReq
	err := c.ShouldBindJSON(&fast)
	if err != nil {
		response.FailWithBadRequest(err.Error()+"参数json格式", c)
		return
	}
	err = utils.Verify(fast, utils.UpdateFastVerify)
	if err != nil {
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	err = fastService.UpdateFast(&fast)

	if err != nil {
		global.GvaLog.Error("更新失败!", zap.Error(err))
		response.FailWithInternalServerError("更新失败"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// DeleteFast
// @Tags      Fast
// @Summary   删除禁食
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id  path string true  " "
// @Success   200  {object}  response.Response{data=bool, msg=string} ""
// @Router    /fast/deleteFast/:id [delete]
func (f *FastApi) DeleteFast(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		response.FailWithBadRequest("id 不能为空", c)
		return
	}
	err := fastService.DeleteFast(id)

	if err != nil {
		global.GvaLog.Error("删除失败!", zap.Error(err))
		response.FailWithInternalServerError("删除失败, "+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

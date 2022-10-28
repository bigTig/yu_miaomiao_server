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

type HealthNewsApi struct{}

// HealthNewsList
// @Tags      Health
// @Summary   获取新闻资讯列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @param data body request.PageInfo true  "页码, 每页大小"
// @Success   200  {object}  response.Response{data=response.PageResult{list=[]system.SysHealthNews,}} ""
// @Router    /health/healthNewsList [post]
func (h *HealthNewsApi) HealthNewsList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		global.GvaLog.Error(err.Error()+"参数为json格式", zap.Error(err))
		response.FailWithBadRequest(err.Error()+"参数为json格式", c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		global.GvaLog.Error(err.Error(), zap.Error(err))
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	list, total, err := healthNewsService.HealthNewsList(pageInfo)

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
	}, "新闻资讯获取成功", c)
}

// InsertHealthNew
// @Tags      Health
// @Summary   添加新闻资讯
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.InsertHealthNewReq true  " "
// @Success   200  {object}  response.Response{msg=string} ""
// @Router    /base/insertHealthNew [post]
func (h *HealthNewsApi) InsertHealthNew(c *gin.Context) {
	var health systemReq.InsertHealthNewReq
	err := c.ShouldBindJSON(&health)
	if err != nil {
		response.FailWithBadRequest("参数json格式", c)
		return
	}
	err = utils.Verify(health, utils.InsertHealthNewVerify)
	if err != nil {
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	err = healthNewsService.InsertHealthNew(&health)
	if err != nil {
		global.GvaLog.Error("新闻资讯添加失败!", zap.Error(err))
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	response.OkWithMessage("新闻资讯添加成功", c)
}

// UpdateHealthNew
// @Tags      Health
// @Summary   更新新闻资讯
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.UpdateHealthNewReq true  " "
// @Success   200  {object}  response.Response{data=bool, msg=string} ""
// @Router    /base/updateHealthNew [put]
func (h *HealthNewsApi) UpdateHealthNew(c *gin.Context) {
	var health systemReq.UpdateHealthNewReq
	err := c.ShouldBindJSON(&health)
	if err != nil {
		response.FailWithBadRequest(err.Error()+"参数json格式", c)
		return
	}
	err = utils.Verify(health, utils.UpdateHealthNewVerify)
	if err != nil {
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	err = healthNewsService.UpdateHealthNew(&health)

	if err != nil {
		global.GvaLog.Error("更新失败!", zap.Error(err))
		response.FailWithInternalServerError("更新失败"+err.Error(), c)
		return
	}

	response.OkWithMessage("新闻资讯更新成功", c)
}

// DeleteHealthNew
// @Tags      Health
// @Summary   删除新闻
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id  path string true  " "
// @Success   200  {object}  response.Response{data=bool, msg=string} ""
// @Router    /base/deleteHealthNew/:id [delete]
func (h *HealthNewsApi) DeleteHealthNew(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		response.FailWithBadRequest("id 不能为空", c)
		return
	}
	err := healthNewsService.DeleteHealthNew(id)

	if err != nil {
		global.GvaLog.Error("删除失败!", zap.Error(err))
		response.FailWithInternalServerError("删除失败, "+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// HealthNewDetail
// @Tags      Health
// @Summary   获取新闻资讯详情
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id  query string true  " "
// @Success   200  {object}  response.Response{data=system.SysHealthNews, msg=string} ""
// @Router    /base/HealthNewDetail [get]
func (h *HealthNewsApi) HealthNewDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		response.FailWithBadRequest("id 不能为空", c)
		return
	}

	health, err := healthNewsService.HealthNewDetail(id)
	if err != nil {
		global.GvaLog.Error(err.Error()+" 获取失败!", zap.Error(err))
		response.FailWithInternalServerError(err.Error()+" 获取失败!", c)
		return
	}
	response.OkWithDetailed(health, "获取成功", c)
}

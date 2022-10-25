package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"yuyu/global"
	"yuyu/model/common/request"
	"yuyu/model/common/response"
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

package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"yuyu/global"
	"yuyu/model/common/response"
	"yuyu/model/system"
	"yuyu/utils"
)

type JwtApi struct{}

// JsonInBlacklist
// @Tags      Jwt
// @Summary   jwt加入黑名单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{msg=string}  "jwt加入黑名单"
// @Router    /jwt/jsonInBlacklist [post]
func (j *JwtApi) JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := system.JwtBlacklist{Jwt: token}
	jwt.CreatedTime = utils.SetCreatedTime()
	jwt.UpdatedTime = utils.SetUpdatedTime()

	err := jwtService.JsonInBlacklist(&jwt)
	if err != nil {
		global.GvaLog.Error("jwt 作废失败!", zap.Error(err))
		response.FailWithMessage("jwt 作废失败!", c)
		return
	}
	response.OkWithMessage("jwt 作废成功", c)
}

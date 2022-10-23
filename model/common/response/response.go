package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 200
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, nil, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "查询成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func FailWithBadRequest(message string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, Response{
		Code: http.StatusBadRequest,
		Data: nil,
		Msg:  message,
	})
}

func FailWithUnauthorized(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		Code: http.StatusUnauthorized,
		Data: nil,
		Msg:  message,
	})
}

func FailWithNotFound(message string, c *gin.Context) {
	c.JSON(http.StatusNotFound, Response{
		Code: http.StatusNotFound,
		Data: nil,
		Msg:  message,
	})
}

func FailWithInternalServerError(message string, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Response{
		Code: http.StatusInternalServerError,
		Data: nil,
		Msg:  message,
	})
}

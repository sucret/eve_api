package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 定义一个统一的返回值结构体
type Response struct {
	ErrorCode int         `json:"error_code"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
}

// Success 成功返回
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		0,
		data,
		"ok",
	})
}

func ValidateFailed(c *gin.Context, msg string) {
	Error(c, ValidateError, msg)
}

func PermissionDenied(c *gin.Context) {
	Error(c, HaveNoPermission, ErrorMap[HaveNoPermission])
}

func BusinessFail(c *gin.Context, msg string) {
	Error(c, BusinessError, msg)
}

// Error 失败返回
func Error(c *gin.Context, errorCode int, msg string) {
	c.JSON(http.StatusOK, Response{
		errorCode,
		nil,
		msg,
	})
}

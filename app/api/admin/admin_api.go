package admin

import (
	"eve/app/request"
	"eve/app/response"
	"eve/app/service"
	"github.com/gin-gonic/gin"
)

type adminApi struct{}

func (a *adminApi) Profile(c *gin.Context) {
	admin := service.AdminService.Profile()

	response.Success(c, admin)
}

// Save 新增/编辑管理员信息
func (a *adminApi) Save(c *gin.Context) {
	var admin request.SaveAdmin
	if err := c.ShouldBind(&admin); err != nil {
		response.ValidateFailed(c, request.GetErrorMsg(admin, err))
	}
}

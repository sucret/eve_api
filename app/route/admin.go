package route

import (
	"eve/app/api/admin"
	"github.com/gin-gonic/gin"
)

func genAdminRouter(rg *gin.RouterGroup) {
	rg.GET("/profile", admin.AdminApi.Profile)
}

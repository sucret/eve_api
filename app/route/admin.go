package route

import (
	"eve/app/api/admin"
	"eve/app/middleware"
	"github.com/gin-gonic/gin"
)

func genAdminRouter(rg *gin.RouterGroup) {
	rg.Use(middleware.CorsMiddleware)

	// 登陆接口，
	rg.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "login"})
	})

	// ------ 以上的接口不走权限校验 ------

	// 注册权限验证中间件
	rg.Use(middleware.PermissionMiddleware)

	rg.GET("/profile", admin.AdminApi.Profile)
	rg.POST("/save", admin.AdminApi.Save)
}

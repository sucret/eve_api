package route

import (
	"github.com/gin-gonic/gin"
)

func genAdminRouter(rg *gin.RouterGroup) {
	rg.GET("/a", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "a"})
	})
}

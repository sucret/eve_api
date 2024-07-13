package route

import "github.com/gin-gonic/gin"

func genAppRouter(rg *gin.RouterGroup) {

	rg.GET("/b", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "b"})
	})
}

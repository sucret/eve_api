package middleware

import (
	"eve/app/response"
	"github.com/gin-gonic/gin"
)

func PermissionMiddleware(c *gin.Context) {
	// 获取路由
	//path := c.FullPath()
	// 使用路由和当前登陆用户去判断是否有当前接口的访问权限
	if isHasPermission() {
		c.Next()
	} else {
		response.PermissionDenied(c)
		c.Abort()
	}

}

// 权限验证，这里根据自己的业务逻辑去实现
func isHasPermission() bool {
	return true
}

package route

import (
	"github.com/gin-gonic/gin"
)

type AppRouter struct{}

func (*AppRouter) AddRoute(e *gin.Engine) {
	genAdminRouter(e.Group("/admin"))
	genAppRouter(e.Group("/app"))
}

func New() *AppRouter {
	return &AppRouter{}
}

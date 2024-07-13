package server

import (
	"context"
	"eve/internal/server/route"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Http struct {
	engine *gin.Engine // gin框架的核心组件
	port   string      // http服务端口号
}

// New 实例化http对象
func New() *Http {
	entity := &Http{
		engine: gin.New(),
		port:   ":8082",
	}

	return entity
}

// GenRouter 注册路由
// 这里需要调用外部的route库，所以先定义一个Interface，并在外部实现Interface
func (h *Http) GenRouter(r route.RouterGeneratorInterface) {
	r.AddRoute(h.engine)
}

// Run 启动服务
func (h *Http) Run() {
	srv := &http.Server{
		Addr:    h.port,
		Handler: h.engine,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
	}()

	h.ListenSignal(srv)
}

func (h *Http) ListenSignal(srv *http.Server) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("shutdown server")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown ")
	}
}

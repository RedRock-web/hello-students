// @program: hello-students
// @author: edte
// @create: 2020-08-02 23:18
package router

import (
	"github.com/gin-gonic/gin"

	"hello-students/log"
	"hello-students/router/handler"
	"hello-students/router/middleware"
)

// InitRouter 初始化路由
func InitRouter() {
	r := gin.Default()

	setupRouter(r)

	err := r.Run()
	if err != nil {
		log.Begin().Fatalf("failed to init router")
	}
}

// setupRouter 设置路由
func setupRouter(r *gin.Engine) {
	// log record，解决跨域
	r.Use(middleware.Cors(), middleware.LoggerToFile())

	r.GET("/api/visitor", handler.Visitor)
	r.GET("/api/college", handler.College)
	r.GET("/api/dorm", handler.Dorm)
	r.GET("/api/canteen", handler.Canteen)
	r.GET("/api/expand", handler.Expand)
}

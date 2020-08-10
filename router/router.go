// @program: hello-students
// @author: edte
// @create: 2020-08-02 23:18
package router

import (
	"log"

	"github.com/gin-gonic/gin"

	"hello-students/router/handler"
	"hello-students/router/middleware"
)

func InitRouter() {
	r := gin.Default()

	setupRouter(r)

	err := r.Run()
	if err != nil {
		log.Fatalf("failed to init router")
	}

}

func setupRouter(r *gin.Engine) {
	// log record，解决跨域
	r.Use(middleware.Cors())

	r.GET("/api/visitor", handler.Visitor)
	r.GET("/api/college", handler.College)
	r.GET("/api/dorm", handler.Dorm)
	r.GET("/api/canteen")
}

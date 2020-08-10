// @program: hello-students
// @author: edte
// @create: 2020-08-10 23:14
package handler

import (
	"github.com/gin-gonic/gin"

	"hello-students/router/response"
	"hello-students/service"
)

// Dorm
func Dorm(c *gin.Context) {
	var d service.DormForm

	if err := c.ShouldBindJSON(&d); err != nil {
		response.FormError(c)
		return
	}

}

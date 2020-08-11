// @program: hello-students
// @author: edte
// @create: 2020-08-11 16:33
package handler

import (
	"github.com/gin-gonic/gin"

	"hello-students/router/response"
	"hello-students/service"
)

// Link
func Link(c *gin.Context) {
	var l service.LinkForm

	if err := c.ShouldBindJSON(&l); err != nil {
		response.FormError(c)
		return
	}
}

// @program: hello-students
// @author: edte
// @create: 2020-08-11 16:32
package handler

import (
	"github.com/gin-gonic/gin"

	"hello-students/router/response"
	"hello-students/service"
)

// Expand
func Expand(c *gin.Context) {
	var e service.ExpandForm

	if err := c.ShouldBindJSON(&e); err != nil {
		response.FormError(c)
		return
	}

	introduction := service.GetExpandIntroduction(e.Name)

	pictures := service.GetExpandPictures(e.Name)

	data := gin.H{
		"pictures":     pictures,
		"introduction": introduction,
	}

	response.OkWithData(c, data)
}

// @program: hello-students
// @author: edte
// @create: 2020-08-11 16:06
package handler

import (
	"github.com/gin-gonic/gin"

	"hello-students/router/response"
	"hello-students/service"
)

// Canteen
func Canteen(c *gin.Context) {
	var cf service.CanteenForm

	if err := c.ShouldBind(&cf); err != nil {
		response.FormError(c)
		return
	}
	introduction := service.GetCanteenIntroduction(cf.Name)

	pictures := service.GetCanteenPictures(cf.Name)

	data := gin.H{
		"pictures":     pictures,
		"introduction": introduction,
	}

	response.OkWithData(c, data)
}

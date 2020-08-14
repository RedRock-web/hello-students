// @program: hello-students
// @author: edte
// @create: 2020-08-14 20:37
package handler

import (
	"github.com/gin-gonic/gin"

	"hello-students/router/response"
	"hello-students/service"
)

// Judge
func Judge(c *gin.Context) {
	var j service.JudgeForm

	if err := c.ShouldBind(&j); err != nil {
		response.FormError(c)
		return
	}

	var data []string

	if j.Type == "region" {
		data = service.GetMajorsByRegion(j.Name)
	} else if j.Type == "major" {
		data = service.GetRegionsByMajor(j.Name)
	} else {
		response.FormError(c)
	}

	response.OkWithData(c, gin.H{"data": data})
}

// @program: hello-students
// @author: edte
// @create: 2020-08-10 19:50
package handler

import (
	"github.com/gin-gonic/gin"

	"hello-students/router/response"
	"hello-students/service"
)

// College
func College(c *gin.Context) {
	var cf service.CollegeForm

	if err := c.ShouldBind(&cf); err != nil {
		response.FormError(c)
		return
	}

	majors := service.GetCollegeMajors(cf.Name)

	introduction := service.GetCollegeIntroduction(cf.Name)

	majorNumber := service.GetCollegeMajorNumber(cf.Name)

	sexRadio := service.GetCollegeSexRatio(cf.Name)

	collegeNumber := service.GetCollegeNumber(cf.Name)

	regionNumber := service.GetCollegeRegionNumber(cf.Name, cf.Region)

	data := gin.H{
		"name":          cf.Name,
		"introduction":  introduction,
		"major_number":  majorNumber,
		"majors":        majors,
		"sex_ratio":     sexRadio,
		"region":        cf.Region,
		"region_number": regionNumber,
		"number":        collegeNumber,
	}

	response.OkWithData(c, data)
}

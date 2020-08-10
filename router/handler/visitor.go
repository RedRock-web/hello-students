// @program: hello-students
// @author: edte
// @create: 2020-08-10 19:49
package handler

import (
	"log"

	"github.com/gin-gonic/gin"

	"hello-students/router/response"
	"hello-students/service"
)

// Visitor
func Visitor(c *gin.Context) {
	var v service.VisitorForm

	if err := c.ShouldBindJSON(&v); err != nil {
		response.FormError(c)
		return
	}

	v.Ip = c.ClientIP()

	err := service.AddVisitor(v)
	if err != nil {
		log.Printf("faield to add visitor:%s", err)
	}

	number, err := service.GetVisitorsNumber()
	if err != nil {
		log.Printf("faield to get visitor number:%s", err)
		return
	}

	collegeName := service.GetCollegeName(v.Major)

	sexRatio := service.GetGenderRatio(v.Gender)

	genderNumber := service.GetGenderNumber(v.Gender)

	regionNumber := service.GetRegionNumber(v.Region)

	genderDesc := service.GetRandomGenderDescription(v.Gender)

	data := gin.H{
		"college":         collegeName,
		"region":          v.Region,
		"fellows_number":  regionNumber,
		"sex_ratio":       sexRatio,
		"gender":          genderDesc,
		"gender_number":   genderNumber,
		"visitors_number": number,
	}

	response.OkWithData(c, data)
}

// @program: hello-students
// @author: edte
// @create: 2020-08-03 14:32
package handler

import (
	"log"

	"github.com/gin-gonic/gin"

	"hello-students/config"
	"hello-students/router/response"
	"hello-students/service"
)

func Visitors(c *gin.Context) {
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

	var data gin.H

	if v.Gender == "male" {
		data["gender"] = "male"
		data["gender_number"] = config.VisitorConfig.MaleNumber
		data["sex_ratio"] = config.VisitorConfig.MalePercent
	} else if v.Gender == "female" {
		data["gender"] = "female"
		data["gender_number"] = config.VisitorConfig.FemaleNumber
		data["sex_ratio"] = config.VisitorConfig.FemalePercent
	} else {
		response.FormError(c)
	}

	data["visitors_number"] = number
	data["region"] = v.Region


}

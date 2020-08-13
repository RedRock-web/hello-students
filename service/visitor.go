// @program: hello-students
// @author: edte
// @create: 2020-08-10 18:51
package service

import (
	"hello-students/model"
)

// VisitorForm
type VisitorForm struct {
	Gender string `form:"gender" binding:"required"`
	Region string `form:"region" binding:"required"`
	Major  string `form:"major" binding:"required"`
	Ip     string `form:"ip"`
}

// GetVisitorsNumber
func GetVisitorsNumber() (int, error) {
	return model.GetVisitorsNumber()
}

// AddVisitor
func AddVisitor(v VisitorForm) error {
	return model.AddVisitor(model.Visitor{
		IP:     v.Ip,
		Gender: v.Gender,
		Region: v.Region,
		Major:  v.Major,
	})
}

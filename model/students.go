// @program: hello-students
// @author: edte
// @create: 2020-08-02 23:12
package model

import (
	"github.com/jinzhu/gorm"
)

// Student 表示访客学生
type Student struct {
	gorm.Model
	IP     string // students ip
	Gender string // students gender
	Region string // students region
	Major  string // students major
}

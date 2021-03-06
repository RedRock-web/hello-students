// @program: hello-students
// @author: edte
// @create: 2020-08-02 17:56
package main

import (
	"hello-students/config"
	"hello-students/model"
	"hello-students/router"
)

func main() {
	config.InitConfig()
	model.InitModel()
	router.InitRouter()
}

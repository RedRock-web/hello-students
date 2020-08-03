// @program: hello-students
// @author: edte
// @create: 2020-08-02 23:09
package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"hello-students/config"
)

// DB
var DB *gorm.DB

// InitModel
func InitModel() {
	db, err := gorm.Open(config.DatabaseConfig.Type,
		fmt.Sprintf(
			"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			config.DatabaseConfig.User,
			config.DatabaseConfig.Password,
			config.DatabaseConfig.Host,
			config.DatabaseConfig.Port,
			config.DatabaseConfig.Name))

	if err != nil {
		log.Fatalf("failed to connect database:%v", err)
	}

	DB = db
	setTable()
	addDefaultData()
}

func setTable() {
	if DB.HasTable(&Visitor{}) {
		DB.AutoMigrate(&Visitor{})
	} else {
		DB.CreateTable(&Visitor{})
	}
}

func addDefaultData() {

}

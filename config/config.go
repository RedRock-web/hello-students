// @program: hello-students
// @author: edte
// @create: 2020-08-02 23:11
package config

import (
	"log"
	"os"
)

type databaseConfig struct {
	Name     string
	Type     string
	User     string
	Password string
	Host     string
	Port     int
}

// DatabaseConfig 用于数据库连接
var DatabaseConfig = databaseConfig{
	Name:     "hello_student",
	Type:     "mysql",
	User:     "root",
	Password: "mima",
	Host:     "127.0.0.1",
	Port:     3306,
}

type cookieConfig struct {
	Name     string
	Value    string
	MaxAge   int
	Path     string
	Domain   string
	Secure   bool
	HttpOnly bool
}

// CookieConfig
var CookieConfig = cookieConfig{
	Name:     "soidfjosd",
	Value:    "",
	Domain:   "localhost:8080",
	Path:     "/",
	MaxAge:   10000,
	Secure:   false,
	HttpOnly: true,
}

type logFileConfig struct {
	Path string
	Name string
}

// LogFileConfig 用于存日志信息
var LogFileConfig = logFileConfig{
	Path: "",
	Name: "log.txt",
}

// InitConfig
func InitConfig() {
	_, err := os.Create(LogFileConfig.Path + LogFileConfig.Name)
	if err != nil {
		log.Fatalf("failed init log file")
	}
}
// @program: hello-students
// @author: edte
// @create: 2020-08-02 23:11
package config

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
	Name:     "cloud_disk",
	Type:     "mysql",
	User:     "root",
	Password: "mima",
	Host:     "127.0.0.1",
	Port:     3306,
}

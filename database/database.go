// database/database.go
package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 是全局数据库连接
var DB *gorm.DB

// Connect 连接到数据库并设置全局 DB 变量
func Connect() error {
	dsn := "chenjianan:222333@tcp(127.0.0.1:3306)/register?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

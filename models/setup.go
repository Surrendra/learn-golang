package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:8889)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "golang")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&Product{})

	DB = database
}

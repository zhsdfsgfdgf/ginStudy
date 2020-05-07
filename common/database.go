package common

import (
	"fmt"
	model "ginStudy/models"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:zhaoheng@(47.105.83.40:3306)/ginStudy?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接失败", err.Error())
	}
	//自动建表
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}
func GetDB() *gorm.DB {
	return DB
}

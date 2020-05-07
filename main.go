package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(100;not null;unique)"`
	Password  string `gorm:"size:255;not null"`
}

func main() {
	db := InitDB()
	defer db.Close()
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		//获取参数
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")
		//数据验证
		if len(telephone) != 11 {
			ctx.JSON(422, gin.H{"code": 422, "msg": "手机号必须位为11位"})
			return
		}
		if len(password) < 6 {
			ctx.JSON(422, gin.H{"code": 422, "msg": "密码不得少于6位"})
			return
		}
		if len(name) == 0 {
			name = "random"
		}
		fmt.Println(name, telephone, password)
		//数据库判断
		if isTelephoneExist(db, telephone) {
			ctx.JSON(422, gin.H{"code": 422, "msg": "用户已经存在"})
			return
		}
		//创建用户
		newUser := User{
			Name:      name,
			Password:  password,
			Telephone: telephone,
		}
		db.Create(&newUser)
		ctx.JSON(422, gin.H{"code": 200, "msg": "注册成功"})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:zhaoheng@(47.105.83.40:3306)/ginStudy?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接失败", err.Error())
	}
	//自动建表
	db.AutoMigrate(&User{})
	return db
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

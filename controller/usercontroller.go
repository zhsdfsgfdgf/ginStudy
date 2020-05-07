package controller

import (
	"ginStudy/common"
	model "ginStudy/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()
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
	//数据库判断
	if isTelephoneExist(db, telephone) {
		ctx.JSON(422, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}
	//创建用户
	newUser := model.User{
		Name:      name,
		Password:  password,
		Telephone: telephone,
	}
	db.Create(&newUser)
	ctx.JSON(422, gin.H{"code": 200, "msg": "注册成功"})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

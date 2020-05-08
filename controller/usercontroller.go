package controller

import (
	"ginStudy/common"
	"ginStudy/dto"
	model "ginStudy/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()
	//获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 422, "msg": "手机号必须位为11位"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 422, "msg": "密码不得少于6位"})
		return
	}
	if len(name) == 0 {
		name = "random"
	}
	//数据库判断
	if isTelephoneExist(db, telephone) {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}
	//创建用户
	basedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "内部加密错误"})
	}
	newUser := model.User{
		Name:      name,
		Password:  string(basedPassword),
		Telephone: telephone,
	}
	db.Create(&newUser)
	ctx.JSON(http.StatusBadRequest, gin.H{"code": 200, "msg": "注册成功"})
}

func Login(ctx *gin.Context) {
	db := common.GetDB()
	//获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 422, "msg": "手机号必须位为11位"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 422, "msg": "密码不得少于6位"})
		return
	}
	//数据库判断
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 422, "msg": "该用户不存在"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "data": gin.H{"token": token}, "msg": "登录成功"})

}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(200, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})

}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

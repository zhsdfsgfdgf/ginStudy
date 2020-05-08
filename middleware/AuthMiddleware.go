package middleware

import (
	"net/http"
	"strings"

	"ginStudy/common"
	model "ginStudy/models"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//假设Token放在Header的Authorization中，并使用Bearer开头
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//去掉Bearer
		tokenString = tokenString[7:]
		//使用之前定义好的解析JWT的函数来解析它
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//验证通过,获取claim中的userid
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//如果用户存在,将user信息写入上下文
		//后续的处理函数可以用过c.Get("user")来获取当前请求的用户信息
		ctx.Set("user", user)
		ctx.Next()
	}
}

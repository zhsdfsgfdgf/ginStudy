package common

import (
	model "ginStudy/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//jwt加密的密钥
var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

//登录成功后用这个方法发放token
func ReleaseToken(user model.User) (string, error) {
	//token的过期时间
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			//token过期时间
			ExpiresAt: expireTime.Unix(),
			//发放时间
			IssuedAt: time.Now().Unix(),
			//发放人
			Issuer:  "zhaoheng",
			Subject: "user token",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}

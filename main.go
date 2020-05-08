package main

import (
	"ginStudy/common"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {
	//读取配置
	InitConfig()
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
}

func InitConfig() {
	//获取当前的工作目录
	workDir, _ := os.Getwd()
	//设置读取的文件名,文件类型,路径
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"mygin/common"
	"mygin/router"
	"os"
)

func main() {
	InitConfig()

	db := common.InitDB()
	defer db.Close()

	app := gin.Default()
	app = router.CollectRouter(app)

	appPort := viper.GetString("server.port")
	if appPort == "" {
		panic(":" + appPort)
	}
	app.Run() // listen and serve on 0.0.0.0:8080
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")

	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

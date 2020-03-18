package router

import (
	"github.com/gin-gonic/gin"
	"mygin/controller"
	"mygin/middleware"
)

func CollectRouter(app *gin.Engine) *gin.Engine {
	app.POST("/api/auth/register", controller.Register)

	//登录
	app.POST("/api/auth/login", controller.Login)

	//用户信息
	app.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return app
}

package router

import (
	"github.com/gin-gonic/gin"
	"mygin/controller"
)

func CollectRouter(app *gin.Engine) *gin.Engine {
	app.POST("/api/auth/register", controller.Register)

	return app
}

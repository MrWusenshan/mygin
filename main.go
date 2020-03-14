package main

import (
	"github.com/gin-gonic/gin"
	"mygin/common"
	"mygin/router"
)

func main() {
	db := common.InitDB()
	defer db.Close()

	app := gin.Default()
	app = router.CollectRouter(app)

	app.Run() // listen and serve on 0.0.0.0:8080
}

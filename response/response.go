package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code":    code,
		"data":    data,
		"message": msg,
	})
}

func SuccessResponse(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

func FailResponse(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
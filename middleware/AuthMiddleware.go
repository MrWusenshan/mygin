package middleware

import (
	"github.com/gin-gonic/gin"
	"mygin/common"
	"mygin/modle"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "权限不足",
			})

			ctx.Abort()

			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParesToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "权限不足",
			})
			ctx.Abort()
			return
		}

		//验证通过token
		userId := claims.UserID
		DB := common.GetDBEngine()
		var user modle.User
		DB.First(&user, userId)

		//用户
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "权限不足",
			})
			ctx.Abort()
			return
		}

		//用户存在 将信息写入 context
		ctx.Set("user", user)

		ctx.Next()
	}
}

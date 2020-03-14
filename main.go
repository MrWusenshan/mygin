package main

import (
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"mygin/orm"
	"net/http"
	"time"
)

func main() {
	db := orm.InitDB()
	defer db.Close()

	app := gin.Default()
	app.POST("/api/auth/register", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")
		//验证电话号码
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "电话号码必须为11位",
			})
			return
		}

		//验证密码是否合法
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "密码不能少于6位",
			})

			return
		}

		//如果没有传name,随机生成一个
		if len(name) == 0 {
			name = RandomString(10)
		}

		log.Println(name, telephone, password)
		//判断手机号是否存在
		if isTelephoneExist(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "该手机号已注册",
			})

			return
		}

		//创建用户
		newUser := orm.User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "注册成功",
		})

	})
	app.Run() // listen and serve on 0.0.0.0:8080
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user orm.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func RandomString(n int) string {
	var latters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = latters[rand.Intn(len(latters))]
	}

	return string(result)
}

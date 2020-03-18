package controller

import (
	"crypto/bcrypt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"mygin/common"
	"mygin/dto"
	"mygin/modle"
	"mygin/response"
	"mygin/util"
	"net/http"
)

func Register(ctx *gin.Context) {
	DB := common.GetDBEngine()

	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//验证电话号码
	if len(telephone) != 11 {

		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "电话号码必须为11位")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"code":    422,
		//	"message": "电话号码必须为11位",
		//})
		return
	}

	//验证密码是否合法
	if len(password) < 6 {

		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"code":    422,
		//	"message": "密码不能少于6位",
		//})

		return
	}

	//如果没有传name,随机生成一个
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name, telephone, password)
	//判断手机号是否存在
	if isTelephoneExist(DB, telephone) {

		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "该手机号已注册")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"code":    422,
		//	"message": "该手机号已注册",
		//})

		return
	}

	//创建用户
	//加密用户提交的密码
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {

		response.Response(ctx, http.StatusInternalServerError, 500, nil, "密码加密错误")
		//ctx.JSON(http.StatusInternalServerError, gin.H{
		//	"code":    500,
		//	"message": "密码加密错误",
		//})
		return
	}

	//
	newUser := modle.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)

	response.SuccessResponse(ctx, nil, "注册成功")
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code":    200,
	//	"message": "注册成功",
	//})

}

func Login(ctx *gin.Context) {
	DB := common.GetDBEngine()

	//获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	var user modle.User
	DB.Where("telephone = ?", telephone).First(&user)

	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户账户不存在")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"code":    422,
		//	"message": "用户账户不存在",
		//})
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.FailResponse(ctx, nil, "密码错误")
		//ctx.JSON(http.StatusBadRequest, gin.H{
		//	"code":    400,
		//	"message": "密码错误",
		//})
		return
	}

	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		//ctx.JSON(http.StatusInternalServerError, gin.H{
		//	"code":    500,
		//	"message": "系统异常",
		//})
		log.Printf("token generate error : %v", err)
		return
	}

	//返回结果
	response.SuccessResponse(ctx, gin.H{"token": token}, "登录成功")
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code":    200,
	//	"data":    gin.H{"token": token},
	//	"message": "登录成功",
	//})
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	response.SuccessResponse(ctx, gin.H{"user": dto.ToUserDto(user.(modle.User)),}, "")
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	"data": gin.H{
	//		"user": dto.ToUserDto(user.(modle.User)),
	//	},
	//})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user modle.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

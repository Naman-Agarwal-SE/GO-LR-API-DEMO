package controller

import (
	"golang-demo/data/request"
	"golang-demo/service"

	"github.com/gin-gonic/gin"
)

func Sott(ctx *gin.Context) {
	apiSecret := ctx.Query("apiSecret")
	apiKey := ctx.Query("apiKey")
	data, statusCode := service.Sott(apiSecret, apiKey)

	ctx.JSON(statusCode, data)
}

func Register(ctx *gin.Context) {
	registerBody := request.RegisterRequest{}
	apiKey := ctx.Query("apiKey")
	sott := ctx.GetHeader("X-LoginRadius-Sott")
	ctx.ShouldBindJSON(&registerBody)
	data, statusCode := service.Register(apiKey, &registerBody, sott)

	ctx.JSON(statusCode, data)
}

func EmailVerify(ctx *gin.Context) {
	verificationtoken := ctx.Query("verificationtoken")
	apiKey := ctx.Query("apiKey")
	data, statusCode := service.EmailVerify(apiKey, verificationtoken)

	ctx.JSON(statusCode, data)
}

func Login(ctx *gin.Context) {
	LoginBody := request.LoginRequest{}
	apiKey := ctx.Query("apiKey")
	ctx.ShouldBindJSON(&LoginBody)
	data, statusCode := service.Login(apiKey, &LoginBody)

	ctx.JSON(statusCode, data)
}

func ProfileUpdate(ctx *gin.Context) { //check
	ProfileBody := request.ProfileUpdateRequest{}
	apiKey := ctx.Query("apiKey")
	token := ctx.GetHeader("Authorization")
	ctx.ShouldBindJSON(&ProfileBody)
	data, statusCode := service.ProfileUpdate(apiKey, &ProfileBody, token)

	ctx.JSON(statusCode, data)
}

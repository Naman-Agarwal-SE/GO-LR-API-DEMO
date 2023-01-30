package router

import (
	"golang-demo/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	routes := gin.Default()

	routes.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome to LR")
	})

	account := routes.Group("/account")
	account.GET("/sott", controller.Sott)               //create sott key
	account.POST("/register", controller.Register)      //create user
	account.GET("/emailVerify", controller.EmailVerify) //email varification
	account.POST("/login", controller.Login)            //login
	account.PUT("/", controller.ProfileUpdate)          //update profile

	return routes
}

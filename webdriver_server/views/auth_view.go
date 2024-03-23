package views

import (
	"net/http"

	"webdriver_server/controllers"
	"webdriver_server/db"
	"webdriver_server/service"

	"github.com/gin-gonic/gin"
)

// 用户登陆
func ClientLogin(ctx *gin.Context) {
	jwtService := services.NewJWTService()
	clientAuthService := services.NewUserService(databases.DB)
	controller := controllers.NewClientAuthController(ctx, jwtService, clientAuthService)

	token, err := controller.SignIn()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"token":   token,
	})
}

// 用户注册
func ClientRegister(ctx *gin.Context) {
	jwtService := services.NewJWTService()
	clientAuthService := services.NewUserService(databases.DB)
	controller := controllers.NewClientAuthController(ctx, jwtService, clientAuthService)

	err := controller.SignUp()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "registration success",
	})
}

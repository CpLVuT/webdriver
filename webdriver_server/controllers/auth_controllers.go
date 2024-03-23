package controllers

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"webdriver_server/models"
	"webdriver_server/service"
)

type ClientAuthController struct {
	ctx               *gin.Context
	jwtService        *services.JWTService
	clientAuthService *services.UserService
}

func NewClientAuthController(ctx *gin.Context, jwt *services.JWTService, clientAuthService *services.UserService) *ClientAuthController {
	return &ClientAuthController{
		ctx:               ctx,
		jwtService:        jwt,
		clientAuthService: clientAuthService,
	}
}

type loginForm struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (c *ClientAuthController) SignIn() (string, error) {
	// Get data from request
	var form loginForm
	err := c.ctx.ShouldBindJSON(&form)
	if err != nil {
		return "", err
	}

	// Check if user exists
	user, err := c.clientAuthService.GetUserByEmail(form.Email)
	if err != nil {
		return "", err
	}

	// Check if password is correct
	err = c.clientAuthService.PasswordCheck(user, form.Password)
	if err != nil {
		return "", err
	}

	// Generate token
	return c.jwtService.GenerateToken(form.Email, true), nil
}

type registerForm struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Mobile   string `json:"mobile"`
}

func (c *ClientAuthController) SignUp() error {
	// Get data from request
	var form registerForm
	err := c.ctx.ShouldBindJSON(&form)
	if err != nil {
		return err
	}
	fmt.Println(form.Email)
	fmt.Println(form.Password)
	fmt.Println(form.Nickname)
	fmt.Println(form.Mobile)

	// Check if user exists
	exists, err := c.clientAuthService.ExistanceCheck(form.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user already exists")
	}

	// Create user
	err = c.clientAuthService.CreateUser(&models.User{
		Email:    form.Email,
		Password: form.Password,
		Nickname: form.Nickname,
		Mobile:   &form.Mobile,
	})
	if err != nil {
		return err
	}

	return nil
}

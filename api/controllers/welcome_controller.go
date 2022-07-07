package controllers

import "github.com/gin-gonic/gin"

type WelcomeController interface {
	Login (ctx *gin.Context)
	Register(ctx *gin.Context)
	Logout (ctx *gin.Context)
}

type welcomeController struct{}

func NewWelcomeController() WelcomeController {
	return &welcomeController{}
}

// Healthcheck is to return app condition
func (c *welcomeController) Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok This is the test of the api",
	})
}
func (c *welcomeController) Register(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok This is the test of the api",
	})
}

func (c *welcomeController) Logout(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok This is the test of the api",
	})
}

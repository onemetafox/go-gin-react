package controllers

import "github.com/gin-gonic/gin"

type UserController interface{
	GetUser(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
	SaveUser(ctx *gin.Context)
	DelUser(ctx *gin.Context)
}
type userController struct {}

func NewUserController() UserController{
	return &userController{}
}

func (c *userController) GetUser(ctx *gin.Context){
	
}

func (c *userController) GetAllUsers(ctx *gin.Context){
	
}

func (c *userController) SaveUser(ctx *gin.Context){
	
}

func (c *userController) DelUser(ctx *gin.Context){
	
}
package controllers

import "github.com/gin-gonic/gin"

type ArticleController interface {
	GetAllArticle(ctx *gin.Context)
	GetArticle(ctx *gin.Context)
	SaveArticle(ctx *gin.Context)
	DelArticle(ctx *gin.Context)
}

type articleController struct{}

func NewArticleController() ArticleController {
	return &articleController{}
}

func (c *articleController) GetAllArticle(ctx *gin.Context)  {
	
}

func (c *articleController) GetArticle(ctx *gin.Context)  {
	
}

func (c *articleController) SaveArticle(ctx *gin.Context)  {
	
}

func (c *articleController) DelArticle(ctx *gin.Context)  {
	
}
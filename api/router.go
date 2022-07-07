package api

import (
	"github.com/gin-gonic/gin"
	"github.com/todaycodemaster/go-gin-react/api/controllers"
)

type Router interface {
	SetupRouter(router *gin.Engine)
}

type router struct{}

func NewRouter() Router {
	return &router{}
}

// SetupRouter configuration router information
func (r *router) SetupRouter(router *gin.Engine) {
	var (
		welcomeController controllers.WelcomeController = controllers.NewWelcomeController()
		userController controllers.UserController = controllers.NewUserController()
		articleController controllers.ArticleController = controllers.NewArticleController()
	)

	// Setup route group for the API
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		v1.POST("/login", welcomeController.Login)
		v1.POST("/register", welcomeController.Register)
		v1.GET("/logout", welcomeController.Logout)

		users := v1.Group("/users")
		users.GET("/:id", userController.GetUser)
		users.POST("/",userController.GetAllUsers)
		users.POST("/save", userController.SaveUser)
		users.DELETE("/:id", userController.DelUser)

		articles := v1.Group("/articles")
		articles.POST("/", articleController.GetAllArticle)
		articles.GET("/:id", articleController.GetArticle)
		articles.POST("/save", articleController.SaveArticle)
		articles.DELETE("/:id", articleController.DelArticle)
	}
}

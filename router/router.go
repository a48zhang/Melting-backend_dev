package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"main/handler"
	"main/router/middleware"
	"net/http"
)

func Register(e *gin.Engine) *gin.Engine {
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	e.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API router.")
	})

	e.GET("/api/v1/user", handler.GetOnesInfo)

	v1 := e.Group("/api/v1")
	{

		v1.POST("/login", handler.Login)
		v1.POST("/register", handler.Register)

		v1.Use(middleware.TokenParser)
		v1.GET("/join", handler.JoinProposal)

		users := v1.Group("/users")
		{
			users.GET("", handler.GetUserInfo)
			users.PUT("", handler.UploadProfile)
			users.PUT("/photo", handler.UploadPhoto)
			users.GET("/myproject", handler.Getprojects)
		}
		project := v1.Group("/project")
		{
			project.GET("", handler.GetProject)
			project.PUT("", handler.UpdateProject)
			project.GET("/template", handler.GetTemplate)
			project.POST("/template", handler.CreateTemplate)
			project.PUT("/template", handler.UpdateTemplate)
			project.DELETE("/template", handler.DeleteTemplate)
			project.POST("/newproject", handler.CreateProject)
			project.DELETE("", handler.DeleteProject)

			games := project.Group("/games")
			{
				games.GET("", handler.GameSelect)
				games.POST("/find", handler.FindGames)
				games.GET("/details", handler.GameDetail)
			}
		}
	}

	return e
}

func WSHandlerRegister(e *gin.Engine) *gin.Engine {
	e.GET("", handler.NewWebSocket)
	e.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "This port is websocket-only.")
	})
	return e
}

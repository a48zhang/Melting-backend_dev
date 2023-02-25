package router

import (
	"main/handler"
	"main/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(e *gin.Engine) {
	e.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API router.")
	})

	e.Static("/resource/games", "./resource/games")
	e.GET("/api/v1/avatar", handler.GetOnesAvatar)

	v1 := e.Group("/api/v1")
	{
		v1.POST("/login", handler.Login)
		v1.POST("/register", handler.Register)
		v1.Use(middleware.TokenParser)
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
			project.POST("/newproject", handler.CreateProject)
			games := project.Group("/games")
			{
				games.GET("", handler.GameSelect)
				games.POST("/find", handler.FindGames)
				games.GET("/details", handler.GameDetail)
			}
		}
	}
}

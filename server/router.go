package server

import (
	"github.com/cowboy-coding/go-gin-boilerplate/controllers"
	"github.com/cowboy-coding/go-gin-boilerplate/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Health controller is public accesseble
	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	// add middlewares here. e.g. router.Use(middlewares.AuthMiddleware())
	// the jwt middleware
	jwtAuthMiddleware := middlewares.JwtAuthMiddleware()

	v1 := router.Group("v1")
	{
		// JWT Auth API
		v1.POST("/auth", jwtAuthMiddleware.LoginHandler)
		v1.GET("/auth/refresh_token", jwtAuthMiddleware.RefreshHandler)

		// following routes require AUTH:
		v1.Use(jwtAuthMiddleware.MiddlewareFunc())
		{
			// User API
			userGroup := v1.Group("users")
			{
				users := new(controllers.UsersController)

				// http localhost:8080/v1/users/112
				userGroup.GET("/:id", users.Show)
			}

		}
	}
	return router
}

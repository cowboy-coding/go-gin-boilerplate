package server

import (
	"github.com/cowboy-coding/go-gin-boilerplate/controllers"
	"github.com/cowboy-coding/go-gin-boilerplate/lib/services"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Health controller is public accessible
	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	// add middlewares here. e.g. router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("v1")
	{
		// Routes without Authentication
		jwtAuthMiddleware := services.JwtAuthService()
		v1.POST("/auth", jwtAuthMiddleware.LoginHandler)
		v1.GET("/auth/refresh_token", jwtAuthMiddleware.RefreshHandler)

		// JWT Auth Middleware
		v1.Use(jwtAuthMiddleware.MiddlewareFunc())

		// Routes with Authentication
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

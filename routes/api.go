package routes

import (
	"github.com/allanurbayramgeldiyev209/learngin/controller"
	"github.com/allanurbayramgeldiyev209/learngin/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ApiRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	auth_routers := router.Group("/api/auth")
	{
		auth_routers.POST("/login", controller.Login)
		auth_routers.POST("/register", controller.Register)
		auth_routers.GET("/refresh", controller.RefreshToken)
	}

	routers := router.Group("/api")
	{
		routers.Use(middleware.JwtAuthMiddleware())
		routers.GET("/user", controller.CurrentUser)
	}

	return router
}

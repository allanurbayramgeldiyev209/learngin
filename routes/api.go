package routes

import (
	"github.com/allanurbayramgeldiyev209/learngin/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ApiRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	auth_routers := router.Group("/api/auth")
	{
		auth_routers.GET("/login", controller.Login)
		auth_routers.GET("/register", controller.Register)
	}

	return router
}

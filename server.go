package main

import (
	controller "github.com/allanurbayramgeldiyev209/learngin/controller"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	
	auth_routers := router.Group("/api/auth")
	{
		auth_routers.GET("/login", controller.Login)
		auth_routers.GET("/register", controller.Register)
	}

	router.Run()
}
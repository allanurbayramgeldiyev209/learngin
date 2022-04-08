package controller

import "github.com/gin-gonic/gin"

func Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "login",
	})
}

func Register(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Register",
	})
}

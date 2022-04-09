package controller

import (
	"net/http"

	"github.com/allanurbayramgeldiyev209/learngin/helpers"
	"github.com/allanurbayramgeldiyev209/learngin/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(ctx *gin.Context) {

	var input LoginInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"resp": helpers.BuildErrResponse(err.Error()),
		})
		return
	}

	u := models.User{}

	u.Email = input.Email
	u.Password = input.Password

}

func Register(ctx *gin.Context) {

	var input RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"resp": helpers.BuildErrResponse(err.Error()),
		})
		return
	}

	u := models.User{}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	helpers.CheckErr(err)

	u.Name = input.Name
	u.Email = input.Email
	u.Password = string(hashedPassword)

	u.Add()

	ctx.JSON(http.StatusOK, gin.H{
		"resp": helpers.BuildResponse(u),
	})

}

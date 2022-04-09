package controller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/allanurbayramgeldiyev209/learngin/helpers"
	"github.com/allanurbayramgeldiyev209/learngin/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func Login(ctx *gin.Context) {

	var input LoginInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"resp": helpers.BuildErrResponse(err.Error()),
		})
		return
	}

	err_get_user, user := models.User{}.GetUser(input.Email)

	err_check_password := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	fmt.Println(err_check_password)

	if err_get_user != nil || err_check_password != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"resp": helpers.BuildErrResponse("Parolynyz ya-da email adresiniz yalnys"),
		})
		return
	}

	claims := &Claims{
		UserID: strconv.FormatUint(user.ID, 10),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	errEnv := godotenv.Load(".env")
	helpers.CheckErr(errEnv)

	tokenString, errCreateToken := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if errCreateToken != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"resp": helpers.BuildErrResponse(errCreateToken.Error()),
		})
		return
	}

	models.User{}.UpdateToken(user.Email, tokenString)

	ctx.JSON(http.StatusOK, gin.H{
		"resp": helpers.BuildResponse(tokenString),
	})

}

func Register(ctx *gin.Context) {

	var input RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"resp": helpers.BuildErrResponse(err.Error()),
		})
		return
	}

	err, _ := models.User{}.GetUser(input.Email)
	if err == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"resp": helpers.BuildErrResponse("Beyle ulanyjy on bar. Ony tazeden gosup bolmayar"),
		})
		return
	}

	u := models.User{}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	helpers.CheckErr(err)

	u.Name = input.Name
	u.Email = input.Email
	u.Password = string(hashedPassword)

	u.Add()

	ctx.JSON(http.StatusOK, gin.H{
		"resp": helpers.BuildResponse(u),
	})

}

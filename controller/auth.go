package controller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
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

	claims_for_access := &Claims{
		UserID: strconv.FormatUint(user.ID, 10),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}

	claims_for_refresh := &Claims{
		UserID: strconv.FormatUint(user.ID, 10),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
		},
	}

	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims_for_access)
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims_for_refresh)

	errEnv := godotenv.Load(".env")
	helpers.CheckErr(errEnv)

	access_tokenString, errCreateAccessToken := access_token.SignedString([]byte(os.Getenv("SECRET_KEY_ACCESS_TOKEN")))
	refresh_tokenString, errCreateRefreshToken := refresh_token.SignedString([]byte(os.Getenv("SECRET_KEY_REFRESH_TOKEN")))

	if errCreateAccessToken != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"resp": helpers.BuildErrResponse(errCreateAccessToken.Error()),
		})
		return
	}

	if errCreateRefreshToken != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"resp": helpers.BuildErrResponse(errCreateRefreshToken.Error()),
		})
		return
	}

	models.User{}.UpdateToken(user.Email, access_tokenString, refresh_tokenString)

	ctx.JSON(http.StatusOK, gin.H{
		"resp": helpers.BuildResponse(map[string]string{"access_toke": access_tokenString, "refresh_token": refresh_tokenString}),
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

func TokenValid(ctx *gin.Context) error {

	claims := &Claims{}

	token_string := ExtractToken(ctx)

	tkn, err := jwt.ParseWithClaims(token_string, claims, func(token *jwt.Token) (interface{}, error) {

		errEnv := godotenv.Load(".env")
		helpers.CheckErr(errEnv)

		return []byte(os.Getenv("SECRET_KEY_ACCESS_TOKEN")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"resp": helpers.BuildErrResponse(err.Error()),
			})
			return err
		}
		ctx.JSON(http.StatusForbidden, gin.H{
			"resp": helpers.BuildErrResponse(err.Error()),
		})
		return err
	}

	if !tkn.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"resp": helpers.BuildErrResponse(err.Error()),
		})
		return err
	}
	return nil

}

func ExtractToken(ctx *gin.Context) string {

	token := ctx.Query("token")
	if token != "" {
		return token
	}

	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return bearerToken

}

func ExtractTokenID(ctx *gin.Context) (uint, error) {

	claims := &Claims{}

	token_string := ExtractToken(ctx)

	tkn, err := jwt.ParseWithClaims(token_string, claims, func(token *jwt.Token) (interface{}, error) {

		errEnv := godotenv.Load(".env")
		helpers.CheckErr(errEnv)

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if !tkn.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"resp": helpers.BuildErrResponse(err.Error()),
		})
		return 0, err
	}

	uid, err := strconv.ParseUint(claims.UserID, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(uid), nil

}

func ExtrTokenID(ctx *gin.Context) (uint, error) {

	claims := &Claims{}

	token_string := ExtractToken(ctx)

	tkn, err := jwt.ParseWithClaims(token_string, claims, func(token *jwt.Token) (interface{}, error) {

		errEnv := godotenv.Load(".env")
		helpers.CheckErr(errEnv)

		return []byte(os.Getenv("SECRET_KEY_REFRESH_TOKEN")), nil
	})

	if !tkn.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"resp": helpers.BuildErrResponse(err.Error()),
		})
		return 0, err
	}

	uid, err := strconv.ParseUint(claims.UserID, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(uid), nil

}

func CurrentUser(ctx *gin.Context) {

	user_id, err := ExtractTokenID(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"resp": helpers.BuildErrResponse(err.Error()),
		})
		return
	}

	err, u := models.User{}.GetUserByID(uint(user_id))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"resp": helpers.BuildErrResponse(err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"resp": helpers.BuildResponse(u),
	})

}

func RefreshToken(ctx *gin.Context) {

	user_id, err := ExtrTokenID(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"resp": helpers.BuildErrResponse(err.Error()),
		})
		return
	}

	claims := &Claims{}

	token_string := ExtractToken(ctx)

	tkn, err := jwt.ParseWithClaims(token_string, claims, func(token *jwt.Token) (interface{}, error) {

		errEnv := godotenv.Load(".env")
		helpers.CheckErr(errEnv)

		return []byte(os.Getenv("SECRET_KEY_REFRESH_TOKEN")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"resp": helpers.BuildErrResponse(err.Error()),
			})
			return
		}
		ctx.JSON(http.StatusForbidden, gin.H{
			"resp": helpers.BuildErrResponse(err.Error()),
		})
		return
	}

	if !tkn.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"resp": helpers.BuildErrResponse(err.Error()),
		})
		return
	}

	claims_for_access := &Claims{
		UserID: strconv.FormatUint(uint64(user_id), 10),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}

	claims_for_refresh := &Claims{
		UserID: strconv.FormatUint(uint64(user_id), 10),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
		},
	}

	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims_for_access)
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims_for_refresh)

	errEnv := godotenv.Load(".env")
	helpers.CheckErr(errEnv)

	access_tokenString, errCreateAccessToken := access_token.SignedString([]byte(os.Getenv("SECRET_KEY_ACCESS_TOKEN")))
	refresh_tokenString, errCreateRefreshToken := refresh_token.SignedString([]byte(os.Getenv("SECRET_KEY_REFRESH_TOKEN")))

	if errCreateAccessToken != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"resp": helpers.BuildErrResponse(errCreateAccessToken.Error()),
		})
		return
	}

	if errCreateRefreshToken != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"resp": helpers.BuildErrResponse(errCreateRefreshToken.Error()),
		})
		return
	}

	models.User{}.UpdateTokenByID(user_id, access_tokenString, refresh_tokenString)

	ctx.JSON(http.StatusOK, gin.H{
		"resp": helpers.BuildResponse(map[string]string{"access_toke": access_tokenString, "refresh_token": refresh_tokenString}),
	})

}

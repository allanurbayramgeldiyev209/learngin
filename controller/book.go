package controller

import (
	"net/http"
	"strconv"

	"github.com/allanurbayramgeldiyev209/learngin/helpers"
	"github.com/allanurbayramgeldiyev209/learngin/models"
	"github.com/gin-gonic/gin"
)

type AddBookInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateBookInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func AddBook(ctx *gin.Context) {

	var input AddBookInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"resp": helpers.BuildErrResponse(err.Error()),
		})
		return
	}

	models.Book{
		Title:       input.Title,
		Description: input.Description,
	}.Add()

	ctx.JSON(http.StatusOK, gin.H{
		"resp": helpers.BuildResponse("Kitap ustunlikli gosuldy"),
	})

}

func GetBooks(ctx *gin.Context) {

	books := models.Book{}.GetAll()

	if books[0].ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"resp": helpers.BuildErrResponse("Kitap yok"),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"resp": helpers.BuildResponse(books),
	})

}

func GetBook(ctx *gin.Context) {

	id_string := ctx.Param("id")
	id, err := strconv.Atoi(id_string)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"resp": helpers.BuildErrResponse("Beyle maglumat yok"),
		})
		return
	}

	book := models.Book{}.Get(uint(id))

	if book.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"resp": helpers.BuildErrResponse("Beyle maglumat yok"),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"resp": helpers.BuildResponse(book),
	})

}

func UpdateBook(ctx *gin.Context) {

	var input UpdateBookInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"resp": helpers.BuildErrResponse(err.Error()),
		})
		return
	}

	id_string := ctx.Param("id")
	id, err := strconv.Atoi(id_string)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"resp": helpers.BuildErrResponse("Beyle maglumat yok"),
		})
		return
	}

	book := models.Book{}.Get(uint(id))

	if book.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"resp": helpers.BuildErrResponse("Beyle maglumat yok"),
		})
		return
	}

	models.Book{}.Update(uint(id), input.Title, input.Description)

	ctx.JSON(http.StatusOK, gin.H{
		"resp": helpers.BuildResponse("Maglumat ustunlikli uytgedildi"),
	})

}

func DeleteBook(ctx *gin.Context) {

	id_string := ctx.Param("id")
	id, err := strconv.Atoi(id_string)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"resp": helpers.BuildErrResponse("Beyle maglumat yok"),
		})
		return
	}

	book := models.Book{}.Get(uint(id))

	if book.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"resp": helpers.BuildErrResponse("Beyle maglumat yok"),
		})
		return
	}

	models.Book{}.Delete(uint(id))

	ctx.JSON(http.StatusOK, gin.H{
		"resp": helpers.BuildResponse("Maglumat ustunlikli pozuldy"),
	})

}

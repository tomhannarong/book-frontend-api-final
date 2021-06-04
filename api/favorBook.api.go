package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/interceptor"
	"book-frontend-api-final/model"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// SetupFavorBookAPI - Favorite Book
func SetupFavorBookAPI(router *gin.Engine) {
	favorBookAPI := router.Group("/api")
	{
		// Get all favorite products owner list
		favorBookAPI.GET("/favorBook", interceptor.JwtVerify, getFavorBook)

		// Get favorite products  by Id
		favorBookAPI.GET("/favorBook/:id", interceptor.JwtVerify, getFavorBookByID)

		// Create or Add , favorite product
		favorBookAPI.POST("/favorBook", interceptor.JwtVerify, createFavorBook)

		// Delete favorite product by Id
		favorBookAPI.DELETE("/favorBook/:id", interceptor.JwtVerify, deleteFavorBook)

		// Delete product by Id from favorite book
		// favorBookAPI.DELETE("/favorBook/product/:id", interceptor.JwtVerify, deleteProductFavor)
	}
}

func getFavorBook(c *gin.Context) {
	var favorBook []model.FavorBook

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%s", keyword)
		db.GetDB().Where("book_id = ?", keyword).Find(&favorBook)
	} else {
		db.GetDB().Find(&favorBook).Where("username = ?", c.GetString("jwt_username")).Order("created_at desc")
	}
	c.JSON(http.StatusOK, favorBook)

}

func getFavorBookByID(c *gin.Context) {
	var favorBook model.FavorBook
	db.GetDB().Where("id = ?", c.Param("id")).First(&favorBook)
	c.JSON(http.StatusOK, favorBook)
}

func createFavorBook(c *gin.Context) {
	BookID, _ := strconv.ParseInt(c.PostForm("book_id"), 10, 64)
	UserID, _ := strconv.ParseInt(c.PostForm("user_id"), 10, 64)

	favorBook := model.FavorBook{}
	favorBook.BookID = BookID
	favorBook.UserID = UserID
	favorBook.IsEbook = c.PostForm("is_ebook")
	favorBook.CreatedAt = time.Now()
	db.GetDB().Create(&favorBook)

	c.JSON(http.StatusOK, gin.H{"result": favorBook})

}

func deleteFavorBook(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	db.GetDB().Delete(&model.FavorBook{}, id)
	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}

func deleteProductFavor(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	db.GetDB().Where("book_id = ?", id).Delete(&model.FavorBook{})
	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}

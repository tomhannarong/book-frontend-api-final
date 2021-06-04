package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupProductBookAPI - pProducts Book
func SetupProductBookAPI(router *gin.Engine) {
	productBookAPI := router.Group("/api/")
	{
		// Get all Products book own list [ book ]
		productBookAPI.GET("/product-book" /*interceptor.JwtVerify,*/, getProductBook)

		// Get product book by Id
		productBookAPI.GET("/product-book/:id" /*interceptor.JwtVerify,*/, getProductBookByID)
	}
}

func getProductBook(c *gin.Context) {
	var product []model.Product

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("book_name like ?", keyword).Where("is_ebook = 0").Find(&product)
	} else {
		db.GetDB().Where("is_ebook = 0").Find(&product)
	}
	c.JSON(http.StatusOK, product)

}

func getProductBookByID(c *gin.Context) {
	var product model.Product
	db.GetDB().Where("id = ?", c.Param("id")).Where("is_ebook = 0").First(&product)
	c.JSON(http.StatusOK, product)
}

package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupProductEbookAPI - Products Ebook
func SetupProductEbookAPI(router *gin.Engine) {
	productEbookAPI := router.Group("/api/")
	{
		// Get all Products e-book own list [ E-book ]
		productEbookAPI.GET("/product-ebook" /*interceptor.JwtVerify,*/, getProductEBook)

		// Get product e-book by Id
		productEbookAPI.GET("/product-ebook/:id" /*interceptor.JwtVerify,*/, getProductEBookByID)
	}
}

func getProductEBook(c *gin.Context) {
	var product []model.Product

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("book_name like ?", keyword).Where("is_ebook = 1").Find(&product)
	} else {
		db.GetDB().Where("is_ebook = 1").Find(&product)
	}
	c.JSON(http.StatusOK, product)

}

func getProductEBookByID(c *gin.Context) {
	var product model.Product
	db.GetDB().Where("id = ?", c.Param("id")).Where("is_ebook = 1").First(&product)
	c.JSON(http.StatusOK, product)
}

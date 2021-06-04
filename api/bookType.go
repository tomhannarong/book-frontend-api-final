package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupBookTypeAPI - Book Type
func SetupBookTypeAPI(router *gin.Engine) {
	bookTypeAPI := router.Group("/api")
	{
		// Get book type list .
		bookTypeAPI.GET("/bookType" /*interceptor.JwtVerify,*/, getBookType)
	}
}

func getBookType(c *gin.Context) {
	var bookType []model.BookType

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("book_type like ?", keyword).Find(&bookType)
	} else {
		db.GetDB().Find(&bookType).Order("created_at desc")
	}
	c.JSON(http.StatusOK, bookType)

}

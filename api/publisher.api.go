package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupPublisherAPI - Publisher owner
func SetupPublisherAPI(router *gin.Engine) {
	publisherAPI := router.Group("/api")
	{
		// Get all Publisher own list
		publisherAPI.GET("/publisher" /*interceptor.JwtVerify,*/, getPublisher)

		// Get Publisher by Id
		publisherAPI.GET("/publisher/:id" /*interceptor.JwtVerify,*/, getPublisherByID)
	}
}

func getPublisher(c *gin.Context) {
	var publisher []model.Publisher

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("publisher like ? ", keyword).Find(&publisher)
	} else {
		db.GetDB().Find(&publisher).Order("created_at desc")
	}
	c.JSON(http.StatusOK, publisher)

}

func getPublisherByID(c *gin.Context) {
	publisher := model.Publisher{}
	db.GetDB().Where("id = ?", c.Param("id")).First(&publisher)
	c.JSON(http.StatusOK, publisher)
}

package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupTransportAPI - delivery
func SetupTransportAPI(router *gin.Engine) {
	transportAPI := router.Group("/api")
	{
		// Get delivery own list
		transportAPI.GET("/transport" /*interceptor.JwtVerify,*/, getTransport)

		// Get delivery by Id
		transportAPI.GET("/transport/:id" /*interceptor.JwtVerify,*/, getTransportByID)
	}
}

func getTransport(c *gin.Context) {
	var transport []model.Transport

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("transport	like ? ", keyword).Find(&transport)
	} else {
		db.GetDB().Find(&transport).Order("created_at desc")
	}
	c.JSON(http.StatusOK, transport)

}

func getTransportByID(c *gin.Context) {
	transport := model.Transport{}
	db.GetDB().Where("id = ?", c.Param("id")).First(&transport)
	c.JSON(http.StatusOK, transport)
}

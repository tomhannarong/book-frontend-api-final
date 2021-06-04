package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupAmphureAPI - Amphure list
func SetupAmphureAPI(router *gin.Engine) {
	amphureAPI := router.Group("/api")
	{
		// Get amphure list .
		amphureAPI.GET("/amphure" /*interceptor.JwtVerify,*/, getAmphure)

		// Get amphure by Id .
		amphureAPI.GET("/amphure/:id" /*interceptor.JwtVerify,*/, getAmphureByID)
	}
}

func getAmphure(c *gin.Context) {
	var amphure []model.Amphure

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%s", keyword)
		db.GetDB().Where("name_th = ? OR name_en = ? ", keyword, keyword).Find(&amphure)
	} else {
		db.GetDB().Find(&amphure).Order("created_at desc")
	}
	c.JSON(http.StatusOK, amphure)

}

func getAmphureByID(c *gin.Context) {
	amphure := model.Amphure{}
	db.GetDB().Where("id = ?", c.Param("id")).First(&amphure)
	c.JSON(http.StatusOK, amphure)
}

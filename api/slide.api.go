package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupSlideAPI - Silde home page
func SetupSlideAPI(router *gin.Engine) {
	slideAPI := router.Group("/api")
	{
		// Get slides in home page
		slideAPI.GET("/slide" /*interceptor.JwtVerify,*/, getSlide)
	}
}

func getSlide(c *gin.Context) {
	var slide []model.Slide

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("slide_name like ? ", keyword).Find(&slide)
	} else {
		db.GetDB().Find(&slide).Order("created_at desc")
	}
	c.JSON(http.StatusOK, slide)

}

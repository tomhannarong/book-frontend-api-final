package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupPrivacyAPI -  Privacy News
func SetupPrivacyAPI(router *gin.Engine) {
	privacyAPI := router.Group("/api")
	{
		// Get privacy blog news list , Post and Update by admin only
		privacyAPI.GET("/privacy" /*interceptor.JwtVerify,*/, getPrivacy)

		// Get privacy by Id
		privacyAPI.GET("/privacy/:id" /*interceptor.JwtVerify,*/, getPrivacyByID)
	}
}

func getPrivacy(c *gin.Context) {
	var privacy []model.Privacy

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("title like ? ", keyword).Find(&privacy)
	} else {
		db.GetDB().Find(&privacy).Order("created_at desc")
	}
	c.JSON(http.StatusOK, privacy)

}

func getPrivacyByID(c *gin.Context) {
	privacy := model.Privacy{}
	db.GetDB().Where("id = ?", c.Param("id")).First(&privacy)
	c.JSON(http.StatusOK, privacy)
}

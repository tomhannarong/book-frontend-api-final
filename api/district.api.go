package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupDistrictAPI - District list
func SetupDistrictAPI(router *gin.Engine) {
	districtAPI := router.Group("/api")
	{
		// Get all district list .
		districtAPI.GET("/district" /*interceptor.JwtVerify,*/, getDistrict)

		// Get district by Id .
		districtAPI.GET("/district/:id" /*interceptor.JwtVerify,*/, getDistrictByID)
	}
}

func getDistrict(c *gin.Context) {
	var district []model.District

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("name_th like ? OR name_en like ? ", keyword, keyword).Find(&district)
	} else {
		db.GetDB().Find(&district).Order("created_at desc")
	}
	c.JSON(http.StatusOK, district)

}

func getDistrictByID(c *gin.Context) {
	district := model.District{}
	db.GetDB().Where("id = ?", c.Param("id")).First(&district)
	c.JSON(http.StatusOK, district)
}

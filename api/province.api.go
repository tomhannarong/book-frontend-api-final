package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupProvinceAPI - Province list
func SetupProvinceAPI(router *gin.Engine) {
	provinceAPI := router.Group("/api")
	{
		// Get province list
		provinceAPI.GET("/province" /*interceptor.JwtVerify,*/, getProvince)

		// Get province by Id
		provinceAPI.GET("/province/:id" /*interceptor.JwtVerify,*/, getProvinceByID)
	}
}

func getProvince(c *gin.Context) {
	var province []model.Province

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("name_th like ? OR name_en like ? ", keyword, keyword).Find(&province)
	} else {
		db.GetDB().Find(&province).Order("created_at desc")
	}
	c.JSON(http.StatusOK, province)

}

func getProvinceByID(c *gin.Context) {
	province := model.Province{}
	db.GetDB().Where("id = ?", c.Param("id")).First(&province)
	c.JSON(http.StatusOK, province)
}

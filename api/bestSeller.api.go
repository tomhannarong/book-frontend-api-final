package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupBestSellerAPI - Best Seller Book
func SetupBestSellerAPI(router *gin.Engine) {
	bestSellerAPI := router.Group("/api")
	{
		// Get Best Seller type book only
		bestSellerAPI.GET("/bestSeller" /*interceptor.JwtVerify,*/, getBestSeller)

		// Get Best Seller type book only By Id .
		bestSellerAPI.GET("/bestSeller/:id" /*interceptor.JwtVerify,*/, getBestSellerByID)
	}
}

func getBestSeller(c *gin.Context) {
	var bestSeller []model.BestSeller

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%s", keyword)
		db.GetDB().Where("top = ?", keyword).Find(&bestSeller)
	} else {
		db.GetDB().Find(&bestSeller).Order("created_at desc")
	}
	c.JSON(http.StatusOK, bestSeller)

}

func getBestSellerByID(c *gin.Context) {
	bestSeller := model.BestSeller{}
	db.GetDB().Where("id = ?", c.Param("id")).First(&bestSeller)
	c.JSON(http.StatusOK, bestSeller)
}

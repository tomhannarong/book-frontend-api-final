package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupEbookBestsellerAPI -  E-book Bestseller list
func SetupEbookBestsellerAPI(router *gin.Engine) {
	ebookBestsellerAPI := router.Group("/api")
	{
		// Get Product Best Seller list .
		ebookBestsellerAPI.GET("/ebookBestseller" /*interceptor.JwtVerify,*/, getEbookBestseller)

		// Get product Best Seller list by Id .
		ebookBestsellerAPI.GET("/ebookBestseller/:id" /*interceptor.JwtVerify,*/, getEbookBestsellerByID)
	}
}

func getEbookBestseller(c *gin.Context) {
	var ebookBestseller []model.EbookBestseller

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%s", keyword)
		db.GetDB().Where("product_id = ? ", keyword).Group("PostDate").Find(&ebookBestseller)
	} else {
		db.GetDB().Select("*, sum(seq) as sumSeq").Find(&ebookBestseller).Group("PostDate").Order("created_at desc")
	}
	c.JSON(http.StatusOK, ebookBestseller)

}

func getEbookBestsellerByID(c *gin.Context) {
	ebookBestseller := model.EbookBestseller{}
	db.GetDB().Where("id = ?", c.Param("id")).First(&ebookBestseller)
	c.JSON(http.StatusOK, ebookBestseller)
}

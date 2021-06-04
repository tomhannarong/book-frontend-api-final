package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/interceptor"
	"book-frontend-api-final/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupEbookApproveEbookAPI - Approve E-book read
func SetupEbookApproveEbookAPI(router *gin.Engine) {
	ebookApproveEbookAPI := router.Group("/api")
	{
		// Get e-book approveed , can have read  owner list .
		ebookApproveEbookAPI.GET("/approveEbook", interceptor.JwtVerify, getBookApproveEbook)

		// Get products e-book approved by Id
		ebookApproveEbookAPI.GET("/approveEbook/:id", interceptor.JwtVerify, getBookApproveEbookByID)
	}
}

func getBookApproveEbook(c *gin.Context) {
	var ebookApproveEbook []model.EbookApproveEbook

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%s", keyword)
		db.GetDB().Where("product_id = ? ", keyword).Where("username = ?", c.GetString("jwt_username")).Find(&ebookApproveEbook)
	} else {
		db.GetDB().Find(&ebookApproveEbook).Where("username = ?", c.GetString("jwt_username")).Order("created_at desc")
	}
	c.JSON(http.StatusOK, ebookApproveEbook)

}

func getBookApproveEbookByID(c *gin.Context) {
	ebookApproveEbook := model.EbookApproveEbook{}
	db.GetDB().Where("id = ?", c.Param("id")).First(&ebookApproveEbook)
	c.JSON(http.StatusOK, ebookApproveEbook)
}

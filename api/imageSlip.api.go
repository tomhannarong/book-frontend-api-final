package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/interceptor"
	"book-frontend-api-final/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// SetupImageSlipAPI - Image Slips
func SetupImageSlipAPI(router *gin.Engine) {
	imageSlipAPI := router.Group("/api")
	{
		// Create Image silp from bank save by order id
		imageSlipAPI.POST("/imageSlip", interceptor.JwtVerify, createImageSlip)
	}
}

func createImageSlip(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("order_id"), 10, 32)
	imageSlip := model.ImageSlip{}
	imageSlip.Filename = c.PostForm("filename")
	imageSlip.OrderID = id
	imageSlip.CreatedAt = time.Now()
	db.GetDB().Create(&imageSlip)

	c.JSON(http.StatusOK, gin.H{"result": imageSlip})

}

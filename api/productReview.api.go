package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/interceptor"
	"book-frontend-api-final/model"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// SetupProductReviewAPI - Product Review from user paid
func SetupProductReviewAPI(router *gin.Engine) {
	productReviewAPI := router.Group("/api")
	{
		// Get all product review , comment product .
		productReviewAPI.GET("/product/:id/product-review" /*interceptor.JwtVerify,*/, getProductReview)

		// Create product review , comment product .
		productReviewAPI.POST("/product/:id/product-review", interceptor.JwtVerify, createProductReview)

	}
}

func getProductReview(c *gin.Context) {
	var productReview []model.ProductReview

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("Message like ?", keyword).Where("username = ?", c.GetString("jwt_username")).Find(&productReview)
	} else {
		db.GetDB().Where("username = ?", c.GetString("jwt_username")).Find(&productReview)
	}
	c.JSON(http.StatusOK, productReview)

}

func createProductReview(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("product_id"), 10, 64)

	// Check user was paid product .
	// Have product
	var queryOrderHistory model.OrderHistory
	if err := db.GetDB().Where("username = ? AND product_id = ?", c.GetString("jwt_username"), id).First(&queryOrderHistory).Error; err == nil {

		// Create rate product , star point
		productReview := model.ProductReview{}
		productReview.ProductID = id
		productReview.Username = c.GetString("jwt_username")
		productReview.Message = c.PostForm("message")
		productReview.CreatedAt = time.Now()
		db.GetDB().Create(&productReview)

		c.JSON(http.StatusOK, gin.H{"result": productReview})
	} else {
		// Not was bought product .
		c.JSON(http.StatusBadRequest, 0)
	}

}

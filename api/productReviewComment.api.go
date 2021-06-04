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

// SetupProductReviewCommentAPI - Comment review product
func SetupProductReviewCommentAPI(router *gin.Engine) {
	productReviewCommentAPI := router.Group("/api")
	{
		// Get product -> review -> all comment .
		productReviewCommentAPI.GET("/product/:id/product-review/:id-review/product-Comment" /*interceptor.JwtVerify,*/, getProductReviewComment)

		// Create review comment .
		productReviewCommentAPI.POST("/product/:id/product-review/:id-review/product-Comment", interceptor.JwtVerify, createProductReviewComment)

	}
}

func getProductReviewComment(c *gin.Context) {
	var productReviewComment []model.ProductReviewComment

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("Message like ?", keyword).Where("product_id = ? AND product_review_comment_id = ?", c.Param("id"), c.Param("id-review")).Find(&productReviewComment)
	} else {
		db.GetDB().Where("product_id = ? AND product_review_comment_id = ?", c.Param("id"), c.Param("id-review")).Find(&productReviewComment)
	}
	c.JSON(http.StatusOK, productReviewComment)

}

func createProductReviewComment(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("product_id"), 10, 64)

	// Check user was paid product .
	// Have product
	var queryOrderHistory model.OrderHistory
	if err := db.GetDB().Where("username = ? AND product_id = ?", c.GetString("jwt_username"), id).First(&queryOrderHistory).Error; err == nil {

		// Create rate product , star point
		productReviewComment := model.ProductReviewComment{}
		productReviewComment.ProductID = id
		productReviewComment.Username = c.GetString("jwt_username")
		productReviewComment.Message = c.PostForm("message")
		productReviewComment.CreatedAt = time.Now()
		db.GetDB().Create(&productReviewComment)

		c.JSON(http.StatusOK, gin.H{"result": productReviewComment})
	} else {
		// Not was bought product .
		c.JSON(http.StatusBadRequest, 0)
	}

}

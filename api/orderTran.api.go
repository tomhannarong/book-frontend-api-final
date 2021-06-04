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

// SetupOrderTranAPI - OrderTran Or product in Order master
func SetupOrderTranAPI(router *gin.Engine) {
	orderTranAPI := router.Group("/api")
	{
		// Get Products item in order master by order id
		orderTranAPI.GET("/order/:id/orderTran", interceptor.JwtVerify, getOrderTran)

		// Create or add product in order master by order id
		orderTranAPI.POST("/order/:id/orderTran", interceptor.JwtVerify, createOrderTran)
	}
}

func getOrderTran(c *gin.Context) {
	var orderTran []model.OrderTran

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%s", keyword)
		db.GetDB().Where("product_name = ?", keyword).Where("username = ?", c.GetString("jwt_username")).Find(&orderTran)
	} else {
		db.GetDB().Where("username = ?", c.GetString("jwt_username")).Order("created_at desc").Find(&orderTran)
	}
	c.JSON(http.StatusOK, orderTran)

}

func createOrderTran(c *gin.Context) {
	// UserID, _ := strconv.ParseInt(c.GetString("jwt_staff_id"), 10, 64)
	BookID, _ := strconv.ParseInt(c.PostForm("book_id"), 10, 64)

	orderTran := model.OrderTran{}
	orderTran.BookID = BookID
	orderTran.Username = c.PostForm("username")
	orderTran.CreatedAt = time.Now()
	db.GetDB().Create(&orderTran)

	c.JSON(http.StatusOK, gin.H{"result": orderTran})

}

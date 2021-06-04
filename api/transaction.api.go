package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/interceptor"
	"book-frontend-api-final/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// SetupTransactionAPI - Transaction order
func SetupTransactionAPI(router *gin.Engine) {
	transactionAPI := router.Group("/api/")
	{
		// Get transaction buy list
		transactionAPI.GET("/transaction", getTransaction)

		// Create transaction buy when buy products
		transactionAPI.POST("/transaction", interceptor.JwtVerify, createTransaction)
	}
}

func getTransaction(c *gin.Context) {
	var transactions []model.Transaction
	db.GetDB().Find(&transactions)
	c.JSON(http.StatusOK, transactions)
}

func createTransaction(c *gin.Context) {
	var transaction model.Transaction
	if err := c.ShouldBind(&transaction); err == nil {
		transaction.StaffID = c.GetString("jwt_staff_id") //get value JSON from JWT token
		transaction.CreatedAt = time.Now()
		db.GetDB().Create(&transaction)
		c.JSON(http.StatusOK, gin.H{"result": "ok", "data": transaction})
	} else {
		c.JSON(404, gin.H{"result": "nok"})
	}
}

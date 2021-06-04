package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupPaymentAPI - Payment list
func SetupPaymentAPI(router *gin.Engine) {
	paymentAPI := router.Group("/api")
	{
		// Get payment bank list ,
		paymentAPI.GET("/payment" /*interceptor.JwtVerify,*/, getPayment)

		// Get payment by Id
		paymentAPI.GET("/payment/:id" /*interceptor.JwtVerify,*/, getPaymentByID)
	}
}

func getPayment(c *gin.Context) {
	var payment []model.Payment

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("payment like ? ", keyword).Find(&payment)
	} else {
		db.GetDB().Find(&payment).Order("created_at desc")
	}
	c.JSON(http.StatusOK, payment)

}

func getPaymentByID(c *gin.Context) {
	payment := model.Payment{}
	db.GetDB().Where("id = ?", c.Param("id")).First(&payment)
	c.JSON(http.StatusOK, payment)
}

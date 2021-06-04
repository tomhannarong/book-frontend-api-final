package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// SetupContactUsAPI - Contact Us
func SetupContactUsAPI(router *gin.Engine) {
	contactUsAPI := router.Group("/api")
	{
		// Create Contact us , using by user customer Or unknow user ,
		contactUsAPI.POST("/contactUs" /*interceptor.JwtVerify,*/, createContactUs)
	}
}

func createContactUs(c *gin.Context) {
	contactUs := model.ContactUs{}
	contactUs.Topic = c.PostForm("topic")
	contactUs.Name = c.PostForm("name")
	contactUs.Tel = c.PostForm("tel")
	contactUs.Email = c.PostForm("email")
	contactUs.Message = c.PostForm("message")
	contactUs.CreatedAt = time.Now()
	db.GetDB().Create(&contactUs)

	c.JSON(http.StatusOK, gin.H{"result": contactUs})

}

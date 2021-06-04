package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/interceptor"
	"book-frontend-api-final/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// SetupAuthenAPI - login, register, reset-password, refresh-token
func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/")
	{

		// Login , Get token
		authenAPI.POST("/login", login)

		// Register , get user customer .
		authenAPI.POST("/register", register)

		// Forget password
		// Requset reset password sendgrid to send email.  ,
		// Input E-mail only , check Have email
		authenAPI.POST("/requset-reset-password", requsetResetPassword)

		// Reset password
		// Set password new , and confirm password .
		authenAPI.POST("/reset-password", resetPassword)

		// Change refresh token , older to new token
		// send refresh token to body requset  and return , {}
		authenAPI.POST("/refresh-token", refreshToken)

	}
}

func login(c *gin.Context) {
	var user model.User

	if c.ShouldBind(&user) == nil {
		var queryUser model.User
		if err := db.GetDB().First(&queryUser, "username = ?", user.Username).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"result": "nok", "error": err})
		} else if checkPasswordHash(user.Password, queryUser.Password) == false {
			c.JSON(http.StatusOK, gin.H{"result": "nok", "error": "invalid password"})
		} else {
			token := interceptor.JwtSign(queryUser)

			c.JSON(http.StatusOK, gin.H{"result": "ok", "token": token})
		}

	} else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}
}

func register(c *gin.Context) {
	var user model.User
	if c.ShouldBind(&user) == nil {
		user.Password, _ = hashPassword(user.Password)
		user.CreatedAt = time.Now()
		if err := db.GetDB().Create(&user).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"result": "nok", "error": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "ok", "data": user})
		}
	} else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func requsetResetPassword(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"result": 0})

}

func resetPassword(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"result": 0})

}

func refreshToken(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"result": 0})

}

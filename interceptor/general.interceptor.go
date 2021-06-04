package interceptor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GeneralInterceptor - call this methos to add interceptor
func GeneralInterceptor(c *gin.Context) {
	token := c.Query("token")
	if token == "5be28cd735d6174a5b4efdc29799565c41ac8fe2" {
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
	}
}

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

// SetupBoardReplyAPI - WebBoard Reply
func SetupBoardReplyAPI(router *gin.Engine) {
	boardReplyAPI := router.Group("/api/boardPost")
	{
		// Create Board Reply , Or comments Post from at WebBoard .
		boardReplyAPI.POST("/:idPost/boardReply", interceptor.JwtVerify, createBoardReply)

		// Get Board Reply by Id .
		boardReplyAPI.GET("/:id/boardReply/:idReply", interceptor.JwtVerify, getboardReplyByID)

		// Update and Edit content in comment .
		boardReplyAPI.PUT("/boardPost/:idPost/boardReply/:idReply/update", interceptor.JwtVerify, getboardReplyUpdate)

		// Delete Comment by Id
		boardReplyAPI.DELETE("/boardPost/:idPost/boardReply/:idReply/delete", interceptor.JwtVerify, getboardReplyDelete)
	}
}

func createBoardReply(c *gin.Context) {
	idPost, _ := strconv.ParseInt(c.Param("idPost"), 10, 32)

	boardReply := model.BoardReply{}
	boardReply.TopicID = idPost
	boardReply.Username = c.PostForm("username")
	boardReply.Reply = c.PostForm("reply")
	boardReply.CreatedAt = time.Now()
	db.GetDB().Create(&boardReply)

	c.JSON(http.StatusOK, gin.H{"result": boardReply})

}

func getboardReplyByID(c *gin.Context) {
	boardReply := model.BoardReply{}
	db.GetDB().Where("id = ?", c.Param("idReply")).First(&boardReply)
	c.JSON(http.StatusOK, boardReply)
}

func getboardReplyUpdate(c *gin.Context) {

	c.JSON(http.StatusOK, 0)
}

func getboardReplyDelete(c *gin.Context) {
	c.JSON(http.StatusOK, 0)
}

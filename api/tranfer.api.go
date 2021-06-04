package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/interceptor"
	"book-frontend-api-final/model"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const vOrderTypeWaitPay string = "รอการชำระเงิน"
const vOrderTypeChecking string = "กำลังตรวจสอบ"
const vOrderTypePaid string = "ชำระเงินแล้ว"
const vOrderTypeSendBack string = "ส่งกลับแก้ไข"
const vOrderTypeSended string = "ส่งสินค้าแล้ว"
const vOrderTypeCancel string = "ยกเลิก"

// SetupTranferAPI - setup Tranfer is send silp
func SetupTranferAPI(router *gin.Engine) {
	tranferAPI := router.Group("/api")
	{
		// Create tranfer when order paid and alrady send slips banking to admin review
		tranferAPI.POST("/order/:id/tranfer", interceptor.JwtVerify, createTranfer)
	}
}

func createTranfer(c *gin.Context) {

	layOut := "2006-01-02 15:04"
	dateStamp, _ := time.Parse(layOut, c.PostForm("tranfer_date"))
	OrderID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	// Create Tranfer
	tranfer := model.Tranfer{}
	tranfer.Amount, _ = strconv.ParseFloat(c.PostForm("amount"), 2)
	tranfer.AccountTranfer = c.PostForm("account_tranfer")
	tranfer.BankTranfer = c.PostForm("bank_tranfer")
	tranfer.Username = c.GetString("jwt_username")
	tranfer.Remark1 = c.PostForm("remark1")
	tranfer.TranferStatus = vOrderTypeChecking
	tranfer.ShowStatus = "y"
	tranfer.TranferDate = dateStamp
	tranfer.CreatedAt = time.Now()
	go db.GetDB().Create(&tranfer)

	// Create slips
	imageSlip := model.ImageSlip{}
	imageSlip.OrderID = OrderID
	imageSlip.CreatedAt = time.Now()
	image, _ := c.FormFile("image")
	go saveImageSlip(image, &imageSlip, c)

	c.JSON(http.StatusOK, gin.H{"result": tranfer})

}

// TempFileName generates a temporary filename for use in testing or whatever
func TempFileName(prefix, suffix string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return filepath.Join(os.TempDir(), prefix+hex.EncodeToString(randBytes)+suffix)
}

// Save Image slips to Local
func saveImageSlip(image *multipart.FileHeader, imageSlip *model.ImageSlip, c *gin.Context) {
	if image != nil {
		runningDir, _ := os.Getwd()
		extension := filepath.Ext(image.Filename)
		// Filename
		filenameLocal := fmt.Sprintf("%d%s", uuid.Must(uuid.NewRandom()), extension)

		// Save Filename in database
		imageSlip.Filename = filenameLocal
		filePath := fmt.Sprintf("%s/uploaded/images/%s", runningDir, filenameLocal)

		if fileExists(filePath) {
			os.Remove(filePath)
		}
		c.SaveUploadedFile(image, filePath)
		db.GetDB().Model(&imageSlip).Update("image", filenameLocal)
	}
}

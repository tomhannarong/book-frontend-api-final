package api

import (
	"book-frontend-api-final/db"

	"github.com/gin-gonic/gin"
)

// Setup - API
func Setup(router *gin.Engine) {

	db.SetupDB()
	SetupAuthenAPI(router)
	SetupProductAPI(router)
	SetupTransactionAPI(router)
	SetupAddressAPI(router)
	SetupBestSellerAPI(router)
	SetupAmphureAPI(router)
	SetupBoardPostAPI(router)
	SetupBoardReplyAPI(router)
	SetupBookTypeAPI(router)
	SetupContactUsAPI(router)
	SetupDistrictAPI(router)
	SetupEbookApproveEbookAPI(router)
	SetupEbookBestsellerAPI(router)
	SetupFavorBookAPI(router)
	SetupImageSlipAPI(router)
	SetupOrderHistoryAPI(router)
	SetupOrderMasAPI(router)
	SetupOrderTranAPI(router)
	SetupPaymentAPI(router)
	SetupPrivacyAPI(router)
	SetupProvinceAPI(router)
	SetupPublisherAPI(router)
	SetupTransportAPI(router)
	SetupSlideAPI(router)
	SetupProductBookAPI(router)
	SetupProductEbookAPI(router)
	SetupProductGalleryPhotoAPI(router)
	SetupProductRateAPI(router)
	SetupProductReviewAPI(router)
	SetupProductReviewCommentAPI(router)
	SetupTranferAPI(router)
}

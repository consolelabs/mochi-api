package wallet

import "github.com/gin-gonic/gin"

type IHandler interface {
	ListOwnedWallets(c *gin.Context)
	ListTrackingWallets(c *gin.Context)
	GetOne(c *gin.Context)
	Track(c *gin.Context)
	Untrack(c *gin.Context)
	ListAssets(c *gin.Context)
	ListTransactions(c *gin.Context)
	GenerateWalletVerification(c *gin.Context)
	UpdateTrackingInfo(c *gin.Context)
}

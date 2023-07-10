package wallet

import "github.com/gin-gonic/gin"

type IHandler interface {
	ListOwnedWallets(c *gin.Context)
	GetOne(c *gin.Context)
	ListAssets(c *gin.Context)
	ListTransactions(c *gin.Context)
}

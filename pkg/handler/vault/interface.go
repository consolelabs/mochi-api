package vault

import "github.com/gin-gonic/gin"

type IHandler interface {
	CreateVault(c *gin.Context)
	GetVaults(c *gin.Context)
	CreateConfigChannel(c *gin.Context)
	GetVaultConfigChannel(c *gin.Context)
	CreateConfigThreshold(c *gin.Context)
	AddTreasurerToVault(c *gin.Context)
	RemoveTreasurerFromVault(c *gin.Context)
	CreateTreasurerRequest(c *gin.Context)
	CreateTreasurerSubmission(c *gin.Context)
	CreateTreasurerResult(c *gin.Context)
	GetVaultDetail(c *gin.Context)
	TransferVaultToken(c *gin.Context)
	GetTreasurerRequest(c *gin.Context)
	GetVaultTransactions(c *gin.Context)
}

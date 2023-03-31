package vault

import "github.com/gin-gonic/gin"

type IHandler interface {
	CreateVault(c *gin.Context)
	GetVault(c *gin.Context)
	GetVaultInfo(c *gin.Context)
	CreateConfigChannel(c *gin.Context)
	GetVaultConfigChannel(c *gin.Context)
	CreateConfigThreshold(c *gin.Context)
	AddTreasurerToVault(c *gin.Context)
	CreateAddTreasurerRequest(c *gin.Context)
}

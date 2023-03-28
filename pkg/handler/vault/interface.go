package vault

import "github.com/gin-gonic/gin"

type IHandler interface {
	CreateVault(c *gin.Context)
	GetVault(c *gin.Context)
}

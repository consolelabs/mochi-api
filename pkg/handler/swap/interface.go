package swap

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetSwapRoutes(c *gin.Context)
	ExecuteSwapRoutes(c *gin.Context)
	OnchainData(c *gin.Context)
}

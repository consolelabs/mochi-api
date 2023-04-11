package swap

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetSwapRoutes(c *gin.Context)
	BuildSwapRoutes(c *gin.Context)
	ExecuteSwapRoutes(c *gin.Context)
}

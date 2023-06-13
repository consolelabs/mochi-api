package apikey

import "github.com/gin-gonic/gin"

type IHandler interface {
	CreateApiKey(c *gin.Context)
	IntegrateBinanceKey(c *gin.Context)
  UnlinkBinance(c *gin.Context)
}

package dex

import "github.com/gin-gonic/gin"

type IHandler interface {
	SumarizeBinanceAsset(c *gin.Context)
	GetBinanceAssets(c *gin.Context)
	GetBinanceFutures(c *gin.Context)
	GetBinanceSpotTxs(c *gin.Context)
}

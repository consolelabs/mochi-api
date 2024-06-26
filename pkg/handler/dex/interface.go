package dex

import "github.com/gin-gonic/gin"

type IHandler interface {
	SumarizeBinanceAsset(c *gin.Context)
	GetBinanceAssets(c *gin.Context)
	GetBinanceFutures(c *gin.Context)
	GetBinanceSpotTxns(c *gin.Context)
	GetBinanceAverageCosts(c *gin.Context)
}

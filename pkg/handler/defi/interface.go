package defi

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetHistoricalMarketChart(c *gin.Context)
	GetFiatHistoricalExchangeRates(c *gin.Context)

	GetUserWatchlist(c *gin.Context)
	AddToWatchlist(c *gin.Context)
	RemoveFromWatchlist(c *gin.Context)

	GetSupportedTokens(c *gin.Context)
	GetSupportedToken(c *gin.Context)
	GetCoin(c *gin.Context)
	SearchCoins(c *gin.Context)
	CompareToken(c *gin.Context)

	ListAllChain(c *gin.Context)
}

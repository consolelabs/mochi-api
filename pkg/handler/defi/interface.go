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
	GetBinanceCoinData(c *gin.Context)
	AddCoingeckoTokenAlias(c *gin.Context)

	ListAllChain(c *gin.Context)

	AddTokenPriceAlert(c *gin.Context)
	GetUserListPriceAlert(c *gin.Context)
	RemoveTokenPriceAlert(c *gin.Context)
}

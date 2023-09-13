package defi

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetHistoricalMarketChart(c *gin.Context)
	GetFiatHistoricalExchangeRates(c *gin.Context)

	GetSupportedTokens(c *gin.Context)
	GetSupportedToken(c *gin.Context)
	GetCoin(c *gin.Context)
	SearchCoins(c *gin.Context)
	CompareToken(c *gin.Context)
	GetBinanceCoinData(c *gin.Context)
	GetCoinsMarketData(c *gin.Context)

	ListAllChain(c *gin.Context)

	AddTokenPriceAlert(c *gin.Context)
	GetUserListPriceAlert(c *gin.Context)
	RemoveTokenPriceAlert(c *gin.Context)

	GetUserRequestTokens(c *gin.Context)
	CreateUserTokenSupportRequest(c *gin.Context)
	ApproveUserTokenSupportRequest(c *gin.Context)
	RejectUserTokenSupportRequest(c *gin.Context)

	GetGasTracker(c *gin.Context)
	GetChainGasTracker(c *gin.Context)
	GetTrendingSearch(c *gin.Context)
	TopGainerLoser(c *gin.Context)

	FindTokenByContractAddress(c *gin.Context)
	GetTokenInfo(c *gin.Context)

	SearchKeys(c *gin.Context)
	TrackFriendTechKey(c *gin.Context)
	UntrackFriendTechKey(c *gin.Context)
	UpdateFriendTechKeyTrack(c *gin.Context)
	GetUserFriendTechKeyWatchlist(c *gin.Context)
}

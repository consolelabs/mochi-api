package watchlist

import "github.com/gin-gonic/gin"

type IHandler interface {
	// wallets
	ListTrackingWallets(c *gin.Context)
	TrackWallet(c *gin.Context)
	UntrackWallet(c *gin.Context)
	UpdateTrackingWalletInfo(c *gin.Context)

	// tokens
	ListTrackingTokens(c *gin.Context)
	TrackToken(c *gin.Context)
	UntrackToken(c *gin.Context)

	// nfts
	TrackNft(c *gin.Context)
	ListTrackingNfts(c *gin.Context)
	UntrackNft(c *gin.Context)
}

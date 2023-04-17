package tip

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetAllTipBotTokens(c *gin.Context)
	TransferToken(c *gin.Context)

	// onchain
	SubmitOnchainTransfer(c *gin.Context)
	ClaimOnchainTransfer(c *gin.Context)
	GetOnchainTransfers(c *gin.Context)
	GetOnchainBalances(c *gin.Context)
}

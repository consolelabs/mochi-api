package tip

import "github.com/gin-gonic/gin"

type IHandler interface {
	TransferToken(c *gin.Context)

	// onchain
	// SubmitOnchainTransfer(c *gin.Context)
	ClaimOnchainTransfer(c *gin.Context)
	GetOnchainTransfers(c *gin.Context)
	GetOnchainBalances(c *gin.Context)
}

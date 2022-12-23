package tip

import "github.com/gin-gonic/gin"

type IHandler interface {
	OffchainTipBotListAllChains(c *gin.Context)
	OffchainTipBotCreateAssignContract(c *gin.Context)
	OffchainTipBotWithdraw(c *gin.Context)
	GetUserBalances(c *gin.Context)
	TransferToken(c *gin.Context)
	TotalBalances(c *gin.Context)
	TotalOffchainBalances(c *gin.Context)
	TotalFee(c *gin.Context)
	UpdateTokenFee(c *gin.Context)
	GetAllTipBotTokens(c *gin.Context)
	GetTransactionHistoryByQuery(c *gin.Context)
	GetContracts(c *gin.Context)
	HandleDeposit(c *gin.Context)
	GetLatestDeposit(c *gin.Context)
}

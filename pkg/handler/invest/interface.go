package invest

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetInvestList(c *gin.Context)
	OnchainInvestStakeData(c *gin.Context)
	OnchainInvestUnstakeData(c *gin.Context)
}

package invest

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetInvestList(c *gin.Context)
	OnchainInvestData(c *gin.Context)
}

package configchannel

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetGmConfig(c *gin.Context)
	UpsertGmConfig(c *gin.Context)

	GetWelcomeChannelConfig(c *gin.Context)
	UpsertWelcomeChannelConfig(c *gin.Context)
	DeleteWelcomeChannelConfig(c *gin.Context)

	GetSalesTrackerConfig(c *gin.Context)
	CreateSalesTrackerConfig(c *gin.Context)

	GetJoinLeaveChannelConfig(c *gin.Context)
	UpsertJoinLeaveChannelConfig(c *gin.Context)
	DeleteJoinLeaveChannelConfig(c *gin.Context)

	CreateConfigNotify(c *gin.Context)
	ListConfigNotify(c *gin.Context)
	DeleteConfigNotify(c *gin.Context)
}

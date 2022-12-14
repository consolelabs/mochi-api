package configcommunity

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetAllTwitterConfig(c *gin.Context)
	CreateTwitterConfig(c *gin.Context)
	GetTwitterHashtagConfig(c *gin.Context)
	GetAllTwitterHashtagConfig(c *gin.Context)
	DeleteTwitterHashtagConfig(c *gin.Context)
	CreateTwitterHashtagConfig(c *gin.Context)
	AddToTwitterBlackList(c *gin.Context)
	GetTwitterBlackList(c *gin.Context)
	DeleteFromTwitterBlackList(c *gin.Context)

	GetLinkedTelegram(c *gin.Context)
	LinkUserTelegramWithDiscord(c *gin.Context)
}

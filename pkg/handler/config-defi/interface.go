package configdefi

import "github.com/gin-gonic/gin"

type IHandler interface {
	UpsertMonikerConfig(c *gin.Context)
	GetMonikerByGuildID(c *gin.Context)
	DeleteMonikerConfig(c *gin.Context)
	GetDefaultMoniker(c *gin.Context)

	SetGuildDefaultTicker(c *gin.Context)
	GetGuildDefaultTicker(c *gin.Context)

	GetGuildTokens(c *gin.Context)
	UpsertGuildTokenConfig(c *gin.Context)

	GetDefaultToken(c *gin.Context)
	ConfigDefaultToken(c *gin.Context)
	RemoveDefaultToken(c *gin.Context)
	CreateDefaultCollectionSymbol(c *gin.Context)

	ListAllCustomToken(c *gin.Context)
	HandlerGuildCustomTokenConfig(c *gin.Context)

	GetGuildDefaultCurrency(c *gin.Context)
	UpsertGuildDefaultCurrency(c *gin.Context)
	DeleteGuildDefaultCurrency(c *gin.Context)
}

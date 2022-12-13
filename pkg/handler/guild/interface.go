package guild

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetGuilds(c *gin.Context)
	GetGuild(c *gin.Context)
	CreateGuild(c *gin.Context)
	GetGuildStatsHandler(c *gin.Context)
	CreateGuildChannel(c *gin.Context)
	ListMyGuilds(c *gin.Context)
	UpdateGuild(c *gin.Context)
}

package guild

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetGuilds(c *gin.Context)
	GetGuild(c *gin.Context)
	ListMyGuilds(c *gin.Context)
	UpdateGuild(c *gin.Context)
	GetGuildRoles(c *gin.Context)
	ValidateUser(c *gin.Context)
}

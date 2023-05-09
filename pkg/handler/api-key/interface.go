package apikey

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetApiKeyByDiscordId(c *gin.Context)
	CreateApiKeyByDiscordId(c *gin.Context)
}

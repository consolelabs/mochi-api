package config

import "github.com/gin-gonic/gin"

type IHandler interface {
	ToggleActivityConfig(c *gin.Context)
	GetListCommandPermissions(c *gin.Context)
	GetInstallBotUrl(c *gin.Context)
}

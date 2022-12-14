package config

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetGuildPruneExclude(c *gin.Context)
	UpsertGuildPruneExclude(c *gin.Context)
	DeleteGuildPruneExclude(c *gin.Context)

	ToggleActivityConfig(c *gin.Context)
}

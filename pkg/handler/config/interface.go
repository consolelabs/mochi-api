package config

import "github.com/gin-gonic/gin"

type IHandler interface {
	ToggleActivityConfig(c *gin.Context)
}

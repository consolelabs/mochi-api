package auto

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetAutoTriggers(c *gin.Context)
}

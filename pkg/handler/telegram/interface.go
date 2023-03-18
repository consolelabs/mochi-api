package telegram

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetByUsername(c *gin.Context)
}

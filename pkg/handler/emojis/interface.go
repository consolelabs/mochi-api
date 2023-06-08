package emojis

import "github.com/gin-gonic/gin"

type IHandler interface {
	ListEmojis(c *gin.Context)
}

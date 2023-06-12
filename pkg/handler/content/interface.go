package content

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetTypeContent(c *gin.Context)
}

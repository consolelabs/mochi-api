package tip

import "github.com/gin-gonic/gin"

type IHandler interface {
	TransferToken(c *gin.Context)
}

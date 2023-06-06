package pkpass

import "github.com/gin-gonic/gin"

type IHandler interface {
	GeneratePkPass(c *gin.Context)
}

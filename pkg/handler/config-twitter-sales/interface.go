package configtwittersale

import "github.com/gin-gonic/gin"

type IHandler interface {
	Get(c *gin.Context)
	Create(c *gin.Context)
}

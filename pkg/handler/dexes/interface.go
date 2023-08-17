package dexes

import "github.com/gin-gonic/gin"

type IHandler interface {
	SearchDexPair(c *gin.Context)
}

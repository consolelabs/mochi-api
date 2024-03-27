package dexes

import "github.com/gin-gonic/gin"

type IHandler interface {
	SearchDexPair(c *gin.Context)
	SearchDexScreenerPair(c *gin.Context)
}

package cache

import "github.com/gin-gonic/gin"

type IHandler interface {
	SetUpvoteMessageCache(c *gin.Context)
}

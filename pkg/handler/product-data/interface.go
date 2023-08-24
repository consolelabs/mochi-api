package productdata

import "github.com/gin-gonic/gin"

type IHandler interface {
	ProductBotCommand(c *gin.Context)
	ProductChangelogs(c *gin.Context)
	CrawlChangelogs(c *gin.Context)
}

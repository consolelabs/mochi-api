package productdata

import "github.com/gin-gonic/gin"

type IHandler interface {
	ProductBotCommand(c *gin.Context)
	ProductChangelogs(c *gin.Context)
	GetProductChangelogByVersion(c *gin.Context)
	CrawlChangelogs(c *gin.Context)
	PublishChangelog(c *gin.Context)
	CreateProductChangelogsView(c *gin.Context)
	GetProductChangelogsView(c *gin.Context)
	GetProductHashtag(c *gin.Context)
	GetProductTheme(c *gin.Context)
}

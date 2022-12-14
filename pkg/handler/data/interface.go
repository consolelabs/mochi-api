package data

import "github.com/gin-gonic/gin"

type IHandler interface {
	AddServersUsageStat(c *gin.Context)
	AddGitbookClick(c *gin.Context)
	MetricByProperties(c *gin.Context)
	MetricNftCollection(c *gin.Context, query string)
	MetricActiveUsers(c *gin.Context, query string, guildId string)
	MetricTotalServers(c *gin.Context, query string)
	MetricVerifiedWallets(c *gin.Context, query string, guildId string)
	MetricSupportedTokens(c *gin.Context, query string, guildId string)
	MetricCommandUsage(c *gin.Context, query string, guildId string)
}

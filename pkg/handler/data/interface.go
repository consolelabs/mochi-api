package data

import "github.com/gin-gonic/gin"

type IHandler interface {
	AddGitbookClick(c *gin.Context)
	MetricProposalUsage(c *gin.Context)
	MetricDaoTracker(c *gin.Context)
	MetricByProperties(c *gin.Context)
	MetricNftCollection(c *gin.Context, query string)
	MetricActiveUsers(c *gin.Context, query string, guildId string)
	MetricTotalServers(c *gin.Context, query string)
	MetricVerifiedWallets(c *gin.Context, query string, guildId string)
	MetricSupportedTokens(c *gin.Context, query string, guildId string)
}

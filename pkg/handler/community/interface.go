package community

import "github.com/gin-gonic/gin"

type IHandler interface {
	HandleUserFeedback(c *gin.Context)
	UpdateUserFeedback(c *gin.Context)
	GetAllUserFeedback(c *gin.Context)

	GetUserQuestList(c *gin.Context)
	ClaimQuestsRewards(c *gin.Context)
	UpdateQuestProgress(c *gin.Context)

	GetRepostReactionConfigs(c *gin.Context)
	ConfigRepostReaction(c *gin.Context)
	RemoveRepostReactionConfig(c *gin.Context)
	CreateConfigRepostReactionConversation(c *gin.Context)
	RemoveConfigRepostReactionConversation(c *gin.Context)
	EditMessageRepost(c *gin.Context)

	CreateBlacklistChannelRepostConfig(c *gin.Context)
	GetGuildBlacklistChannelRepostConfig(c *gin.Context)
	DeleteBlacklistChannelRepostConfig(c *gin.Context)

	CreateTwitterPost(c *gin.Context)
	GetTwitterLeaderboard(c *gin.Context)

	UpsertLevelUpMessage(c *gin.Context)
	GetLevelUpMessage(c *gin.Context)
	DeleteLevelUpMessage(c *gin.Context)

	GetAllAd(c *gin.Context)
	GetAdById(c *gin.Context)
	CreateAd(c *gin.Context)
	InitAdSubmission(c *gin.Context)
	DeleteAdById(c *gin.Context)
	UpdateAdById(c *gin.Context)
}

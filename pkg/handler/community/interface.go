package community

import "github.com/gin-gonic/gin"

type IHandler interface {
	HandleUserFeedback(c *gin.Context)
	UpdateUserFeedback(c *gin.Context)
	GetAllUserFeedback(c *gin.Context)

	GetUserQuestList(c *gin.Context)
	ClaimQuestsRewards(c *gin.Context)
	UpdateQuestProgress(c *gin.Context)

	UpsertLevelUpMessage(c *gin.Context)
	GetLevelUpMessage(c *gin.Context)
	DeleteLevelUpMessage(c *gin.Context)

	GetAllAd(c *gin.Context)
	GetAdById(c *gin.Context)
	CreateAd(c *gin.Context)
	InitAdSubmission(c *gin.Context)
	DeleteAdById(c *gin.Context)
	UpdateAdById(c *gin.Context)

	GetUserTag(c *gin.Context)
	UpsertUserTag(c *gin.Context)
}

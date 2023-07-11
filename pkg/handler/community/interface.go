package community

import "github.com/gin-gonic/gin"

type IHandler interface {
	HandleUserFeedback(c *gin.Context)
	UpdateUserFeedback(c *gin.Context)
	GetAllUserFeedback(c *gin.Context)

	GetUserQuestList(c *gin.Context)
	ClaimQuestsRewards(c *gin.Context)
	UpdateQuestProgress(c *gin.Context)

	GetUserTag(c *gin.Context)
	UpsertUserTag(c *gin.Context)
}

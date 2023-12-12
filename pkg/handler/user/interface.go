package user

import "github.com/gin-gonic/gin"

type IHandler interface {
	IndexUsers(c *gin.Context)
	GetUser(c *gin.Context)
	GetUserCurrentGMStreak(c *gin.Context)
	GetTopUsers(c *gin.Context)
	GetUserProfile(c *gin.Context)

	SendUserXP(c *gin.Context)
	GetUserBalance(c *gin.Context)
	GetGlobalProfileInfo(c *gin.Context)
}

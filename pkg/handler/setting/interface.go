package setting

import "github.com/gin-gonic/gin"

type IHandler interface {
	// general
	GetUserGeneralSettings(c *gin.Context)
	UpdateUserGeneralSettings(c *gin.Context)

	// notification
	GetUserNotificationSettings(c *gin.Context)
	UpdateUserNotificationSettings(c *gin.Context)
}

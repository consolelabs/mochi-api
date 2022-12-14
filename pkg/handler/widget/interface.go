package widget

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetUserTokenAlert(c *gin.Context)
	UpsertUserTokenAlert(c *gin.Context)
	DeleteUserTokenAlert(c *gin.Context)
	GetUserDevice(c *gin.Context)
	UpsertUserDevice(c *gin.Context)
	DeleteUserDevice(c *gin.Context)
}

package earn

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetEarnInfoList(c *gin.Context)
	CreateEarnInfo(c *gin.Context)
	CreateUserEarn(c *gin.Context)
	GetUserEarnListByUserId(c *gin.Context)
	DeleteUserEarn(c *gin.Context)
}

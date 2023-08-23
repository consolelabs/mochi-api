package onboarding

import "github.com/gin-gonic/gin"

type IHandler interface {
	Start(c *gin.Context)
}

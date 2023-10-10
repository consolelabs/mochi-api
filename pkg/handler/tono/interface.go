package tono

import "github.com/gin-gonic/gin"

type IHandler interface {
	TonoCommandPermissions(c *gin.Context)
}

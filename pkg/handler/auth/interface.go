package auth

import "github.com/gin-gonic/gin"

type IHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

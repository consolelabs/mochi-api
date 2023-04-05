package verify

import "github.com/gin-gonic/gin"

type IHandler interface {
	NewGuildConfigWalletVerificationMessage(c *gin.Context)
	GetGuildConfigWalletVerificationMessage(c *gin.Context)
	UpdateGuildConfigWalletVerificationMessage(c *gin.Context)
	DeleteGuildConfigWalletVerificationMessage(c *gin.Context)
	GenerateVerification(c *gin.Context)
	VerifyWalletAddress(c *gin.Context)
	AssignVerifiedRole(c *gin.Context)
}

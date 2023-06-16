package airdropcampaign

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetAirdropCampaigns(c *gin.Context)
	CreateAirdropCampaign(c *gin.Context)
	GetAirdropCampaignStats(c *gin.Context)
	CreateProfileAirdropCampaign(c *gin.Context)
	GetProfileAirdropCampaigns(c *gin.Context)
	DeleteProfileAirdropCampaign(c *gin.Context)
}

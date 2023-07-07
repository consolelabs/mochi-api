package configrole

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetAllRoleReactionConfigs(c *gin.Context)
	AddReactionRoleConfig(c *gin.Context)
	RemoveReactionRoleConfig(c *gin.Context)
	FilterConfigByReaction(c *gin.Context)

	GetDefaultRolesByGuildID(c *gin.Context)
	CreateDefaultRole(c *gin.Context)
	DeleteDefaultRoleByGuildID(c *gin.Context)

	ListGuildGroupNFTRoles(c *gin.Context)
	NewGuildGroupNFTRole(c *gin.Context)
	RemoveGuildNFTRole(c *gin.Context)
	RemoveGuildGroupNFTRole(c *gin.Context)

	ConfigLevelRole(c *gin.Context)
	GetLevelRoleConfigs(c *gin.Context)
	RemoveLevelRoleConfig(c *gin.Context)

	CreateGuildTokenRole(c *gin.Context)
	ListGuildTokenRoles(c *gin.Context)
	UpdateGuildTokenRole(c *gin.Context)
	RemoveGuildTokenRole(c *gin.Context)

	CreateGuildXPRole(c *gin.Context)
	ListGuildXPRoles(c *gin.Context)
	RemoveGuildXPRole(c *gin.Context)

	CreateGuildAdminRoles(c *gin.Context)
	ListGuildAdminRoles(c *gin.Context)
	RemoveGuildAdminRole(c *gin.Context)
}

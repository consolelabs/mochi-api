package handler

import (
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllDefaultRoles(c *gin.Context) {
	data, err := h.entities.GetAllDefaultRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *Handler) CreateDefaultRole(c *gin.Context) {
	roleID, isExist := c.GetQuery("role_id")
	if !isExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role_id is required"})
	}

	guildID, isExist := c.GetQuery("guild_id")
	if !isExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}

	if err := h.entities.CreateDefaultRoleConfig(guildID, roleID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defaultRole := response.DefaultRole{
		RoleID:  roleID,
		GuildID: guildID,
	}

	c.JSON(http.StatusOK, response.DefaultRoleCreationResponse{
		Data:    defaultRole,
		Success: true,
	})
}

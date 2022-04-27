package handler

import (
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetDefaultRolesByGuildID(c *gin.Context) {
	guildID, isExist := c.GetQuery("guild_id")
	if !isExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}

	data, err := h.entities.GetAllDefaultRoles(guildID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *Handler) CreateDefaultRole(c *gin.Context) {
	body := request.CreateDefaultRoleRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.CreateDefaultRoleConfig(body.GuildID, body.RoleID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defaultRole := response.DefaultRole{
		RoleID:  body.GuildID,
		GuildID: body.RoleID,
	}

	c.JSON(http.StatusOK, response.DefaultRoleCreationResponse{
		Data:    defaultRole,
		Success: true,
	})
}

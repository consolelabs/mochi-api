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

	data, err := h.entities.GetDefaultRoleByGuildID(guildID)
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

	c.JSON(http.StatusOK, response.DefaultRoleResponse{
		Data:    defaultRole,
		Success: true,
	})
}

func (h *Handler) DeleteDefaultRoleByGuildID(c *gin.Context) {
	type DeleteResponse struct {
		Success bool `json:"success"`
	}
	guildID, isExist := c.GetQuery("guild_id")
	if !isExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}
	err := h.entities.DeleteDefaultRoleConfig(guildID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &DeleteResponse{
		Success: true,
	})
}

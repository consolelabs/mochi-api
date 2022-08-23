package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetDefaultRolesByGuildID(c *gin.Context) {
	guildID, isExist := c.GetQuery("guild_id")
	if !isExist {
		h.log.Info("[handler.GetDefaultRolesByGuildID] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}

	data, err := h.entities.GetDefaultRoleByGuildID(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetDefaultRolesByGuildID] - failed to get default roles")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *Handler) CreateDefaultRole(c *gin.Context) {
	body := request.CreateDefaultRoleRequest{}

	if err := c.BindJSON(&body); err != nil {
		h.log.Fields(logger.Fields{"body": body}).Error(err, "[handler.CreateDefaultRole] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.CreateDefaultRoleConfig(body.GuildID, body.RoleID); err != nil {
		h.log.Fields(logger.Fields{"guildID": body.GuildID, "roleID": body.RoleID}).Error(err, "[handler.CreateDefaultRole] - failed to create default role config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defaultRole := response.DefaultRole{
		RoleID:  body.RoleID,
		GuildID: body.GuildID,
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
		h.log.Info("[handler.DeleteDefaultRoleByGuildID] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}
	err := h.entities.DeleteDefaultRoleConfig(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.DeleteDefaultRoleByGuildID] - failed to delete default role config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &DeleteResponse{
		Success: true,
	})
}

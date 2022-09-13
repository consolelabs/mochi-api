package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

// GetDefaultRolesByGuildID     godoc
// @Summary     Get default roles by guild id
// @Description Get default roles by guild id
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.DefaultRoleResponse
// @Router      /configs/default-roles [get]
func (h *Handler) GetDefaultRolesByGuildID(c *gin.Context) {
	guildID, isExist := c.GetQuery("guild_id")
	if !isExist {
		h.log.Info("[handler.GetDefaultRolesByGuildID] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	data, err := h.entities.GetDefaultRoleByGuildID(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetDefaultRolesByGuildID] - failed to get default roles")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

// CreateDefaultRole     godoc
// @Summary     Create default role
// @Description Create default role
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.CreateDefaultRoleRequest true "Create default role request"
// @Success     200 {object} response.DefaultRoleResponse
// @Router      /configs/default-roles [post]
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
		Data: defaultRole,
		Ok:   true,
	})
}

// DeleteDefaultRole     godoc
// @Summary     Delete default role
// @Description Delete default role
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.ResponseSucess
// @Router      /configs/default-roles [delete]
func (h *Handler) DeleteDefaultRoleByGuildID(c *gin.Context) {
	guildID, isExist := c.GetQuery("guild_id")
	if !isExist {
		h.log.Info("[handler.DeleteDefaultRoleByGuildID] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}
	err := h.entities.DeleteDefaultRoleConfig(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.DeleteDefaultRoleByGuildID] - failed to delete default role config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ResponseSucess{
		Success: true,
	})
}

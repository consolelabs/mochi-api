package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

// GetAllRoleReactionConfigs     godoc
// @Summary     Get all role reaction configs
// @Description Get all role reaction configs
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.DataListRoleReactionResponse
// @Router      /configs/reaction-roles [get]
func (h *Handler) GetAllRoleReactionConfigs(c *gin.Context) {
	guildID, guildIDExist := c.GetQuery("guild_id")
	if !guildIDExist {
		h.log.Info("[handler.GetAllRoleReactionConfigs] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	resp, err := h.entities.ListAllReactionRoles(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetAllRoleReactionConfigs] - failed to list all reaction roles")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.DataListRoleReactionResponse{
		Data: *resp,
	})

}

// AddReactionRoleConfig     godoc
// @Summary     Add reaction role config
// @Description Add reaction role config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.RoleReactionUpdateRequest true "Add reaction role config request"
// @Success     200 {object} response.RoleReactionConfigResponse
// @Router      /configs/reaction-roles [post]
func (h *Handler) AddReactionRoleConfig(c *gin.Context) {
	var req request.RoleReactionUpdateRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.AddReactionRoleConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config, err := h.entities.UpdateConfigByMessageID(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.AddReactionRoleConfig] - failed to update config my message id")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, config)
}

// RemoveReactionRoleConfig     godoc
// @Summary     Remove reaction role config
// @Description Remove reaction role config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.RoleReactionUpdateRequest true "Remove reaction role config request"
// @Success     200 {object} response.ResponseSucess
// @Router      /configs/reaction-roles [delete]
func (h *Handler) RemoveReactionRoleConfig(c *gin.Context) {
	var req request.RoleReactionUpdateRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.RemoveReactionRoleConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err error

	if req.RoleID != "" && req.Reaction != "" {
		err = h.entities.RemoveSpecificRoleReaction(req)
	} else {
		err = h.entities.ClearReactionMessageConfig(req)
	}

	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.RemoveReactionRoleConfig] - failed to remove reaction config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.ResponseSucess{
		Success: true,
	})
}

// FilterConfigByReaction     godoc
// @Summary     Filter config by reaction
// @Description Filter config by reaction
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.RoleReactionRequest true "Filter config by reaction request"
// @Success     200 {object} response.DataFilterConfigByReaction
// @Router      /configs/reaction-roles/filter [post]
func (h *Handler) FilterConfigByReaction(c *gin.Context) {
	var req request.RoleReactionRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.FilterConfigByReaction] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config, err := h.entities.GetReactionRoleByMessageID(req.GuildID, req.MessageID, req.Reaction)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "messageID": req.MessageID, "reaction": req.Reaction}).Error(err, "[handler.FilterConfigByReaction] - failed to get reaction role by message id")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.DataFilterConfigByReaction{
		Data: config,
	})
}

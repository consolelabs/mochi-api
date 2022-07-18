package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

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

	c.JSON(http.StatusOK, resp)

}

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
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

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

	c.JSON(http.StatusOK, config)
}

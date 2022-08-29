package handler

import (
	"fmt"
	"net/http"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateCustomCommand     godoc
// @Summary     Create custom command
// @Description Create custom command
// @Tags        Custom Command
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Param       Request  body model.GuildCustomCommand true "Create custom command request"
// @Success     200 {object} response.CreateCustomCommandResponse
// @Router      /guilds/{guild_id}/custom-commands [post]
func (h *Handler) CreateCustomCommand(c *gin.Context) {
	var (
		guildID = c.Param("guild_id")
	)

	if guildID == "" {
		h.log.Info("[handler.CreateCustomCommand] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required."})
		return
	}

	var customCommand model.GuildCustomCommand

	if err := c.BindJSON(&customCommand); err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "body": customCommand}).Error(err, "[handler.CreateCustomCommand] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customCommand.GuildID = guildID
	customCommand.Enabled = true

	if err := h.entities.CreateCustomCommand(customCommand); err != nil {
		if err == gorm.ErrRecordNotFound {
			h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.CreateCustomCommand] - failed to find guild")
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("guild %s not found", customCommand.GuildID)})
			return
		}
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.CreateCustomCommand] - failed to create custom command")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.CreateCustomCommandResponse{Data: customCommand})
}

// UpdateCustomCommand     godoc
// @Summary     Update custom command
// @Description Update custom command
// @Tags        Custom Command
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Param       command_id   path  string true  "Command ID"
// @Param       Request  body model.GuildCustomCommand true "Update custom command request"
// @Success     200 {object} response.UpdateCustomCommandResponse
// @Router      /guilds/{guild_id}/custom-commands/{command_id} [put]
func (h *Handler) UpdateCustomCommand(c *gin.Context) {
	var (
		guildID = c.Param("guild_id")
		ID      = c.Param("command_id")
	)

	if guildID == "" {
		h.log.Info("[handler.UpdateCustomCommand] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required."})
		return
	}

	if ID == "" {
		h.log.Info("[handler.UpdateCustomCommand] - command id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required."})
		return
	}

	var customCommand model.GuildCustomCommand

	if err := c.BindJSON(&customCommand); err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "id": ID, "body": customCommand}).Error(err, "[handler.UpdateCustomCommand] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customCommand.GuildID = guildID

	if err := h.entities.UpdateCustomCommand(ID, guildID, customCommand); err != nil {
		if err == gorm.ErrRecordNotFound {
			h.log.Fields(logger.Fields{"guildID": guildID, "id": ID, "body": customCommand}).Error(err, "[handler.UpdateCustomCommand] - failed to find guild")
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("command %s of guild %s not found", ID, guildID)})
			return
		}
		h.log.Fields(logger.Fields{"guildID": guildID, "id": ID, "body": customCommand}).Error(err, "[handler.UpdateCustomCommand] - failed to update custom command")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.UpdateCustomCommandResponse{Data: customCommand})
}

// ListCustomCommands     godoc
// @Summary     List custom commands
// @Description List custom commands
// @Tags        Custom Command
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Param       enabled   query  string true  "Enabled"
// @Success     200 {object} response.ListCustomCommandsResponse
// @Router      /guilds/{guild_id}/custom-commands [get]
func (h *Handler) ListCustomCommands(c *gin.Context) {
	var (
		guildID    = c.Param("guild_id")
		enabledStr = c.Query("enabled")
		enabled    bool
		enabledQ   *bool
	)

	if guildID == "" {
		h.log.Info("[handler.ListCustomCommands] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required."})
		return
	}

	switch enabledStr {
	case "true":
		enabled = true
		enabledQ = &enabled
	case "false":
		enabledQ = &enabled
	}

	customCommands, err := h.entities.ListCustomCommands(guildID, enabledQ)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "enabled": enabledStr}).Error(err, "[handler.ListCustomCommands] - failed to list all custom commands")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ListCustomCommandsResponse{Data: customCommands})
}

// GetCustomCommand     godoc
// @Summary     Get custom commands
// @Description Get custom commands
// @Tags        Custom Command
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Param       command_id   path  string true  "Command ID"
// @Success     200 {object} response.GetCustomCommandResponse
// @Router      /guilds/{guild_id}/custom-commands/{command_id} [get]
func (h *Handler) GetCustomCommand(c *gin.Context) {
	var (
		guildID = c.Param("guild_id")
		ID      = c.Param("command_id")
	)

	if guildID == "" {
		h.log.Info("[handler.GetCustomCommand] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required."})
		return
	}

	if ID == "" {
		h.log.Info("[handler.GetCustomCommand] - command id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required."})
		return
	}

	customCommand, err := h.entities.GetCustomCommand(ID, guildID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			h.log.Fields(logger.Fields{"guildID": guildID, "id": ID}).Error(err, "[handler.GetCustomCommand] - failed to find guild")
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("command %s of guild %s not found", ID, guildID)})
			return
		}
		h.log.Fields(logger.Fields{"guildID": guildID, "id": ID}).Error(err, "[handler.GetCustomCommand] - failed to get custom command")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.GetCustomCommandResponse{Data: customCommand})
}

// DeleteCustomCommand     godoc
// @Summary     Delete custom commands
// @Description Delete custom commands
// @Tags        Custom Command
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Param       command_id   path  string true  "Command ID"
// @Success     204
// @Router      /guilds/{guild_id}/custom-commands/{command_id} [delete]
func (h *Handler) DeleteCustomCommand(c *gin.Context) {
	var (
		guildID = c.Param("guild_id")
		ID      = c.Param("command_id")
	)

	if guildID == "" {
		h.log.Info("[handler.DeleteCustomCommand] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required."})
		return
	}

	if ID == "" {
		h.log.Info("[handler.DeleteCustomCommand] - command id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required."})
		return
	}

	if err := h.entities.DeleteCustomCommand(ID, guildID); err != nil {
		if err == gorm.ErrRecordNotFound {
			h.log.Fields(logger.Fields{"guildID": guildID, "id": ID}).Error(err, "[handler.DeleteCustomCommand] - failed to find guild")
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("command %s of guild %s not found", ID, guildID)})
			return
		}
		h.log.Fields(logger.Fields{"guildID": guildID, "id": ID}).Error(err, "[handler.DeleteCustomCommand] - failed to delete custom command")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

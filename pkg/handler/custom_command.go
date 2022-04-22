package handler

import (
	"fmt"
	"net/http"

	"github.com/defipod/mochi/pkg/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) CreateCustomCommand(c *gin.Context) {
	var (
		guildID = c.Param("guild_id")
	)

	if guildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required."})
		return
	}

	var customCommand model.GuildCustomCommand

	if err := c.BindJSON(&customCommand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customCommand.GuildID = guildID
	customCommand.Enabled = true

	if err := h.entities.CreateCustomCommand(customCommand); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("guild %s not found", customCommand.GuildID)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": customCommand})
}

func (h *Handler) UpdateCustomCommand(c *gin.Context) {
	var (
		guildID = c.Param("guild_id")
		ID      = c.Param("command_id")
	)

	if guildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required."})
		return
	}

	if ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required."})
		return
	}

	var customCommand model.GuildCustomCommand

	if err := c.BindJSON(&customCommand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customCommand.GuildID = guildID

	if err := h.entities.UpdateCustomCommand(ID, guildID, customCommand); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("command %s of guild %s not found", ID, guildID)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customCommand})
}

func (h *Handler) ListCustomCommands(c *gin.Context) {
	var (
		guildID    = c.Param("guild_id")
		enabledStr = c.Query("enabled")
		enabled    bool
		enabledQ   *bool
	)

	if guildID == "" {
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customCommands})
}

func (h *Handler) GetCustomCommand(c *gin.Context) {
	var (
		guildID = c.Param("guild_id")
		ID      = c.Param("command_id")
	)

	if guildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required."})
		return
	}

	if ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required."})
		return
	}

	customCommand, err := h.entities.GetCustomCommand(ID, guildID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("command %s of guild %s not found", ID, guildID)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customCommand})
}

func (h *Handler) DeleteCustomCommand(c *gin.Context) {
	var (
		guildID = c.Param("guild_id")
		ID      = c.Param("command_id")
	)

	if guildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required."})
		return
	}

	if ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required."})
		return
	}

	if err := h.entities.DeleteCustomCommand(ID, guildID); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("command %s of guild %s not found", ID, guildID)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllReactionRolesByGuildID(c *gin.Context) {

	guildID, guildIDExist := c.GetQuery("guild_id")
	if !guildIDExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	configs, err := h.entities.GetReactionRoles(guildID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, configs)
}

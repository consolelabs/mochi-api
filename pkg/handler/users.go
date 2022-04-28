package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) IndexUsers(c *gin.Context) {
	var req request.CreateUserRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.CreateUser(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (h *Handler) GetUser(c *gin.Context) {
	discordID := c.Param("id")
	if discordID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	user, err := h.entities.GetUser(discordID)
	if err != nil {
		if err == entities.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *Handler) GetUserCurrentGMStreak(c *gin.Context) {

	discordID := c.Query("discord_id")
	guildID := c.Query("guild_id")

	if discordID == "" || guildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "discord_id and guild_id is required"})
		return
	}

	res, code, err := h.entities.GetUserCurrentGMStreak(discordID, guildID)
	if err != nil {
		if code >= 500 {
			logrus.WithError(err).Errorf("Failed to get user gm streak discord_id: %s guild_id: %s err: %s", discordID, guildID, err.Error())
		}
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(code, gin.H{"data": res})
}

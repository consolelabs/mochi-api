package handler

import (
	"net/http"
	"strconv"

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

func (h *Handler) GetMyInfo(c *gin.Context) {
	accessToken := c.GetString("discord_access_token")

	du, err := h.entities.GetMyDiscordInfo(accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": du})
}

func (h *Handler) GetTopUsers(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = "0"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "page must be an integer"})
		return
	}

	limitStr := c.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "limit must be an integer"})
		return
	}

	data, err := h.entities.GetTopUsers(guildID, userID, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

package handler

import (
	"net/http"
	"strconv"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

func (h *Handler) IndexUsers(c *gin.Context) {
	var req request.CreateUserRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.IndexUsers] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.CreateUser(req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.IndexUsers] - failed to create user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (h *Handler) GetUser(c *gin.Context) {
	discordID := c.Param("id")
	if discordID == "" {
		h.log.Info("[handler.GetUser] - discord id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	user, err := h.entities.GetUser(discordID)
	if err != nil {
		if err == entities.ErrRecordNotFound {
			h.log.Fields(logger.Fields{"discordId": discordID}).Error(err, "[handler.GetUser] - users not found")
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		h.log.Fields(logger.Fields{"discordId": discordID}).Error(err, "[handler.GetUser] - failed to get user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *Handler) GetUserCurrentGMStreak(c *gin.Context) {

	discordID := c.Query("discord_id")
	guildID := c.Query("guild_id")

	if discordID == "" || guildID == "" {
		h.log.Infof("[handler.GetUserCurrentGMStreak] - missing params, discordID: %v, guildID: %v", discordID, guildID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "discord_id and guild_id is required"})
		return
	}

	res, code, err := h.entities.GetUserCurrentGMStreak(discordID, guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"discordId": discordID, "guildID": guildID}).Error(err, "[handler.GetUserCurrentGMStreak] - failed to get user current gm streak")
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(code, gin.H{"data": res})
}

func (h *Handler) GetMyInfo(c *gin.Context) {
	accessToken := c.GetString("discord_access_token")

	du, err := h.entities.GetMyDiscordInfo(accessToken)
	if err != nil {
		h.log.Fields(logger.Fields{"token": accessToken}).Error(err, "[handler.GetMyInfo] - failed to get discord info")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": du})
}

func (h *Handler) GetTopUsers(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetTopUsers] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	userID := c.Query("user_id")
	if userID == "" {
		h.log.Info("[handler.GetTopUsers] - user id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = "0"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		h.log.Fields(logger.Fields{"page": pageStr}).Error(err, "[handler.GetTopUsers] - invalid page")
		c.JSON(http.StatusBadRequest, gin.H{"error": "page must be an integer"})
		return
	}

	limitStr := c.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		h.log.Fields(logger.Fields{"limit": limit}).Error(err, "[handler.GetTopUsers] - invalid limit")
		c.JSON(http.StatusBadRequest, gin.H{"error": "limit must be an integer"})
		return
	}

	data, err := h.entities.GetTopUsers(guildID, userID, limit, page)
	if err != nil {
		h.log.Fields(logger.Fields{"page": pageStr, "limit": limit, "guildID": guildID, "userID": userID}).Error(err, "[handler.GetTopUsers] - failed to get top users")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) GetUserProfile(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetUserProfile] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	userID := c.Query("user_id")
	if userID == "" {
		h.log.Info("[handler.GetUserProfile] - user id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	data, err := h.entities.GetUserProfile(guildID, userID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "userID": userID}).Error(err, "[handler.GetUserProfile] - failed to get user profile")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

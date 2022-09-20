package handler

import (
	"net/http"
	"strconv"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

// IndexUsers     godoc
// @Summary     Create User
// @Description Create User
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       Request  body request.CreateUserRequest true "Create user request"
// @Success     200 {object} response.ResponseMessage
// @Router      /users [post]
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

	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
}

// GetUser     godoc
// @Summary     Get user
// @Description Get user
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       user_id path     string true "User ID"
// @Success     200 {object} response.GetUserResponse
// @Router      /users/{user_id} [get]
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
			c.JSON(http.StatusOK, gin.H{"data": nil})
			return
		}
		h.log.Fields(logger.Fields{"discordId": discordID}).Error(err, "[handler.GetUser] - failed to get user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GetUserCurrentGMStreak     godoc
// @Summary     Get user current GM streak
// @Description Get user current GM streak
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       discord_id query     string true "Discord ID"
// @Param       guild_id query     string true "Guild ID"
// @Success     200 {object} response.GetUserCurrentGMStreakResponse
// @Router      /users/gmstreak [get]
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

// GetUserCurrentUpvoteStreak     godoc
// @Summary     Get user current upvote streak
// @Description Get user current upvote streak
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       discord_id query     string true "Discord ID"
// @Success     200 {object} response.CurrentUserUpvoteStreakResponse
// @Router      /users/upvote-streak [get]
func (h *Handler) GetUserCurrentUpvoteStreak(c *gin.Context) {

	discordID := c.Query("discord_id")
	if discordID == "" {
		h.log.Infof("[handler.GetUserCurrentUpvoteStreak] - missing params, discordID: %v", discordID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "discord_id is required"})
		return
	}

	res, code, err := h.entities.GetUserCurrentUpvoteStreak(discordID)
	if err != nil {
		h.log.Fields(logger.Fields{"discordId": discordID}).Error(err, "[handler.GetUserCurrentUpvoteStreak] - failed to get user current upvote streak")
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(code, gin.H{"data": res})
}

// GetUserUpvoteLeaderboard     godoc
// @Summary     Get user upvote leaderboard
// @Description Get user upvote leaderboard
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       by query     string true "streak / total"
// @Param       guild_id query     string true "Guild ID"
// @Success     200 {object} response.GetUserUpvoteLeaderboardResponse
// @Router      /users/upvote-leaderboard [get]
func (h *Handler) GetUserUpvoteLeaderboard(c *gin.Context) {
	by := c.Query("by")
	if by == "" {
		by = "total"
	}
	guildId := c.Query("guild_id")
	res, err := h.entities.GetUpvoteLeaderboard(by, guildId)
	if err != nil {
		h.log.Error(err, "[handler.GetUserUpvoteLeaderboard] - failed to get upvote leaderboard by total")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.GetUserUpvoteLeaderboardResponse{
		Message: "ok",
		Data:    &res,
	})
}

// GetMyInfo     godoc
// @Summary     Get user info
// @Description Get user info
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       Authorization header   string true "Authorization"
// @Success     200 {object} response.GetMyInfoResponse
// @Router      /users/me [get]
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

// GetTopUsers     godoc
// @Summary     Get top users
// @Description Get top users
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       guild_id query     string true "Guild ID"
// @Param       user_id query     string true "User ID"
// @Param       page query     int false "Page" default(0)
// @Param       limit query     int false "Limit" default(10)
// @Success     200 {object} response.GetMyInfoResponse
// @Router      /users/top [get]
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

// GetUserProfile     godoc
// @Summary     Get user profile
// @Description Get user profile
// @Tags        Profile
// @Accept      json
// @Produce     json
// @Param       guild_id query     string true "Guild ID"
// @Param       user_id query     string true "User ID"
// @Success     200 {object} response.GetDataUserProfileResponse
// @Router      /profiles/ [get]
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

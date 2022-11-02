package handler

import (
	"errors"
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
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := h.entities.CreateUser(req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.IndexUsers] - failed to create user")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
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
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("id is required"), nil))
		return
	}

	user, err := h.entities.GetUser(discordID)
	if err != nil {
		if err == entities.ErrRecordNotFound {
			h.log.Fields(logger.Fields{"discordId": discordID}).Error(err, "[handler.GetUser] - users not found")
			c.JSON(http.StatusOK, response.CreateResponse[any](nil, nil, nil, nil))
			return
		}
		h.log.Fields(logger.Fields{"discordId": discordID}).Error(err, "[handler.GetUser] - failed to get user")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(user, nil, nil, nil))
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
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("discord_id and guild_id is required"), nil))
		return
	}

	res, code, err := h.entities.GetUserCurrentGMStreak(discordID, guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"discordId": discordID, "guildID": guildID}).Error(err, "[handler.GetUserCurrentGMStreak] - failed to get user current gm streak")
		c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
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
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("discord_id is required"), nil))
		return
	}

	res, code, err := h.entities.GetUserCurrentUpvoteStreak(discordID)
	if err != nil {
		h.log.Fields(logger.Fields{"discordId": discordID}).Error(err, "[handler.GetUserCurrentUpvoteStreak] - failed to get user current upvote streak")
		c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
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
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
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
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(du, nil, nil, nil))
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
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	userID := c.Query("user_id")
	if userID == "" {
		h.log.Info("[handler.GetTopUsers] - user id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("user_id is required"), nil))
		return
	}

	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = "0"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		h.log.Fields(logger.Fields{"page": pageStr}).Error(err, "[handler.GetTopUsers] - invalid page")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("page must be an integer"), nil))
		return
	}

	limitStr := c.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		h.log.Fields(logger.Fields{"limit": limit}).Error(err, "[handler.GetTopUsers] - invalid limit")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("limit must be an integer"), nil))
		return
	}

	data, err := h.entities.GetTopUsers(guildID, userID, limit, page)
	if err != nil {
		h.log.Fields(logger.Fields{"page": pageStr, "limit": limit, "guildID": guildID, "userID": userID}).Error(err, "[handler.GetTopUsers] - failed to get top users")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
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
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	userID := c.Query("user_id")
	if userID == "" {
		h.log.Info("[handler.GetUserProfile] - user id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("user_id is required"), nil))
		return
	}

	data, err := h.entities.GetUserProfile(guildID, userID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "userID": userID}).Error(err, "[handler.GetUserProfile] - failed to get user profile")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// GetUserDevice     godoc
// @Summary     Get user current device data
// @Description Get user current device data
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       device_id query     string true "Device ID"
// @Success     200 {object} response.UserDeviceResponse
// @Router      /users/device [get]
func (h *Handler) GetUserDevice(c *gin.Context) {
	deviceID := c.Query("device_id")
	if deviceID == "" {
		h.log.Info("[handler.GetUserDevice] - discord id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("user_id is required"), nil))
		return
	}

	data, err := h.entities.GetUserDevice(deviceID)
	if err != nil {
		h.log.Fields(logger.Fields{"deviceID": deviceID}).Error(err, "[handler.GetUserDevice] - failed to get user device")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// UpsertUserDevice     godoc
// @Summary     Upsert user current device data
// @Description Upsert user current device data
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertUserDeviceRequest true "Upsert user device"
// @Success     200 {object} response.ResponseMessage
// @Router      /users/device [post]
func (h *Handler) UpsertUserDevice(c *gin.Context) {
	req := request.UpsertUserDeviceRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpsertUserDevice] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("failed to read JSON"), nil))
		return
	}

	err = h.entities.UpsertUserDevice(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpsertUserDevice] - failed to upsert user device")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// DeleteUserDevice     godoc
// @Summary     Delete user current device data
// @Description Delete user current device data
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       Request  body request.DeleteUserDeviceRequest true "Delete user device"
// @Success     200 {object} response.ResponseMessage
// @Router      /users/device [delete]
func (h *Handler) DeleteUserDevice(c *gin.Context) {
	req := request.DeleteUserDeviceRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.DeleteUserDevice] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("failed to read JSON"), nil))
		return
	}

	err = h.entities.DeleteUserDevice(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.DeleteUserDevice] - failed to upsert user device")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetUserWalletByGuildIDAddress     godoc
// @Summary     Get user by guild_id address
// @Description Get user by guild_id address
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       address path     string true "Address"
// @Success     200 {object} response.GetUserWalletByGuildIDAddressResponse
// @Router      /users/{user_id} [get]
func (h *Handler) GetUserWalletByGuildIDAddress(c *gin.Context) {
	address := c.Param("address")
	if address == "" {
		h.log.Info("[handler.GetUserWalletByGuildIDAddress] - address id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("address is required"), nil))
		return
	}
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetUserWalletByGuildIDAddress] - guild_id id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	uw, err := h.entities.GetUserWalletByGuildIDAddress(guildID, address)
	if err != nil {
		if err == entities.ErrRecordNotFound {
			h.log.Fields(logger.Fields{"address": address}).Error(err, "[handler.GetUserWalletByGuildIDAddress] - users not found")
			c.JSON(http.StatusOK, response.CreateResponse[any](nil, nil, nil, nil))
			return
		}
		h.log.Fields(logger.Fields{"address": address}).Error(err, "[handler.GetUserWalletByGuildIDAddress] - failed to get user")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return

	}
	c.JSON(http.StatusOK, response.CreateResponse(uw, nil, nil, nil))
}

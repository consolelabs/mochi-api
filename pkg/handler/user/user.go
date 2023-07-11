package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	errs "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

type Handler struct {
	entities *entities.Entity
	log      logger.Logger
}

func New(entities *entities.Entity, logger logger.Logger) IHandler {
	return &Handler{
		entities: entities,
		log:      logger,
	}
}

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
			c.JSON(http.StatusNotFound, response.CreateResponse[any](nil, nil, nil, nil))
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

// GetTopUsers     godoc
// @Summary     Get top users
// @Description Get top users
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       req query     request.GetTopUsersRequest true "query"
// @Success     200 {object} response.TopUser
// @Router      /users/top [get]
func (h *Handler) GetTopUsers(c *gin.Context) {
	var req request.GetTopUsersRequest
	if err := c.BindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetTopUsers] BindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	data, err := h.entities.GetTopUsers(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.GetTopUsers] entity.GetTopUsers() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// GetUserProfile     godoc
// @Summary     Get user profile
// @Description Get user profile
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       req query     request.GetUserProfileRequest true "query"
// @Success     200 {object} response.GetDataUserProfileResponse
// @Router      /users/profiles/ [get]
func (h *Handler) GetUserProfile(c *gin.Context) {
	var req request.GetUserProfileRequest
	if err := c.BindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetUserProfile] BindQuery() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	data, err := h.entities.GetUserProfile(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.GetUserProfile] - failed to get user profile")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// GetInvites     godoc
// @Summary     Get invites
// @Description Get invites
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       member_id query     string true "Member ID"
// @Param       guild_id query     string true "Guild ID"
// @Router      /community/invites/ [get]
func (h *Handler) GetInvites(c *gin.Context) {
	//TODO: add test
	memberID := c.Query("member_id")
	guildID := c.Query("guild_id")

	invites, err := h.entities.GetUserGlobalInviteCodes(guildID, memberID)
	if err != nil {
		h.log.Fields(logger.Fields{"memberID": memberID, "guildID": guildID}).Error(err, "[handler.GetInvites] - failed to get user global invite code")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(invites, nil, nil, nil))
}

// SendUserXP     godoc
// @Summary     Send User XP
// @Description Send User XP
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       Request  body request.SendUserXPRequest true "Send user XP request"
// @Success     200 {object} response.ResponseMessage
// @Router      /users/xp [post]
func (h *Handler) SendUserXP(c *gin.Context) {
	var req request.SendUserXPRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.SendUserXP] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.SendUserXP(req)
	if err != nil {
		h.log.Error(err, "[handler.SendUserXP] - failed to send user XP")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

func (h *Handler) GetUserBalance(c *gin.Context) {
	userID := c.Param("id")
	balance, err := h.entities.GetUserBalance(userID)
	if err != nil {
		h.log.Fields(logger.Fields{"userID": userID}).Error(err, "[handler.GetUserBalance] - entities.GetUserBalance failed")
		c.JSON(errs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(balance, nil, nil, nil))
}

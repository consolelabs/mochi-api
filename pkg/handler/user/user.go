package user

import (
	"errors"
	"net/http"
	"strconv"

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
// @Param       query query     string false "Query to search by name"
// @Param       sort query     string false "ASC / DESC"
// @Success     200 {object} response.TopUser
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

	query := c.Query("query")

	sort := c.Query("sort")

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

	data, err := h.entities.GetTopUsers(guildID, userID, query, sort, limit, page)
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
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       guild_id query     string true "Guild ID"
// @Param       user_id query     string true "User ID"
// @Success     200 {object} response.GetDataUserProfileResponse
// @Router      /users/profiles/ [get]
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

// GetUserTransaction     godoc
// @Summary     Get user transaction
// @Description Get user transaction
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       id path     string true "user discord ID"
// @Success     200 {object} response.UserTransactionResponse
// @Router      /users/{id}/transactions [get]
func (h *Handler) GetUserTransaction(c *gin.Context) {
	userDiscordId := c.Param("id")
	if userDiscordId == "" {
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("user_discord_id is required"), nil))
		return
	}

	userTransaction, err := h.entities.GetUserTransaction(userDiscordId)
	if err != nil {
		h.log.Error(err, "[handler.GetUserTransaction] - failed to get transaction for user")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(userTransaction, nil, nil, nil))
}

// GetTransactionsByQuery     godoc
// @Summary     Get transactions by query
// @Description Get transactions by query
// @Tags        Tip
// @Accept      json
// @Produce     json
// @Param       sender_id   query  string false  "sender ID"
// @Param       receiver_id   query  string false  "receiver ID"
// @Param       token   query  string true  "token"
// @Success     200 {object} response.TransactionsResponse
// @Router      /tip/transactions [get]
func (h *Handler) GetTransactionsByQuery(c *gin.Context) {
	senderId := c.Query("sender_id")
	receiverId := c.Query("receiver_id")
	if senderId == "" && receiverId == "" {
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("sender_id or receiver_id is required"), nil))
		return
	}
	token := c.Query("token")
	transactions, err := h.entities.GetTransactionsByQuery(senderId, receiverId, token)
	if err != nil {
		h.log.Error(err, "[handler.GetUserTransactionsByQuery] - failed to get transactions")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(transactions, nil, nil, nil))
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

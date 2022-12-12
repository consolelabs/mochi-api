package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

// GetInvites     godoc
// @Summary     Get invites
// @Description Get invites
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       member_id query     string true "Member ID"
// @Param       guild_id query     string true "Guild ID"
// @Success     200 {object} response.GetInvitesResponse
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

// GetInvitesLeaderboard     godoc
// @Summary     Get invites leaderboard
// @Description Get invites leaderboard
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       id path     string true "Guild ID"
// @Success     200 {object} response.GetInvitesLeaderboardResponse
// @Router      /community/invites/leaderboard/{id} [get]
func (h *Handler) GetInvitesLeaderboard(c *gin.Context) {
	guildID := c.Param("id")
	if guildID == "" {
		h.log.Info("[handler.GetInvitesLeaderboard] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	leaderboard, err := h.entities.GetInvitesLeaderboard(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetInvitesLeaderboard] - failed to get invite leaderboards")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(leaderboard, nil, nil, nil))
}

// GetInviteTrackerConfig     godoc
// @Summary     Get invites tracker config
// @Description Get invites tracker config
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       guild_id query     string true "Guild ID"
// @Success     200 {object} response.GetInviteTrackerConfigResponse
// @Router      /community/invites/config [get]
func (h *Handler) GetInviteTrackerConfig(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetInviteTrackerConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	config, err := h.entities.GetInviteTrackerLogChannel(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetInviteTrackerConfig] - failed to get invite tracker log channel")

		code := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// ConfigureInvites     godoc
// @Summary     Configure invites
// @Description Configure invites
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       Request  body request.ConfigureInviteRequest true "Configure Invites request"
// @Success     200 {object} response.ConfigureInvitesResponse
// @Router      /community/invites/config [post]
func (h *Handler) ConfigureInvites(c *gin.Context) {
	var req request.ConfigureInviteRequest
	if err := req.Bind(c); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.ConfigureInvites] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := req.Validate(); err != nil {
		h.log.Fields(logger.Fields{"logChannel": req.LogChannel, "guildID": req.GuildID}).Error(err, "[handler.ConfigureInvites] - failed to validate request")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := h.entities.CreateOrUpdateInviteTrackerLogChannel(req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.ConfigureInvites] - failed to upsert invite tracker log channel")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// InvitesAggregation     godoc
// @Summary     Invites Aggregation
// @Description Invites Aggregation
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       guild_id query     string true "Guild ID"
// @Param       inviter query     string true "Inviter ID"
// @Success     200 {object} response.InvitesAggregationResponse
// @Router      /community/invites/aggregation [get]
func (h *Handler) InvitesAggregation(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.InvitesAggregation] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	inviterID := c.Query("inviter_id")
	if inviterID == "" {
		h.log.Info("[handler.InvitesAggregation] - inviter id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("inviter_id is required"), nil))
		return
	}

	aggregation, err := h.entities.GetUserInvitesAggregation(guildID, inviterID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "inviterID": inviterID}).Error(err, "[handler.InvitesAggregation] - failed to get user invites aggregation")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(aggregation, nil, nil, nil))
}

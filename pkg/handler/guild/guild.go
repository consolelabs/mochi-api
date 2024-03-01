package guild

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	_ "github.com/defipod/mochi/pkg/model"
	baseerrs "github.com/defipod/mochi/pkg/model/errors"
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

// GetGuilds     godoc
// @Summary     Get guilds
// @Description Get guilds
// @Tags        Guild
// @Accept      json
// @Produce     json
// @Success     200 {object} response.GetGuildsResponse
// @Router      /guilds [get]
func (h *Handler) GetGuilds(c *gin.Context) {
	guilds, err := h.entities.GetGuilds()
	if err != nil {
		h.log.Error(err, "[handler.GetGuilds] - failed to get all guilds")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, guilds)
}

// GetGuild      godoc
// @Summary     Get guild
// @Description Get guild
// @Tags        Guild
// @Accept      json
// @Produce     json
// @Param       guild_id path     string true "Guild ID"
// @Success     200      {object} response.GetGuildResponse
// @Router      /guilds/{guild_id} [get]
func (h *Handler) GetGuild(c *gin.Context) {
	guildID := c.Param("guild_id")

	if guildID == "" {
		h.log.Info("[handler.GetGuild] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	guild, err := h.entities.GetGuild(guildID)
	if err != nil {
		if err == entities.ErrRecordNotFound {
			h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuild] - guild not exist")
			c.JSON(http.StatusOK, response.CreateResponse[any](nil, nil, nil, nil))
			return
		}
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuild] - failed to get guild")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, guild)
}

// ListMyGuilds      godoc
// @Summary     Get my guilds list
// @Description Get my guild list
// @Tags        Guild
// @Accept      json
// @Produce     json
// @Param       Authorization header   string true "Authorization"
// @Success     200           {object} response.ListMyGuildsResponse
// @Router      /guilds/user-managed [get]
func (h *Handler) ListMyGuilds(c *gin.Context) {
	accessToken := c.GetString("discord_access_token")

	resp, err := h.entities.ListMyDiscordGuilds(accessToken)
	if err != nil {
		h.log.Fields(logger.Fields{"token": accessToken}).Error(err, "[handler.ListMyGuilds] - failed to list discord guilds")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(resp, nil, nil, nil))
}

// UpdateGuild      godoc
// @Summary     Update guild
// @Description Update guild
// @Tags        Guild
// @Accept      json
// @Produce     json
// @Param       guild_id path string                     true "Guild ID"
// @Param       Request  body request.UpdateGuildRequest true "Update guild request"
// @Success     200        {object} response.ResponseMessage
// @Router      /guilds/{guild_id} [put]
func (h *Handler) UpdateGuild(c *gin.Context) {
	guildID := c.Param("guild_id")
	if guildID == "" {
		h.log.Info("[handler.UpdateGuild] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	var req request.UpdateGuildRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "globalXP": req.GlobalXP, "logChannel": req.LogChannel}).Error(err, "[handler.UpdateGuild] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("global_xp is required"), nil))
		return
	}

	if err := h.entities.UpdateGuild(guildID, req); err != nil {
		if err == baseerrs.ErrRecordNotFound {
			h.log.Fields(logger.Fields{"guildID": guildID, "req": req}).Info("[handler.UpdateGuild] - guild not exist")
			c.JSON(http.StatusOK, response.CreateResponse[any](nil, nil, nil, nil))
			return
		}
		h.log.Fields(logger.Fields{"guildID": guildID, "req": req}).Error(err, "[handler.UpdateGuild] - failed to update guild")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// UpdateGuild      godoc
// @Summary     Update guild
// @Description Update guild
// @Tags        Guild
// @Accept      json
// @Produce     json
// @Param       guild_id path string                     true "Guild ID"
// @Success     200        {object} response.DiscordGuildRoles
// @Router      /guilds/{guild_id}/roles [get]
func (h *Handler) GetGuildRoles(c *gin.Context) {
	guildID := c.Param("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetGuildRoles] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	roles, err := h.entities.GetGuildRoles(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuildRoles] - failed to get guild roles")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(roles, nil, nil, nil))
}

func (h *Handler) ValidateUser(c *gin.Context) {
	var req request.ValidateUserRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.ValidateUser] - failed to read query")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("global_xp is required"), nil))
		return
	}

	ids := strings.Split(req.Ids, ",")

	resp, err := h.entities.ValidateUser(ids, req.GuildId)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.ValidateUser] - failed to validate user")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(resp, nil, nil, nil))
}

func (h *Handler) CreateGuild(c *gin.Context) {
	var req request.CreateGuildRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.ID, "name": req.Name}).Error(err, "[handler.CreateGuild] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := h.entities.CreateGuild(req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.UpdateGuild] - failed to update guild")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

func (h *Handler) GuildReportRoles(c *gin.Context) {
	var req request.GuildReportAdditionalRoleRequest
	if err := c.ShouldBindUri(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.GuildReportAdditionalRoles] - failed to read query")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("invalid request"), nil))
		return
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.GuildReportAdditionalRoles] - failed to read query")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("invalid request"), nil))
		return
	}

	if req.AdditionalRoles == "" {
		resp, err := h.entities.GuildReportRoles(req.GuildId)
		if err != nil {
			h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.GuildReportRoles] - failed to GuildReportRoles")
			c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
			return
		}

		c.JSON(http.StatusOK, response.CreateResponse(resp, nil, nil, nil))
	} else {
		req.AdditionalRolesList = strings.Split(req.AdditionalRoles, ",")

		resp, err := h.entities.GuildReportAdditionalRoles(req)
		if err != nil {
			h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.GuildReportAdditionalRoles] - failed to GuildReportAdditionalRoles")
			c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
			return
		}

		c.JSON(http.StatusOK, response.CreateResponse(resp, nil, nil, nil))
	}
}

func (h *Handler) GuildReportMembers(c *gin.Context) {
	var req request.GuildRequest
	if err := c.ShouldBindUri(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.GuildReportMembers] - failed to read query")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("invalid request"), nil))
		return
	}

	resp, err := h.entities.GuildReportMembers(req.GuildId)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.GuildReportMembers] - failed to GuildReportMembers")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(resp, nil, nil, nil))
}

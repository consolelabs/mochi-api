package handler

import (
	"errors"
	"net/http"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	_ "github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

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

	guild, err := h.entities.GetGuild(guildID)
	if err != nil {
		if err == entities.ErrRecordNotFound {
			h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuild] - guild not exist")
			c.JSON(http.StatusNotFound, response.CreateResponse[any](nil, nil, err, nil))
			return
		}
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuild] - failed to get guild")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, guild)
}

// Createguild      godoc
// @Summary     Create guild
// @Description Create guild
// @Tags        Guild
// @Accept      json
// @Produce     json
// @Param       Request body     request.CreateGuildRequest true "Create Guild request"
// @Success     200     {object} request.CreateGuildRequest
// @Router      /guilds [post]
func (h *Handler) CreateGuild(c *gin.Context) {
	body := request.CreateGuildRequest{}

	if err := c.BindJSON(&body); err != nil {
		h.log.Fields(logger.Fields{"body": body}).Error(err, "[handler.CreateGuild] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.CreateGuild(body); err != nil {
		h.log.Fields(logger.Fields{"body": body}).Error(err, "[handler.CreateGuild] - failed to creat guild")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, body)
}

// GetGuildStats      godoc
// @Summary     Get guild stats
// @Description Get guild stats
// @Tags        Guild
// @Accept      json
// @Produce     json
// @Param       guild_id path     string true "Guild ID"
// @Success     200      {object} model.DiscordGuildStat
// @Router      /guilds/{guild_id}/stats [get]
func (h *Handler) GetGuildStatsHandler(c *gin.Context) {
	guildID := c.Param("guild_id")

	guildStat, err := h.entities.GetByGuildID(guildID)
	if err != nil {
		if err == entities.ErrRecordNotFound {
			h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuildStatsHandler] - guild not exist")
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuildStatsHandler] - failed to get guild")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, guildStat)
}

// CreateGuildChannel      godoc
// @Summary     Create guild channel
// @Description Create guild channel
// @Tags        Guild
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Param       count_type query string false "Guild ID"
// @Success     200      {string} string "ok"
// @Router      /guilds/{guild_id}/channels [post]
func (h *Handler) CreateGuildChannel(c *gin.Context) {
	log := logger.NewLogrusLogger()
	guildID := c.Param("guild_id")
	countType := c.Query("count_type")

	log.Infof("Creating stats channel for counting. GuildId: %v, CountType: %v", guildID, countType)

	err := h.entities.CreateGuildChannel(guildID, countType)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "countType": countType}).Error(err, "[handler.CreateGuildChannel] - failed to create guild channel")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// ListMyGuilds      godoc
// @Summary     Get my guilds list
// @Description Get my guild list
// @Tags        Guild
// @Accept      json
// @Produce     json
// @Param       Authorization header   string true "Authorization"
// @Success     200           {object} entities.ListMyGuildsResponse
// @Router      /guilds/user-managed [get]
func (h *Handler) ListMyGuilds(c *gin.Context) {
	accessToken := c.GetString("discord_access_token")

	resp, err := h.entities.ListMyDiscordGuilds(accessToken)
	if err != nil {
		h.log.Fields(logger.Fields{"token": accessToken}).Error(err, "[handler.ListMyGuilds] - failed to list discord guilds")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
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
		h.log.Fields(logger.Fields{"guildID": guildID, "req": req}).Error(err, "[handler.UpdateGuild] - failed to update guild")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

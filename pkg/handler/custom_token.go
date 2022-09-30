package handler

import (
	"errors"
	"net/http"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

// HandlerGuildCustomTokenConfig     godoc
// @Summary     Guild custom token config
// @Description Guild custom token config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertCustomTokenConfigRequest true "Custom guild custom token config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/custom-tokens [post]
func (h *Handler) HandlerGuildCustomTokenConfig(c *gin.Context) {
	var req request.UpsertCustomTokenConfigRequest

	// handle input validate
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.HandlerGuildCustomTokenConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.HandlerGuildCustomTokenConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	if req.Symbol == "" {
		h.log.Info("[handler.HandlerGuildCustomTokenConfig] - symbol empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("symbol is required"), nil))
		return
	}
	if req.Address == "" {
		h.log.Info("[handler.HandlerGuildCustomTokenConfig] - address empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("address is required"), nil))
		return
	}
	if req.Chain == "" {
		h.log.Info("[handler.HandlerGuildCustomTokenConfig] - chain empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Chain is required"})
		return
	}

	if err := h.entities.CreateCustomToken(req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.HandlerGuildCustomTokenConfig] - fail to create custom token")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// ListAllCustomToken     godoc
// @Summary     List all guild custom token
// @Description List all guild custom token
// @Tags        Guild
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Success     200 {object} response.ListAllCustomTokenResponse
// @Router      /guilds/custom-tokens [get]
func (h *Handler) ListAllCustomToken(c *gin.Context) {
	guildID := c.Param("guild_id")

	// get all token with guildID
	returnToken, err := h.entities.GetAllSupportedToken(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.ListAllCustomToken] - failed to get all tokens")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ListAllCustomTokenResponse{Data: returnToken})
}

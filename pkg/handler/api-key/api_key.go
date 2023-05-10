package apikey

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
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

// GetApiKeyByDiscordId     godoc
// @Summary     Get api key by discordId
// @Description Get api key by discordId
// @Tags        ApiKey
// @Accept      json
// @Produce     json
// @Param       discord_id   path  string true  "Discord ID"
// @Success     200 {object} response.ProfileApiKeyResponse
// @Router      /api-key/{discord_id} [get]
func (h *Handler) GetApiKeyByDiscordId(c *gin.Context) {
	discordId := c.Param("discord_id")
	if discordId == "" {
		h.log.Info("[handler.GetApiKeyByDiscordId] - discordId is empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, nil, nil))
		return
	}
	data, err := h.entities.GetApiKeyByDiscordId(discordId)
	if err != nil {
		h.log.Fields(logger.Fields{"discord_id": discordId}).Error(err, "[handler.GetApiKeyByDiscordId] - fail to get apiKey")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](data, nil, nil, nil))
}

// CreateApiKeyByDiscordId     godoc
// @Summary     Create apiKey by discordId
// @Description Create apiKey by discordId
// @Tags        ApiKey
// @Accept      json
// @Produce     json
// @Param       discord_id   path  string true  "Discord ID"
// @Success     200 {object} response.ProfileApiKeyResponse
// @Router      /api-key/{discord_id} [post]
func (h *Handler) CreateApiKeyByDiscordId(c *gin.Context) {
	discordId := c.Param("discord_id")
	if discordId == "" {
		h.log.Info("[handler.CreateApiKeyByDiscordId] - discordId is empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, nil, nil))
		return
	}
	data, err := h.entities.CreateApiKeyByDiscordId(discordId)
	if err != nil {
		if strings.Contains(err.Error(), "record already existed") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "record already existed"})
			return
		}
		h.log.Fields(logger.Fields{"discord_id": discordId}).Error(err, "[handler.CreateApiKeyByDiscordId] - fail to create apiKey")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse[any](data, nil, nil, nil))
}

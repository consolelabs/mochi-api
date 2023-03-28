package vault

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
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

func (h *Handler) CreateVault(c *gin.Context) {
	var req request.CreateVaultRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildId, "name": req.Name, "threshold": req.Threshold}).Error(err, "[handler.CreateVault] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	vault, err := h.entities.CreateVault(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildId, "name": req.Name, "threshold": req.Threshold}).Error(err, "[handler.CreateVault] - failed to create vault")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](vault, nil, nil, nil))
}
func (h *Handler) GetVault(c *gin.Context) {
	guildId := c.Query("guild_id")
	if guildId == "" {
		h.log.Info("[handler.GetVault] - guildId is empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, nil, nil))
		return
	}

	vault, err := h.entities.GetVault(guildId)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildId}).Error(err, "[handler.GetVault] - failed to get vault")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](vault, nil, nil, nil))
}

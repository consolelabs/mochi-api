package telegram

import (
	"errors"
	"net/http"

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

func (h *Handler) GetByUsername(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("username is required"), nil))
		return
	}
	userTelegram, err := h.entities.GetUserTelegramByUsername(username)
	if err != nil {
		h.log.Fields(logger.Fields{"username": username}).Error(err, "[handler.Telegram.GetByUsername] entity.GetUserTelegramByUsername() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(userTelegram, nil, nil, nil))
}

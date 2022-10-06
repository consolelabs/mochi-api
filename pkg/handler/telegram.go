package handler

import (
	"errors"
	"net/http"

	baseerrs "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	_ "github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

// GetLinkedTelegram     godoc
// @Summary     Get telegram account linked with discord ID
// @Description Get telegram account linked with discord ID
// @Tags        Telegram
// @Accept      json
// @Produce     json
// @Param       telegram_username query string true "request"
// @Success     200 {object} response.GetLinkedTelegramResponse
// @Router      /configs/telegram [get]
func (h *Handler) GetLinkedTelegram(c *gin.Context) {
	telegramUsername := c.Query("telegram_username")
	if telegramUsername == "" {
		h.log.Info("[handler.GetLinkedTelegram] - telegram_username is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "telegram_username is required"})
		return
	}
	res, err := h.entities.GetByTelegramUsername(telegramUsername)
	if err != nil {
		h.log.Error(err, "[handler.GetLinkedTelegram] entity.GetByTelegramUsername() failed")
		code := http.StatusInternalServerError
		if err == baseerrs.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// LinkUserTelegramWithDiscord     godoc
// @Summary     Link user's telegram with discord account
// @Description Link user's telegram with discord account
// @Tags        Telegram
// @Accept      json
// @Produce     json
// @Param       req body request.LinkUserTelegramWithDiscordRequest true "request"
// @Success     201 {object} response.LinkUserTelegramWithDiscordResponse
// @Router      /configs/telegram [post]
func (h *Handler) LinkUserTelegramWithDiscord(c *gin.Context) {
	var req request.LinkUserTelegramWithDiscordRequest
	if err := c.Bind(&req); err != nil {
		h.log.Error(err, "[handler.LinkUserTelegramWithDiscord] Bind() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("discord_id and telegramusernam are required"), nil))
		return
	}
	res, err := h.entities.LinkUserTelegramWithDiscord(req)
	if err != nil {
		h.log.Error(err, "[handler.LinkUserTelegramWithDiscord] entity.LinkUserTelegramWithDiscord() failed")
		c.JSON(baseerrs.GetStatusCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

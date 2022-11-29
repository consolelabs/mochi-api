package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

// CreateConfigNotify   godoc
// @Summary     OffChain Tip Bot - Config notify
// @Description API config notify channel for token
// @Tags        OffChain
// @Accept      json
// @Produce     json
// @Param       Request body request.CreateTipConfigNotify true "config notify request"
// @Success     200 {object} response.ResponseMessage
// @Router      /offchain-tip-bot/config-notify [post]
func (h *Handler) CreateConfigNotify(c *gin.Context) {
	req := request.CreateTipConfigNotify{}

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateConfigNotify] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.CreateConfigNotify(req)
	if err != nil {
		h.log.Error(err, "[handler.CreateConfigNotify] - failed to create config notify")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// ListConfigNotify   godoc
// @Summary     OffChain Tip Bot - Config notify
// @Description API get list config notify channel for token
// @Tags        OffChain
// @Accept      json
// @Produce     json
// @Param       guild_id query string true "guild id"
// @Success     200 {object} response.ListConfigNotifyResponse
// @Router      /offchain-tip-bot/config-notify [get]
func (h *Handler) ListConfigNotify(c *gin.Context) {
	guildId := c.Query("guild_id")
	listConfigs, err := h.entities.ListConfigNotify(guildId)
	if err != nil {
		h.log.Error(err, "[handler.ListConfigNotify] - failed to list config notify")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(listConfigs, nil, nil, nil))
}

// DeleteConfigNotify   godoc
// @Summary     OffChain Tip Bot - Config notify
// @Description API delete config notify channel for token
// @Tags        OffChain
// @Accept      json
// @Produce     json
// @Param       id path string true "id of config notify"
// @Success     200 {object} response.ResponseMessage
// @Router      /offchain-tip-bot/config-notify/{id} [delete]
func (h *Handler) DeleteConfigNotify(c *gin.Context) {
	id := c.Param("id")

	err := h.entities.DeleteConfigNotify(id)
	if err != nil {
		h.log.Error(err, "[handler.DeleteConfigNotify] - failed to delete config notify")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

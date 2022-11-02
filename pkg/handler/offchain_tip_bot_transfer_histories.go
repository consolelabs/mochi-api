package handler

import (
	"errors"
	"net/http"

	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

// GetUserBalances     godoc
// @Summary     Get user transaction
// @Description Get user transaction
// @Tags        OffChain
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

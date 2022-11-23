package handler

import (
	"errors"
	"net/http"

	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

// GetUserTransaction     godoc
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

// GetTransactionsByQuery     godoc
// @Summary     Get transactions by query
// @Description Get transactions by query
// @Tags        OffChain
// @Accept      json
// @Produce     json
// @Param       sender_id   query  string false  "sender ID"
// @Param       receiver_id   query  string false  "receiver ID"
// @Param       token   query  string true  "token"
// @Success     200 {object} response.TransactionsResponse
// @Router      /offchain-tip-bot/transactions [get]
func (h *Handler) GetTransactionsByQuery(c *gin.Context) {
	senderId := c.Query("sender_id")
	receiverId := c.Query("receiver_id")
	if senderId == "" && receiverId == "" {
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("sender_id or receiver_id is required"), nil))
		return
	}
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("token is required"), nil))
		return
	}
	transactions, err := h.entities.GetTransactionsByQuery(senderId, receiverId, token)
	if err != nil {
		h.log.Error(err, "[handler.GetUserTransactionsByQuery] - failed to get transactions")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(transactions, nil, nil, nil))
}

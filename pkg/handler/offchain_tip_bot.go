package handler

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	offchaintipbotchain "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_chain"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func (h *Handler) OffchainTipBotListAllChains(c *gin.Context) {
	tokenID := c.Query("token_id")
	tokenSymbol := c.Query("token_symbol")
	returnChain, err := h.entities.OffchainTipBotListAllChains(
		offchaintipbotchain.Filter{
			TokenID:     tokenID,
			TokenSymbol: tokenSymbol,
		},
	)
	if err != nil {
		h.log.Error(err, "[handler.OffchainTipBotListAllChains] - failed to list chains")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(returnChain, nil, nil, nil))
}

// AddContract   godoc
// @Summary     OffChain Tip Bot - Create an assign contract
// @Description Create an assign contract when user want to deposit a specific token to contract
// @Tags        OffChain, TipBot, Deposit
// @Accept      json
// @Produce     json
// @Param       Request  body request.CreateAssignContract true "Create assign contract request"
// @Success     200 {object} response.GetAssignedContract
// @Router      /offchain-tip-bot/assign-contract [post]
func (h *Handler) OffchainTipBotCreateAssignContract(c *gin.Context) {
	body := request.CreateAssignContract{}

	if err := c.BindJSON(&body); err != nil {
		h.log.Fields(logger.Fields{"body": body}).Error(err, "[handler.CreateDefaultRole] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	chains, err := h.entities.OffchainTipBotListAllChains(
		offchaintipbotchain.Filter{
			TokenSymbol:         body.TokenSymbol,
			IsContractAvailable: true,
		},
	)
	if err != nil {
		h.log.Error(err, "[handler.OffchainTipBotCreateAssignContract] - failed to list chains")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if len(chains) == 0 {
		err := errors.New("contract not found or already assigned")
		h.log.Error(err, "[handler.OffchainTipBotCreateAssignContract] - contract not found or already assigned")
		c.JSON(http.StatusNotFound, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	ac := &model.OffchainTipBotAssignContract{
		ContractID:  chains[0].Contracts[0].ID.String(),
		ChainID:     chains[0].ID.String(),
		UserID:      body.UserID,
		ExpiredTime: time.Now().Add(3 * 24 * time.Hour),
	}
	for _, t := range chains[0].Tokens {
		if strings.EqualFold(strings.ToLower(t.TokenSymbol), strings.ToLower(body.TokenSymbol)) {
			ac.TokenID = t.ID.String()
			break
		}
	}

	userAssignedContract, err := h.entities.OffchainTipBotCreateAssignContract(ac)
	if err != nil {
		h.log.Error(err, "[handler.OffchainTipBotCreateAssignContract] - failed to create assign contract")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(userAssignedContract, nil, nil, nil))
}

// GetUserBalances     godoc
// @Summary     Get offchain user bals
// @Description Get offchain user bals
// @Tags        OffChain
// @Accept      json
// @Produce     json
// @Param       user_id query     string true "user ID"
// @Success     200 {object} response.GetUserBalancesResponse
// @Router      /offchain-tip-bot/balances [get]
func (h *Handler) GetUserBalances(c *gin.Context) {
	userID := c.Query("user_id")

	if userID == "" {
		h.log.Info("[handler.GetUserBalances] - missing user id")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("user_id is required"), nil))
		return
	}

	userBalances, err := h.entities.GetUserBalances(userID)
	if err != nil {
		h.log.Error(err, "[handler.GetUserBalances] - failed to get user balances")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(userBalances, nil, nil, nil))
}

// OffchainTipBotWithdraw     godoc
// @Summary     OffChain Tip Bot - Withdraw
// @Description OffChain Tip Bot - Withdraw
// @Tags        OffChain
// @Accept      json
// @Produce     json
// @Success     200 {object} response.OffchainTipBotWithdrawResponse
// @Router      /offchain-tip-bot/withdraw [post]
func (h *Handler) OffchainTipBotWithdraw(c *gin.Context) {
	resp := response.OffchainTipBotWithdrawResponse{
		FromDiscordID:  "393034938028392449",
		ToAddress:      "0x140dd183e18ba39bd9BE82286ea2d96fdC48117A",
		Amount:         3.1,
		Cryptocurrency: "usdt",
		TxHash:         "0x940518031ecb2c29e6c7b6f73aa7574a8d12b897f46cca9b5d67ea111012cd5e",
		TxUrl:          "https://etherscan.io/tx/0x940518031ecb2c29e6c7b6f73aa7574a8d12b897f46cca9b5d67ea111012cd5e",
		WithdrawAmount: 1.4,
		TransactionFee: 0.0001,
	}
	c.JSON(http.StatusOK, response.CreateResponse(resp, nil, nil, nil))
}

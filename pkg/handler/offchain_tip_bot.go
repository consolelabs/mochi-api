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
// @Tags        OffChain
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
// @Param       Request  body request.OffchainWithdrawRequest true "Withdraw token request"
// @Success     200 {object} response.OffchainTipBotWithdrawResponse
// @Router      /offchain-tip-bot/withdraw [post]
func (h *Handler) OffchainTipBotWithdraw(c *gin.Context) {
	req := request.OffchainWithdrawRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.OffchainTipBotWithdraw] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	res, err := h.entities.OffchainTipBotWithdraw(req)
	if err != nil {
		if strings.Contains(err.Error(), "Token not supported") || strings.Contains(err.Error(), "Not enough balance") {
			c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
			return
		}
		h.log.Error(err, "[handler.OffchainTipBotWithdraw] - failed to withdraw")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

// TransferToken   godoc
// @Summary     OffChain Tip Bot - Transfer token
// @Description API transfer token for tip, airdrop, ...
// @Tags        OffChain
// @Accept      json
// @Produce     json
// @Param       Request  body request.OffchainTransferRequest true "Transfer token request"
// @Success     200 {object} response.OffchainTipBotTransferTokenResponse
// @Router      /offchain-tip-bot/transfer [post]
func (h *Handler) TransferToken(c *gin.Context) {
	req := request.OffchainTransferRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.TransferToken] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	transferHistories, err := h.entities.TransferToken(req)
	if err != nil {
		if strings.Contains(err.Error(), "Token not supported") || strings.Contains(err.Error(), "Not enough balance") {
			c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
			return
		}
		h.log.Error(err, "[entities.TransferToken] - failed to transfer token")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(transferHistories, nil, nil, nil))
}

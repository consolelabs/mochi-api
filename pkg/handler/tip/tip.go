package tip

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

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

// TransferToken   godoc
// @Summary     OffChain Tip Bot - Transfer token
// @Description API transfer token for tip, airdrop, ...
// @Tags        Tip
// @Accept      json
// @Produce     json
// @Param       Request  body request.OffchainTransferRequest true "Transfer token request"
// @Success     200 {object} response.OffchainTipBotTransferTokenResponse
// @Router      /tip/transfer [post]
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

// SubmitOnchainTransfer   godoc
// @Summary     Onchain Tip Bot - Submit transfer transaction
// @Description Onchain Tip Bot - Submit transfer transaction
// @Tags        Tip
// @Accept      json
// @Produce     json
// @Param       Request  body request.SubmitOnchainTransferRequest true "req"
// @Success     200 {object} response.SubmitOnchainTransferResponse
// @Router      /tip/onchain/submit [post]
func (h *Handler) SubmitOnchainTransfer(c *gin.Context) {
	req := request.SubmitOnchainTransferRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.SubmitOnchainTransfer] ShouldBindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	data, err := h.entities.SubmitOnchainTransfer(req)
	if err != nil {
		if strings.Contains(err.Error(), "Token not supported") || strings.Contains(err.Error(), "Not enough balance") {
			c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
			return
		}
		h.log.Error(err, "[handler.SubmitOnchainTransfer] entity.SubmitOnchainTransfer() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// ClaimOnchainTransfer   godoc
// @Summary     Onchain Tip Bot - Submit transfer transaction
// @Description Onchain Tip Bot - Submit transfer transaction
// @Tags        Tip
// @Accept      json
// @Produce     json
// @Param       Request  body request.ClaimOnchainTransferRequest true "req"
// @Success     200 {object} response.ClaimOnchainTransferResponse
// @Router      /tip/onchain/claim [post]
func (h *Handler) ClaimOnchainTransfer(c *gin.Context) {
	req := request.ClaimOnchainTransferRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.ClaimOnchainTransfer] ShouldBindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	data, err := h.entities.ClaimOnchainTransfer(req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, response.CreateResponse[any](nil, nil, err, nil))
			return
		}
		h.log.Error(err, "[handler.ClaimOnchainTransfer] entity.ClaimOnchainTransfer() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// GetOnchainTransfers   godoc
// @Summary     Onchain Tip Bot - Get user's onchain transfers
// @Description Onchain Tip Bot - Get user's onchain transfers
// @Tags        Tip
// @Accept      json
// @Produce     json
// @Param       user_id  query string false "userId"
// @Param       status  query string false "status"
// @Success     200 {object} response.GetOnchainTransfersResponse
// @Router      /tip/onchain/{user_id}/transfers [get]
func (h *Handler) GetOnchainTransfers(c *gin.Context) {
	userId := c.Param("user_id")
	if userId == "" {
		err := errors.New("user_id is required")
		h.log.Errorf(err, "[handler.GetOnchainTransfers] %s", err.Error())
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	status := c.Query("status")
	data, err := h.entities.GetUserOnchainTransfers(userId, status)
	if err != nil {
		h.log.Fields(logger.Fields{"user_id": userId, "status": status}).Error(err, "[handler.GetOnchainTransfers] entity.GetUserOnchainTransfers() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// GetOnchainBalances   godoc
// @Summary     Onchain Tip Bot - Get user's onchain balances
// @Description Onchain Tip Bot - Get user's onchain balances
// @Tags        Tip
// @Accept      json
// @Produce     json
// @Param       user_id  path string true "userId"
// @Success     200 {object} response.GetUserBalancesResponse
// @Router      /tip/onchain/{user_id}/balances [get]
func (h *Handler) GetOnchainBalances(c *gin.Context) {
	userId := c.Param("user_id")
	if userId == "" {
		err := errors.New("user_id is required")
		h.log.Errorf(err, "[handler.GetOnchainBalances] %s", err.Error())
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	data, err := h.entities.GetPendingOnchainBalances(userId)
	if err != nil {
		h.log.Fields(logger.Fields{"user_id": userId}).Error(err, "[handler.GetOnchainBalances] entity.GetPendingOnchainBalances() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

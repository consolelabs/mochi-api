package tip

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
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

// TransferTokenV2   godoc
// @Summary     OffChain Tip Bot - Transfer token v2
// @Description API transfer token for tip, airdrop, ...
// @Tags        Tip
// @Accept      json
// @Produce     json
// @Param       Request  body request.TransferV2Request true "Transfer token request"
// @Success     200 {object} response.TransferTokenV2Response
// @Router      /tip/transfer-v2 [post]
func (h *Handler) TransferTokenV2(c *gin.Context) {
	req := request.TransferV2Request{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.TransferToken] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if !util.ValidateNumberSeries(req.Sender) {
		err := errors.New("sender is invalid")
		h.log.Error(err, "[handler.TransferTokenV2] validate sender failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	ctxProfileId := c.GetString("profile_id")
	isMochi := c.GetBool("is_mochi")
	if !isMochi && ctxProfileId != req.Sender {
		err := errors.New("you can only transfer tokens from your profile")
		h.log.Error(err, "[handler.TransferTokenV2] unauthorized sender")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	for _, recipient := range req.Recipients {
		if !util.ValidateNumberSeries(recipient) {
			err := errors.New("recipient is invalid")
			h.log.Error(err, "[handler.TransferTokenV2] validate recipient failed")
			c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
			return
		}
	}

	transferHistories, err := h.entities.TransferTokenV2(req)
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

package handler

import (
	"errors"
	"net/http"

	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

// CreateTradeOffer     godoc
// @Summary     Create Trade Offer
// @Description Create Trade Offer
// @Tags        Trade
// @Accept      json
// @Produce     json
// @Param       Request  body request.CreateTradeOfferRequest true "Create Trade Offer Request"
// @Success     200 {object} response.CreateTradeOfferResponse
// @Router      /trades [post]
func (h *Handler) CreateTradeOffer(c *gin.Context) {
	req := request.CreateTradeOfferRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err, "[handler.CreateTradeOffer] c.ShouldBindJSON failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	data, err := h.entities.CreateTradeOffer(req)
	if err != nil {
		h.log.Error(err, "[handler.CreateTradeOffer] entities.CreateTradeOffer failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusCreated, response.CreateResponse(data, nil, nil, nil))
}

// GetTradeOffer     godoc
// @Summary     Get Trade Offer
// @Description Get Trade Offer
// @Tags        Trade
// @Accept      json
// @Produce     json
// @Param       id path  string true  "Trade Offer ID"
// @Success     200 {object} response.GetTradeOfferResponse
// @Router      /trades/{id} [get]
func (h *Handler) GetTradeOffer(c *gin.Context) {
	tradeId := c.Param("id")
	if tradeId == "" {
		h.log.Info("[handler.GetTradeOffer] - trade id missing")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("id is required"), nil))
		return
	}
	data, err := h.entities.GetTradeOffer(tradeId)
	if err != nil {
		h.log.Error(err, "[handler.GetTradeOffer] entities.GetTradeOffer failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

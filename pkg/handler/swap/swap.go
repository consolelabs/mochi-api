package swap

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	baseerrs "github.com/defipod/mochi/pkg/model/errors"
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

// GetSwapRoutes     godoc
// @Summary     Get swap route for token pairs
// @Description Get swap route for token pairs
// @Tags        Swap
// @Accept      json
// @Produce     json
// @Param       from   query  string true  "from token symbol"
// @Param       to   query  string true  "to token symbol"
// @Param       amount   query  string true  "from amount value"
// @Param       chain_name   query  string false  "chain name"
// @Param       chain_id   query  string false  "chain id"
// @Success     200 {object} response.SwapRouteResponseData
// @Router      /swap/route [get]
func (h *Handler) GetSwapRoutes(c *gin.Context) {
	chainId := 0
	if c.Query("chain_id") != "" {
		chainId, _ = strconv.Atoi(c.Query("chain_id"))
	}

	req := request.GetSwapRouteRequest{
		From:        c.Query("from"),
		To:          c.Query("to"),
		Amount:      c.Query("amount"),
		ChainId:     int64(chainId),
		ChainName:   c.Query("chain_name"),
		FromTokenId: c.Query("from_token_id"),
		ToTokenId:   c.Query("to_token_id"),
	}

	data, err := h.entities.GetSwapRoutes(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"chainId": req.ChainId, "chainName": req.ChainName, "from": req.From, "to": req.To, "amount": req.Amount}).Error(err, "[handler.GetSwapRoutes] - cannot get data from kyber")
		c.JSON(baseerrs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](data, nil, nil, nil))
}

// ExecuteSwapRoutes     godoc
// @Summary     Execute swap token
// @Description Execute swap token
// @Tags        Swap
// @Accept      json
// @Produce     json
// @Param       Request  body request.SwapRequest true "swap request"
// @Success     200 {object} response.ResponseMessage
// @Router      /swap [post]
func (h *Handler) ExecuteSwapRoutes(c *gin.Context) {
	var req request.SwapRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.ExecuteSwapRoutes] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	_, err := h.entities.Swap(req)
	if err != nil {
		if strings.Contains(err.Error(), "insufficient balance") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
			return
		}
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.BuildSwapRoutes] - failed to build swap route")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

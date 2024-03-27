package dexes

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

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

// SearchDexPair godoc
// @Summary     Search dex pair
// @Description Search dex pair
// @Tags        Dex
// @Accept      json
// @Produce     json
// @Param       query   query  string true  "dex query"
// @Success     200 {object} response.SearchDexPairResponse
// @Router      /dexes/search [get]
func (h *Handler) SearchDexPair(c *gin.Context) {
	req := request.SearchDexPairRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.SearchPair] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	res, err := h.entities.SearchDexPair(req)
	if err != nil {
		h.log.Error(err, "[handler.SearchPair] entity.SearchPair() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

func (h *Handler) SearchDexScreenerPair(c *gin.Context) {
	req := request.SearchDexScreenerPairRequest{}
	if err := c.BindQuery(&req); err != nil {
		h.log.Error(err, "[handler.SearchDexScreenerPair] BindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	req.TokenAddress = strings.TrimSpace(req.TokenAddress)
	req.Symbol = strings.TrimSpace(req.Symbol)
	if req.TokenAddress == "" && req.Symbol == "" {
		err := errors.New("either symbol or token_address is required")
		h.log.Error(err, "[handler.SearchDexScreenerPair] required field(s) are missing")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	res, err := h.entities.SearchDexScreenerPair(req)
	if err != nil {
		h.log.Error(err, "[handler.SearchDexScreenerPair] entity.SearchDexScreenerPair() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

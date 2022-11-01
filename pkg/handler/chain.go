package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/response"
)

// AddContract   godoc
// @Summary     List All Chain
// @Description List All Chain
// @Tags        Chain
// @Accept      json
// @Produce     json
// @Success     200 {object} response.GetListAllChainsResponse
// @Router      /chains [get]
func (h *Handler) ListAllChain(c *gin.Context) {
	returnChain, err := h.entities.ListAllChain()
	if err != nil {
		h.log.Error(err, "[handler.ListAllChain] - failed to list all chains")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(returnChain, nil, nil, nil))
}

func (h *Handler) Test(c *gin.Context) {

	c.JSON(http.StatusOK, "ok")
}

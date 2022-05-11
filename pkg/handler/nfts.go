package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) GetNFTsDetailByParam(c *gin.Context) {

	address := strings.ToLower(c.Query("address"))
	tokenId := c.Query("token_id")
	chain := c.Query("chain")

	data, err := h.entities.GetUserNFTS(address, tokenId, chain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

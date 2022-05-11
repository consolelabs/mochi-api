package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetNFTDetail(c *gin.Context) {
	symbol := c.Param("symbol")
	tokenId := c.Param("id")

	data, err := h.entities.GetNFTDetail(symbol, tokenId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

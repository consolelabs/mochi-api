package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetNftSalesHandler(c *gin.Context) {
	data, err := h.entities.GetNftSales()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, data)
}

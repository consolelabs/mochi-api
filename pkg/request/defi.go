package request

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetMarketChartRequest struct {
	CoinID   string `json:"coin_id"`
	Currency string `json:"currency"`
	Days     int    `json:"days"`
}

func ValidateRequest(c *gin.Context) (*GetMarketChartRequest, error) {
	params := c.Request.URL.Query()

	days, err := strconv.Atoi(params.Get("days"))
	if err != nil {
		return nil, fmt.Errorf("invalid days")
	}
	req := &GetMarketChartRequest{
		CoinID:   params.Get("coin_id"),
		Currency: params.Get("currency"),
		Days:     days,
	}

	if req.CoinID == "" {
		return nil, fmt.Errorf("coin_id is required")
	}
	if req.Currency == "" {
		req.Currency = "usd"
	}

	return req, nil
}

package request

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetMarketChartRequest struct {
	CoinID    string `json:"coin_id"`
	Currency  string `json:"currency"`
	Days      int    `json:"days"`
	DiscordID string `json:"discord_id"`
}

type TransferRequest struct {
	Sender         string   `json:"sender"`
	Recipients     []string `json:"recipients"`
	GuildID        string   `json:"guildId"`
	ChannelID      string   `json:"channelId"`
	Amount         float64  `json:"amount"`
	Cryptocurrency string   `json:"cryptocurrency"`
	Each           bool     `json:"each"`
	All            bool     `json:"all"`
	TransferType   string   `json:"transferType"`
}

func ValidateRequest(c *gin.Context) (*GetMarketChartRequest, error) {
	params := c.Request.URL.Query()

	days, err := strconv.Atoi(params.Get("days"))
	if err != nil {
		return nil, fmt.Errorf("invalid days")
	}
	req := &GetMarketChartRequest{
		CoinID:    params.Get("coin_id"),
		Currency:  params.Get("currency"),
		DiscordID: params.Get("discord_id"),
		Days:      days,
	}

	if req.CoinID == "" {
		return nil, fmt.Errorf("coin_id is required")
	}
	if req.Currency == "" {
		req.Currency = "usd"
	}

	return req, nil
}

func (input *TransferRequest) Bind(c *gin.Context) (err error) {
	err = c.BindJSON(input)
	if err != nil {
		return err
	}

	return err
}

type GetUserWatchlistRequest struct {
	UserID string `json:"user_id" form:"user_id" binding:"required"`
	Page   int    `json:"page" form:"page"`
	Size   int    `json:"size" form:"size"`
}

type AddToWatchlistRequest struct {
	UserID      string `json:"user_id"`
	Symbol      string `json:"symbol"`
	CoinGeckoID string `json:"coin_gecko_id"`
}

type RemoveFromWatchlistRequest struct {
	UserID string `json:"user_id" form:"user_id" binding:"required"`
	Symbol string `json:"symbol" form:"symbol" binding:"required"`
}

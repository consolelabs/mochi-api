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

type TransferRequest struct {
	FromDiscordID  string   `json:"fromDiscordId"`
	ToDiscordIDs   []string `json:"toDiscordIds"`
	GuildID        string   `json:"guildId"`
	ChannelID      string   `json:"channelId"`
	Amount         float64  `json:"amount"`
	Cryptocurrency string   `json:"cryptocurrency"`
	Each           bool     `json:"each"`
	All            bool     `json:"all"`
}

type WithdrawRequest struct {
	FromDiscordID  string  `json:"fromDiscordId"`
	ToAddress      string  `json:"toAddress"`
	GuildID        string  `json:"guildId"`
	ChannelID      string  `json:"channelId"`
	Amount         float64 `json:"amount"`
	Cryptocurrency string  `json:"cryptocurrency"`
	All            bool    `json:"all"`
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

func (input *TransferRequest) Bind(c *gin.Context) (err error) {
	err = c.BindJSON(input)
	if err != nil {
		return err
	}

	return err
}

func (input *WithdrawRequest) Bind(c *gin.Context) (err error) {
	err = c.BindJSON(input)
	if err != nil {
		return err
	}

	return err
}

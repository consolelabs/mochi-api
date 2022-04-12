package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type TokenType string

var (
	SingleTokenType TokenType = "single"
	LPTokenType     TokenType = "lp"
)

type Token struct {
	ID                  int             `json:"id"`
	Address             string          `json:"address"`
	Symbol              string          `json:"symbol"`
	ChainID             int             `json:"chain_id"`
	Decimals            int             `json:"decimal"`
	DiscordBotSupported bool            `json:"discord_bot_supported"`
	Icon                JSONArrayString `json:"icon"`
	Type                TokenType       `json:"type"`
	Components          TokenComponents `json:"components,omitempty"`
}

type TokenComponent struct {
	*Token           `json:"token,omitempty" gorm:"-"`
	TokenAddress     string    `json:"token_address"`
	AmountPerLpToken *BigFloat `json:"amount_per_lp_token"`
}

type TokenComponents []*TokenComponent

func (c TokenComponents) Value() (driver.Value, error) {
	if c == nil || len(c) == 0 {
		return "[]", nil
	}

	b, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return string(b), nil
}

func (c *TokenComponents) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	switch t := value.(type) {
	case []uint8:
		jsonData := value.([]uint8)

		if string(jsonData) == "null" || string(jsonData) == "[]" {
			return nil
		}

		if err := json.Unmarshal([]byte(jsonData), c); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("could not scan type %T into json", t)
	}
}

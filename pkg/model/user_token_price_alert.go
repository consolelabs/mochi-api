package model

import "github.com/defipod/mochi/pkg/model/errors"

type AlertType string

const (
	PriceReaches    AlertType = "price_reaches"
	PriceRisesAbove AlertType = "price_rises_above"
	PriceDropsTo    AlertType = "price_drops_to"
	ChangeIsOver    AlertType = "change_is_over"
	ChangeIsUnder   AlertType = "change_is_under"
)

type AlertFrequency string

const (
	OnlyOnce AlertFrequency = "only_once"
	OnceADay AlertFrequency = "once_a_day"
	Always   AlertFrequency = "always"
)

type UserTokenPriceAlert struct {
	UserID      string         `json:"user_id"`
	CoinGeckoID string         `json:"coin_gecko_id"`
	AlertType   AlertType      `json:"alert_type"`
	Frequency   AlertFrequency `json:"frequency"`
	Price       float64        `json:"price"`
}

func (c AlertFrequency) IsValidAlertFrequency() error {
	switch c {
	case OnlyOnce, OnceADay, Always:
		return nil
	}
	return errors.ErrInvalidAlertFrequencyType
}

func (c AlertType) IsValidAlertType() error {
	switch c {
	case PriceDropsTo, PriceReaches, PriceRisesAbove, ChangeIsOver, ChangeIsUnder:
		return nil
	}
	return errors.ErrInvalidAlertType
}

func (UserTokenPriceAlert) TableName() string {
	return "user_token_price_alerts"
}

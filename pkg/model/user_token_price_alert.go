package model

import (
	"time"

	"github.com/defipod/mochi/pkg/model/errors"
)

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
	UserDiscordID string         `json:"user_discord_id"`
	Symbol        string         `json:"symbol"`
	Currency      string         `json:"currency"`
	AlertType     AlertType      `json:"alert_type"`
	Frequency     AlertFrequency `json:"frequency"`
	Price         float64        `json:"price"`
	SnoozedTo     time.Time      `json:"snoozed_to"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
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

func (c AlertType) GetRedisKeyPrefix() string {
	switch c {
	case PriceDropsTo, ChangeIsUnder:
		return "alert_direction_down"
	case PriceReaches, PriceRisesAbove, ChangeIsOver:
		return "alert_direction_up"
	default:
		return ""
	}
}

func (UserTokenPriceAlert) TableName() string {
	return "user_token_price_alerts"
}

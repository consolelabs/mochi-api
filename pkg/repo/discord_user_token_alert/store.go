package discord_user_token_alert

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	UpsertOne(config *model.UpsertDiscordUserTokenAlert) error
	RemoveOne(config *model.DiscordUserTokenAlert) error
	GetByDiscordID(discordId string) ([]model.DiscordUserTokenAlert, error)
	GetByDeviceID(deviceId string) ([]model.DiscordUserTokenAlert, error)
	GetAll() ([]model.DiscordUserTokenAlert, error)
	GetAllActive() ([]model.DiscordUserTokenAlert, error)
}

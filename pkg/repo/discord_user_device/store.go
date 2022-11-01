package discord_user_device

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	UpsertOne(config *model.DiscordUserDevice) error
	GetByDeviceID(deviceId string) (*model.DiscordUserDevice, error)
	RemoveByDeviceID(deviceId string) error
}

package guild_config_welcome_channel

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByGuildID(guildID string) (*model.GuildConfigWelcomeChannel, error)
	UpsertOne(config *model.GuildConfigWelcomeChannel) (*model.GuildConfigWelcomeChannel, error)
	DeleteOne(config *model.GuildConfigWelcomeChannel) error
}

package discord_guilds

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Gets() ([]model.DiscordGuild, error)
	GetByID(id string) (*model.DiscordGuild, error)
	CreateIfNotExists(guild model.DiscordGuild) error
	Update(omit string, guild model.DiscordGuild) error
}

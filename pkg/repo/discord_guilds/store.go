package discord_guilds

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Gets() ([]model.DiscordGuild, error)
	GetByID(id string) (*model.DiscordGuild, error)
	CreateOrReactivate(guild model.DiscordGuild) error
	Update(guild *model.DiscordGuild) error
	GetNonLeftGuilds() (guilds []model.DiscordGuild, err error)
}

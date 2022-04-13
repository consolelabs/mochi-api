package discord_guilds

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Gets() ([]*model.DiscordGuild, error)
}

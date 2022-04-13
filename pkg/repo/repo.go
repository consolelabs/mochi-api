package repo

import (
	discordguilds "github.com/defipod/mochi/pkg/repo/discord_guilds"
)

type Repo struct {
	DiscordGuilds discordguilds.Store
}

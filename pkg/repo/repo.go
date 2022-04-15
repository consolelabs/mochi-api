package repo

import (
	discordguilds "github.com/defipod/mochi/pkg/repo/discord_guilds"
	guildusers "github.com/defipod/mochi/pkg/repo/guild_users"
	invitehistories "github.com/defipod/mochi/pkg/repo/invite_histories"
	users "github.com/defipod/mochi/pkg/repo/users"
)

type Repo struct {
	DiscordGuilds   discordguilds.Store
	InviteHistories invitehistories.Store
	Users           users.Store
	GuildUsers      guildusers.Store
}

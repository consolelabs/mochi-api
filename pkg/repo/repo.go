package repo

import (
	discordbottransaction "github.com/defipod/mochi/pkg/repo/discord_bot_transaction"
	discordguilds "github.com/defipod/mochi/pkg/repo/discord_guilds"
	discordusergmstreak "github.com/defipod/mochi/pkg/repo/discord_user_gm_streak"
	guildconfiggmgn "github.com/defipod/mochi/pkg/repo/guild_config_gm_gn"
	guildconfiginvitetracker "github.com/defipod/mochi/pkg/repo/guild_config_invite_tracker"
	guildcustomcommand "github.com/defipod/mochi/pkg/repo/guild_custom_command"
	guildusers "github.com/defipod/mochi/pkg/repo/guild_users"
	invitehistories "github.com/defipod/mochi/pkg/repo/invite_histories"
	token "github.com/defipod/mochi/pkg/repo/token"
	users "github.com/defipod/mochi/pkg/repo/users"
)

type Repo struct {
	DiscordUserGMStreak      discordusergmstreak.Store
	GuildConfigGmGn          guildconfiggmgn.Store
	DiscordGuilds            discordguilds.Store
	InviteHistories          invitehistories.Store
	Users                    users.Store
	GuildUsers               guildusers.Store
	GuildCustomCommand       guildcustomcommand.Store
	Token                    token.Store
	DiscordBotTransaction    discordbottransaction.Store
	GuildConfigInviteTracker guildconfiginvitetracker.Store
}

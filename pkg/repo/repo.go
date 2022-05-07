package repo

import (
	discordbottransaction "github.com/defipod/mochi/pkg/repo/discord_bot_transaction"
	discordguildstatchannels "github.com/defipod/mochi/pkg/repo/discord_guild_stat_channels"
	discordguildstats "github.com/defipod/mochi/pkg/repo/discord_guild_stats"
	discordguilds "github.com/defipod/mochi/pkg/repo/discord_guilds"
	discordusergmstreak "github.com/defipod/mochi/pkg/repo/discord_user_gm_streak"
	discordwalletverification "github.com/defipod/mochi/pkg/repo/discord_wallet_verification"
	guildconfigdefaultrole "github.com/defipod/mochi/pkg/repo/guild_config_default_roles"
	guildconfiggmgn "github.com/defipod/mochi/pkg/repo/guild_config_gm_gn"
	guildconfiginvitetracker "github.com/defipod/mochi/pkg/repo/guild_config_invite_tracker"
	guildconfigreactionrole "github.com/defipod/mochi/pkg/repo/guild_config_reaction_roles"
	guildconfigtoken "github.com/defipod/mochi/pkg/repo/guild_config_token"
	guildconfigwalletverificationmessage "github.com/defipod/mochi/pkg/repo/guild_config_wallet_verification_message"
	guildcustomcommand "github.com/defipod/mochi/pkg/repo/guild_custom_command"
	guildusers "github.com/defipod/mochi/pkg/repo/guild_users"
	invitehistories "github.com/defipod/mochi/pkg/repo/invite_histories"
	token "github.com/defipod/mochi/pkg/repo/token"
	userwallet "github.com/defipod/mochi/pkg/repo/user_wallet"
	users "github.com/defipod/mochi/pkg/repo/users"
)

type Repo struct {
	DiscordUserGMStreak                  discordusergmstreak.Store
	GuildConfigGmGn                      guildconfiggmgn.Store
	GuildConfigWalletVerificationMessage guildconfigwalletverificationmessage.Store
	DiscordGuilds                        discordguilds.Store
	DiscordWalletVerification            discordwalletverification.Store
	InviteHistories                      invitehistories.Store
	Users                                users.Store
	UserWallet                           userwallet.Store
	GuildUsers                           guildusers.Store
	GuildCustomCommand                   guildcustomcommand.Store
	Token                                token.Store
	DiscordBotTransaction                discordbottransaction.Store
	GuildConfigInviteTracker             guildconfiginvitetracker.Store
	GuildConfigReactionRole              guildconfigreactionrole.Store
	GuildConfigDefaultRole               guildconfigdefaultrole.Store
	DiscordGuildStats                    discordguildstats.Store
	DiscordGuildStatChannels             discordguildstatchannels.Store
	GuildConfigToken                     guildconfigtoken.Store
}

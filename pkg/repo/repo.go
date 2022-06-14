package repo

import (
	"github.com/defipod/mochi/pkg/repo/activity"
	"github.com/defipod/mochi/pkg/repo/chain"
	configxplevel "github.com/defipod/mochi/pkg/repo/config_xp_level"
	discordguildstatchannels "github.com/defipod/mochi/pkg/repo/discord_guild_stat_channels"
	discordguildstats "github.com/defipod/mochi/pkg/repo/discord_guild_stats"
	discordguilds "github.com/defipod/mochi/pkg/repo/discord_guilds"
	discordusergmstreak "github.com/defipod/mochi/pkg/repo/discord_user_gm_streak"
	discordwalletverification "github.com/defipod/mochi/pkg/repo/discord_wallet_verification"
	guildconfigactivity "github.com/defipod/mochi/pkg/repo/guild_config_activity"
	guildconfigdefaultrole "github.com/defipod/mochi/pkg/repo/guild_config_default_roles"
	guildconfiggmgn "github.com/defipod/mochi/pkg/repo/guild_config_gm_gn"
	guildconfiginvitetracker "github.com/defipod/mochi/pkg/repo/guild_config_invite_tracker"
	guildconfiglevelrole "github.com/defipod/mochi/pkg/repo/guild_config_level_role"
	guildconfignftrole "github.com/defipod/mochi/pkg/repo/guild_config_nft_role"
	guildconfigreactionrole "github.com/defipod/mochi/pkg/repo/guild_config_reaction_roles"
	guildconfigrepostreaction "github.com/defipod/mochi/pkg/repo/guild_config_repost_reaction"
	guildconfigtoken "github.com/defipod/mochi/pkg/repo/guild_config_token"
	guildconfigwalletverificationmessage "github.com/defipod/mochi/pkg/repo/guild_config_wallet_verification_message"
	guildcustomcommand "github.com/defipod/mochi/pkg/repo/guild_custom_command"
	guilduseractivitylog "github.com/defipod/mochi/pkg/repo/guild_user_activity_log"
	guilduserxp "github.com/defipod/mochi/pkg/repo/guild_user_xp"
	guildusers "github.com/defipod/mochi/pkg/repo/guild_users"
	invitehistories "github.com/defipod/mochi/pkg/repo/invite_histories"
	messagereposthistory "github.com/defipod/mochi/pkg/repo/message_repost_history"
	nftcollection "github.com/defipod/mochi/pkg/repo/nft_collection"
	token "github.com/defipod/mochi/pkg/repo/token"
	userbalance "github.com/defipod/mochi/pkg/repo/user_balance"
	userdefiactivitylog "github.com/defipod/mochi/pkg/repo/user_defi_activity_log"
	usernftbalance "github.com/defipod/mochi/pkg/repo/user_nft_balance"
	userwallet "github.com/defipod/mochi/pkg/repo/user_wallet"
	users "github.com/defipod/mochi/pkg/repo/users"
	whitelistcampaignusers "github.com/defipod/mochi/pkg/repo/whitelist_campaign_users"
	whitelistcampaigns "github.com/defipod/mochi/pkg/repo/whitelist_campaigns"
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
	GuildConfigInviteTracker             guildconfiginvitetracker.Store
	GuildConfigReactionRole              guildconfigreactionrole.Store
	GuildConfigDefaultRole               guildconfigdefaultrole.Store
	GuildConfigRepostReaction            guildconfigrepostreaction.Store
	DiscordGuildStats                    discordguildstats.Store
	DiscordGuildStatChannels             discordguildstatchannels.Store
	GuildConfigToken                     guildconfigtoken.Store
	WhitelistCampaigns                   whitelistcampaigns.Store
	WhitelistCampaignUsers               whitelistcampaignusers.Store
	NFTCollection                        nftcollection.Store
	Activity                             activity.Store
	GuildConfigActivity                  guildconfigactivity.Store
	ConfigXPLevel                        configxplevel.Store
	GuildUserActivityLog                 guilduseractivitylog.Store
	GuildUserXP                          guilduserxp.Store
	GuildConfigLevelRole                 guildconfiglevelrole.Store
	Chain                                chain.Store
	GuildConfigNFTRole                   guildconfignftrole.Store
	UserNFTBalance                       usernftbalance.Store
	MessageRepostHistory                 messagereposthistory.Store
	UserDefiActivityLog                  userdefiactivitylog.Store
	UserBalance                          userbalance.Store
}

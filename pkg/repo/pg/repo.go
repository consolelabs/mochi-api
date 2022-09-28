package pg

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/repo"
	"github.com/defipod/mochi/pkg/repo/activity"
	"github.com/defipod/mochi/pkg/repo/chain"
	coingeckosupportedtokens "github.com/defipod/mochi/pkg/repo/coingecko_supported_tokens"
	configxplevel "github.com/defipod/mochi/pkg/repo/config_xp_level"
	discordguildstatchannels "github.com/defipod/mochi/pkg/repo/discord_guild_stat_channels"
	discordguildstats "github.com/defipod/mochi/pkg/repo/discord_guild_stats"
	discordguilds "github.com/defipod/mochi/pkg/repo/discord_guilds"
	discordusergmstreak "github.com/defipod/mochi/pkg/repo/discord_user_gm_streak"
	discorduserupvotelog "github.com/defipod/mochi/pkg/repo/discord_user_upvote_log"
	discorduserupvotestreak "github.com/defipod/mochi/pkg/repo/discord_user_upvote_streak"
	discordwalletverification "github.com/defipod/mochi/pkg/repo/discord_wallet_verification"
	guildconfigactivity "github.com/defipod/mochi/pkg/repo/guild_config_activity"
	guildconfigdefaultcollection "github.com/defipod/mochi/pkg/repo/guild_config_default_collection"
	guildconfigdefaultrole "github.com/defipod/mochi/pkg/repo/guild_config_default_roles"
	guildconfigdefaultticker "github.com/defipod/mochi/pkg/repo/guild_config_default_ticker"
	guildconfiggmgn "github.com/defipod/mochi/pkg/repo/guild_config_gm_gn"
	guildconfiggroupnftrole "github.com/defipod/mochi/pkg/repo/guild_config_group_nft_role"
	guildconfiginvitetracker "github.com/defipod/mochi/pkg/repo/guild_config_invite_tracker"
	guildconfiglevelrole "github.com/defipod/mochi/pkg/repo/guild_config_level_role"
	guildconfignftrole "github.com/defipod/mochi/pkg/repo/guild_config_nft_role"
	guildconfigpruneexclude "github.com/defipod/mochi/pkg/repo/guild_config_prune_exclude"
	guildconfigreactionrole "github.com/defipod/mochi/pkg/repo/guild_config_reaction_roles"
	guildconfigrepostreaction "github.com/defipod/mochi/pkg/repo/guild_config_repost_reaction"
	guildconfigsalestracker "github.com/defipod/mochi/pkg/repo/guild_config_sales_tracker"
	guildconfigtoken "github.com/defipod/mochi/pkg/repo/guild_config_token"
	guildconfigtwitterfeed "github.com/defipod/mochi/pkg/repo/guild_config_twitter_feed"
	guildconfigtwitterhashtag "github.com/defipod/mochi/pkg/repo/guild_config_twitter_hashtag"
	guildconfigvotechannel "github.com/defipod/mochi/pkg/repo/guild_config_vote_channel"
	guildconfigwalletverificationmessage "github.com/defipod/mochi/pkg/repo/guild_config_wallet_verification_message"
	guildconfigwelcomechannel "github.com/defipod/mochi/pkg/repo/guild_config_welcome_channel"
	guildcustomcommand "github.com/defipod/mochi/pkg/repo/guild_custom_command"
	guildscheduledevent "github.com/defipod/mochi/pkg/repo/guild_scheduled_event"
	guilduseractivitylog "github.com/defipod/mochi/pkg/repo/guild_user_activity_log"
	guilduserxp "github.com/defipod/mochi/pkg/repo/guild_user_xp"
	guildusers "github.com/defipod/mochi/pkg/repo/guild_users"
	invitehistories "github.com/defipod/mochi/pkg/repo/invite_histories"
	messagereaction "github.com/defipod/mochi/pkg/repo/message_reaction"
	messagereposthistory "github.com/defipod/mochi/pkg/repo/message_repost_history"
	mochinftsales "github.com/defipod/mochi/pkg/repo/mochi_nft_sales"
	nftcollection "github.com/defipod/mochi/pkg/repo/nft_collection"
	nftsalestracker "github.com/defipod/mochi/pkg/repo/nft_sales_tracker"
	serversusagestats "github.com/defipod/mochi/pkg/repo/servers_usage_stats"
	"github.com/defipod/mochi/pkg/repo/token"
	twitterpost "github.com/defipod/mochi/pkg/repo/twitter_post"
	upvotestreaktier "github.com/defipod/mochi/pkg/repo/upvote_streak_tiers"
	usernftbalance "github.com/defipod/mochi/pkg/repo/user_nft_balance"
	usernftwatchlistitem "github.com/defipod/mochi/pkg/repo/user_nft_watchlist_items"
	usertelegramdiscordassociation "github.com/defipod/mochi/pkg/repo/user_telegram_discord_association"
	userwallet "github.com/defipod/mochi/pkg/repo/user_wallet"
	userwatchlistitem "github.com/defipod/mochi/pkg/repo/user_watchlist_item"
	"github.com/defipod/mochi/pkg/repo/users"
	whitelistcampaignusers "github.com/defipod/mochi/pkg/repo/whitelist_campaign_users"
	whitelistcampaigns "github.com/defipod/mochi/pkg/repo/whitelist_campaigns"
)

// NewRepo new pg repo implementation
func NewRepo(db *gorm.DB) *repo.Repo {
	return &repo.Repo{
		DiscordGuilds:                        discordguilds.NewPG(db),
		DiscordWalletVerification:            discordwalletverification.NewPG(db),
		InviteHistories:                      invitehistories.NewPG(db),
		Users:                                users.NewPG(db),
		UserWallet:                           userwallet.NewPG(db),
		GuildUsers:                           guildusers.NewPG(db),
		GuildCustomCommand:                   guildcustomcommand.NewPG(db),
		Token:                                token.NewPG(db),
		DiscordUserGMStreak:                  discordusergmstreak.NewPG(db),
		GuildConfigWelcomeChannel:            guildconfigwelcomechannel.NewPG(db),
		GuildConfigVoteChannel:               guildconfigvotechannel.NewPG(db),
		DiscordUserUpvoteStreak:              discorduserupvotestreak.NewPG(db),
		GuildConfigGmGn:                      guildconfiggmgn.NewPG(db),
		DiscordUserUpvoteLog:                 discorduserupvotelog.NewPG(db),
		GuildConfigSalesTracker:              guildconfigsalestracker.NewPG(db),
		GuildConfigInviteTracker:             guildconfiginvitetracker.NewPG(db),
		GuildConfigReactionRole:              guildconfigreactionrole.NewPG(db),
		GuildConfigDefaultRole:               guildconfigdefaultrole.NewPG(db),
		GuildConfigDefaultCollection:         guildconfigdefaultcollection.NewPG(db),
		GuildConfigRepostReaction:            guildconfigrepostreaction.NewPG(db),
		GuildConfigWalletVerificationMessage: guildconfigwalletverificationmessage.NewPG(db),
		UpvoteStreakTier:                     upvotestreaktier.NewPG(db),
		GuildConfigTwitterFeed:               guildconfigtwitterfeed.NewPG(db),
		GuildConfigTwitterHashtag:            guildconfigtwitterhashtag.NewPG(db),
		GuildConfigPruneExclude:              guildconfigpruneexclude.NewPG(db),
		DiscordGuildStats:                    discordguildstats.NewPG(db),
		DiscordGuildStatChannels:             discordguildstatchannels.NewPG(db),
		GuildConfigToken:                     guildconfigtoken.NewPG(db),
		WhitelistCampaigns:                   whitelistcampaigns.NewPG(db),
		WhitelistCampaignUsers:               whitelistcampaignusers.NewPG(db),
		NFTCollection:                        nftcollection.NewPG(db),
		TwitterPost:                          twitterpost.NewPG(db),
		NFTSalesTracker:                      nftsalestracker.NewPG(db),
		Activity:                             activity.NewPG(db),
		GuildConfigActivity:                  guildconfigactivity.NewPG(db),
		ConfigXPLevel:                        configxplevel.NewPG(db),
		GuildUserActivityLog:                 guilduseractivitylog.NewPG(db),
		GuildUserXP:                          guilduserxp.NewPG(db),
		GuildConfigLevelRole:                 guildconfiglevelrole.NewPG(db),
		Chain:                                chain.NewPG(db),
		GuildConfigNFTRole:                   guildconfignftrole.NewPG(db),
		UserNFTBalance:                       usernftbalance.NewPG(db),
		MessageRepostHistory:                 messagereposthistory.NewPG(db),
		GuildScheduledEvent:                  guildscheduledevent.NewPG(db),
		MochiNFTSales:                        mochinftsales.NewPG(db),
		GuildConfigDefaultTicker:             guildconfigdefaultticker.NewPG(db),
		UserWatchlistItem:                    userwatchlistitem.NewPG(db),
		GuildConfigGroupNFTRole:              guildconfiggroupnftrole.NewPG(db),
		CoingeckoSupportedTokens:             coingeckosupportedtokens.NewPG(db),
		UserTelegramDiscordAssociation:       usertelegramdiscordassociation.NewPG(db),
		ServersUsageStats:                    serversusagestats.NewPG(db),
		MessageReaction:                      messagereaction.NewPG(db),
		UserNftWatchlistItem:                 usernftwatchlistitem.NewPG(db),
	}
}

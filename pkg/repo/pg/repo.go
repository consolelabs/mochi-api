package pg

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/repo"
	"github.com/defipod/mochi/pkg/repo/activity"
	"github.com/defipod/mochi/pkg/repo/chain"
	configxplevel "github.com/defipod/mochi/pkg/repo/config_xp_level"
	discordguildstatchannels "github.com/defipod/mochi/pkg/repo/discord_guild_stat_channels"
	discordguildstats "github.com/defipod/mochi/pkg/repo/discord_guild_stats"
	discordguilds "github.com/defipod/mochi/pkg/repo/discord_guilds"
	discordusergmstreak "github.com/defipod/mochi/pkg/repo/discord_user_gm_streak"
	discordwalletverification "github.com/defipod/mochi/pkg/repo/discord_wallet_verification"
	guildconfigactivity "github.com/defipod/mochi/pkg/repo/guild_config_activity"
	guildconfigdefaultcollection "github.com/defipod/mochi/pkg/repo/guild_config_default_collection"
	guildconfigdefaultrole "github.com/defipod/mochi/pkg/repo/guild_config_default_roles"
	guildconfiggmgn "github.com/defipod/mochi/pkg/repo/guild_config_gm_gn"
	guildconfiginvitetracker "github.com/defipod/mochi/pkg/repo/guild_config_invite_tracker"
	guildconfiglevelrole "github.com/defipod/mochi/pkg/repo/guild_config_level_role"
	guildconfignftrole "github.com/defipod/mochi/pkg/repo/guild_config_nft_role"
	guildconfigreactionrole "github.com/defipod/mochi/pkg/repo/guild_config_reaction_roles"
	guildconfigrepostreaction "github.com/defipod/mochi/pkg/repo/guild_config_repost_reaction"
	guildconfigsalestracker "github.com/defipod/mochi/pkg/repo/guild_config_sales_tracker"
	guildconfigtoken "github.com/defipod/mochi/pkg/repo/guild_config_token"
	guildconfigtwitterfeed "github.com/defipod/mochi/pkg/repo/guild_config_twitter_feed"
	guildconfigtwitterhashtag "github.com/defipod/mochi/pkg/repo/guild_config_twitter_hashtag"
	guildconfigwalletverificationmessage "github.com/defipod/mochi/pkg/repo/guild_config_wallet_verification_message"
	guildcustomcommand "github.com/defipod/mochi/pkg/repo/guild_custom_command"
	guildscheduledevent "github.com/defipod/mochi/pkg/repo/guild_scheduled_event"
	guilduseractivitylog "github.com/defipod/mochi/pkg/repo/guild_user_activity_log"
	guilduserxp "github.com/defipod/mochi/pkg/repo/guild_user_xp"
	guildusers "github.com/defipod/mochi/pkg/repo/guild_users"
	invitehistories "github.com/defipod/mochi/pkg/repo/invite_histories"
	messagereposthistory "github.com/defipod/mochi/pkg/repo/message_repost_history"
	mochinftsales "github.com/defipod/mochi/pkg/repo/mochi_nft_sales"
	nftcollection "github.com/defipod/mochi/pkg/repo/nft_collection"
	nftsalestracker "github.com/defipod/mochi/pkg/repo/nft_sales_tracker"
	"github.com/defipod/mochi/pkg/repo/token"
	twitterpost "github.com/defipod/mochi/pkg/repo/twitter_post"
	usernftbalance "github.com/defipod/mochi/pkg/repo/user_nft_balance"
	userwallet "github.com/defipod/mochi/pkg/repo/user_wallet"
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
		GuildConfigGmGn:                      guildconfiggmgn.NewPG(db),
		GuildConfigSalesTracker:              guildconfigsalestracker.NewPG(db),
		GuildConfigInviteTracker:             guildconfiginvitetracker.NewPG(db),
		GuildConfigReactionRole:              guildconfigreactionrole.NewPG(db),
		GuildConfigDefaultRole:               guildconfigdefaultrole.NewPG(db),
		GuildConfigDefaultCollection:         guildconfigdefaultcollection.NewPG(db),
		GuildConfigRepostReaction:            guildconfigrepostreaction.NewPG(db),
		GuildConfigWalletVerificationMessage: guildconfigwalletverificationmessage.NewPG(db),
		GuildConfigTwitterFeed:               guildconfigtwitterfeed.NewPG(db),
		GuildConfigTwitterHashtag:            guildconfigtwitterhashtag.NewPG(db),
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
	}
}

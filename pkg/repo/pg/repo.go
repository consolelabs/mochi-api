package pg

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/repo"
	"github.com/defipod/mochi/pkg/repo/activity"
	"github.com/defipod/mochi/pkg/repo/chain"
	coingeckosupportedtokens "github.com/defipod/mochi/pkg/repo/coingecko_supported_tokens"
	configxplevel "github.com/defipod/mochi/pkg/repo/config_xp_level"
	conversationreposthistories "github.com/defipod/mochi/pkg/repo/conversation_repost_histories"
	discordguildstatchannels "github.com/defipod/mochi/pkg/repo/discord_guild_stat_channels"
	discordguildstats "github.com/defipod/mochi/pkg/repo/discord_guild_stats"
	discordguilds "github.com/defipod/mochi/pkg/repo/discord_guilds"
	discorduserdevice "github.com/defipod/mochi/pkg/repo/discord_user_device"
	discordusergmstreak "github.com/defipod/mochi/pkg/repo/discord_user_gm_streak"
	discordusertokenalert "github.com/defipod/mochi/pkg/repo/discord_user_token_alert"
	discorduserupvotelog "github.com/defipod/mochi/pkg/repo/discord_user_upvote_log"
	discorduserupvotestreak "github.com/defipod/mochi/pkg/repo/discord_user_upvote_streak"
	discordwalletverification "github.com/defipod/mochi/pkg/repo/discord_wallet_verification"
	gitbookclickcollector "github.com/defipod/mochi/pkg/repo/gitbook_click_collectors"
	guildblacklistchannelrepostconfigs "github.com/defipod/mochi/pkg/repo/guild_blacklist_channel_repost_configs"
	guildconfigactivity "github.com/defipod/mochi/pkg/repo/guild_config_activity"
	guildconfigdefaultcollection "github.com/defipod/mochi/pkg/repo/guild_config_default_collection"
	guildconfigdefaultcurrency "github.com/defipod/mochi/pkg/repo/guild_config_default_currency"
	guildconfigdefaultrole "github.com/defipod/mochi/pkg/repo/guild_config_default_roles"
	guildconfigdefaultticker "github.com/defipod/mochi/pkg/repo/guild_config_default_ticker"
	guildconfiggmgn "github.com/defipod/mochi/pkg/repo/guild_config_gm_gn"
	guildconfiggroupnftrole "github.com/defipod/mochi/pkg/repo/guild_config_group_nft_role"
	guildconfiginvitetracker "github.com/defipod/mochi/pkg/repo/guild_config_invite_tracker"
	guildconfigjoinleavechannel "github.com/defipod/mochi/pkg/repo/guild_config_join_leave_channel"
	guildconfiglevelrole "github.com/defipod/mochi/pkg/repo/guild_config_level_role"
	guildconfignftrole "github.com/defipod/mochi/pkg/repo/guild_config_nft_role"
	guildconfigpruneexclude "github.com/defipod/mochi/pkg/repo/guild_config_prune_exclude"
	guildconfigreactionrole "github.com/defipod/mochi/pkg/repo/guild_config_reaction_roles"
	guildconfigrepostreaction "github.com/defipod/mochi/pkg/repo/guild_config_repost_reaction"
	guildconfigsalestracker "github.com/defipod/mochi/pkg/repo/guild_config_sales_tracker"
	guildconfigtoken "github.com/defipod/mochi/pkg/repo/guild_config_token"
	guildconfigtwitterblacklist "github.com/defipod/mochi/pkg/repo/guild_config_twitter_blacklist"
	guildconfigtwitterfeed "github.com/defipod/mochi/pkg/repo/guild_config_twitter_feed"
	guildconfigtwitterhashtag "github.com/defipod/mochi/pkg/repo/guild_config_twitter_hashtag"
	guildconfigvotechannel "github.com/defipod/mochi/pkg/repo/guild_config_vote_channel"
	guildconfigwalletverificationmessage "github.com/defipod/mochi/pkg/repo/guild_config_wallet_verification_message"
	guildconfigwelcomechannel "github.com/defipod/mochi/pkg/repo/guild_config_welcome_channel"
	guildscheduledevent "github.com/defipod/mochi/pkg/repo/guild_scheduled_event"
	guilduseractivitylog "github.com/defipod/mochi/pkg/repo/guild_user_activity_log"
	guilduserxp "github.com/defipod/mochi/pkg/repo/guild_user_xp"
	guildusers "github.com/defipod/mochi/pkg/repo/guild_users"
	invitehistories "github.com/defipod/mochi/pkg/repo/invite_histories"
	messagereaction "github.com/defipod/mochi/pkg/repo/message_reaction"
	messagereposthistory "github.com/defipod/mochi/pkg/repo/message_repost_history"
	mochinftsales "github.com/defipod/mochi/pkg/repo/mochi_nft_sales"
	monikerconfig "github.com/defipod/mochi/pkg/repo/moniker_config"
	nftaddrequesthistory "github.com/defipod/mochi/pkg/repo/nft_add_request_history"
	nftcollection "github.com/defipod/mochi/pkg/repo/nft_collection"
	nftsalestracker "github.com/defipod/mochi/pkg/repo/nft_sales_tracker"
	offchaintipbotactivitylogs "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_activity_logs"
	offchaintipbotchain "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_chain"
	offchaintipbotconfignotify "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_config_notify"
	offchaintipbotcontract "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_contract"
	offchaintipbotdepositlog "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_deposit_log"
	offchaintipbottokens "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_tokens"
	offchaintipbottransferhistories "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_transfer_histories"
	offchaintipbotuserbalances "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_user_balances"
	"github.com/defipod/mochi/pkg/repo/quest"
	questpass "github.com/defipod/mochi/pkg/repo/quest_pass"
	questreward "github.com/defipod/mochi/pkg/repo/quest_reward"
	questrewardtype "github.com/defipod/mochi/pkg/repo/quest_reward_type"
	queststreak "github.com/defipod/mochi/pkg/repo/quest_streak"
	questuserlist "github.com/defipod/mochi/pkg/repo/quest_user_list"
	questuserlog "github.com/defipod/mochi/pkg/repo/quest_user_log"
	questuserpass "github.com/defipod/mochi/pkg/repo/quest_user_pass"
	questuserreward "github.com/defipod/mochi/pkg/repo/quest_user_reward"
	serversusagestats "github.com/defipod/mochi/pkg/repo/servers_usage_stats"
	"github.com/defipod/mochi/pkg/repo/token"
	tradeoffer "github.com/defipod/mochi/pkg/repo/trade_offer"
	twitterpost "github.com/defipod/mochi/pkg/repo/twitter_post"
	twitterpoststreak "github.com/defipod/mochi/pkg/repo/twitter_post_streak"
	upvotestreaktier "github.com/defipod/mochi/pkg/repo/upvote_streak_tiers"
	userfeedback "github.com/defipod/mochi/pkg/repo/user_feedback"
	usernftbalance "github.com/defipod/mochi/pkg/repo/user_nft_balance"
	usernftwatchlistitem "github.com/defipod/mochi/pkg/repo/user_nft_watchlist_items"
	usertelegramdiscordassociation "github.com/defipod/mochi/pkg/repo/user_telegram_discord_association"
	userwallet "github.com/defipod/mochi/pkg/repo/user_wallet"
	userwatchlistitem "github.com/defipod/mochi/pkg/repo/user_watchlist_item"
	"github.com/defipod/mochi/pkg/repo/users"
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
		Token:                                token.NewPG(db),
		DiscordUserGMStreak:                  discordusergmstreak.NewPG(db),
		GuildConfigWelcomeChannel:            guildconfigwelcomechannel.NewPG(db),
		GuildConfigVoteChannel:               guildconfigvotechannel.NewPG(db),
		DiscordUserUpvoteStreak:              discorduserupvotestreak.NewPG(db),
		GuildConfigGmGn:                      guildconfiggmgn.NewPG(db),
		DiscordUserUpvoteLog:                 discorduserupvotelog.NewPG(db),
		GuildConfigSalesTracker:              guildconfigsalestracker.NewPG(db),
		DiscordUserTokenAlert:                discordusertokenalert.NewPG(db),
		DiscordUserDevice:                    discorduserdevice.NewPG(db),
		GuildConfigInviteTracker:             guildconfiginvitetracker.NewPG(db),
		GuildConfigReactionRole:              guildconfigreactionrole.NewPG(db),
		GuildConfigDefaultCurrency:           guildconfigdefaultcurrency.NewPG(db),
		GuildConfigDefaultRole:               guildconfigdefaultrole.NewPG(db),
		GuildConfigJoinLeaveChannel:          guildconfigjoinleavechannel.NewPG(db),
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
		NFTCollection:                        nftcollection.NewPG(db),
		TwitterPost:                          twitterpost.NewPG(db),
		TwitterPostStreak:                    twitterpoststreak.NewPG(db),
		NFTSalesTracker:                      nftsalestracker.NewPG(db),
		UserFeedback:                         userfeedback.NewPG(db),
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
		Quest:                                quest.NewPG(db),
		QuestRewardType:                      questrewardtype.NewPG(db),
		QuestUserLog:                         questuserlog.NewPG(db),
		QuestUserList:                        questuserlist.NewPG(db),
		QuestPass:                            questpass.NewPG(db),
		QuestReward:                          questreward.NewPG(db),
		QuestUserReward:                      questuserreward.NewPG(db),
		QuestUserPass:                        questuserpass.NewPG(db),
		QuestStreak:                          queststreak.NewPG(db),
		ConversationRepostHistories:          conversationreposthistories.NewPG(db),
		GitbookClickCollector:                gitbookclickcollector.NewPG(db),
		OffchainTipBotChain:                  offchaintipbotchain.NewPG(db),
		OffchainTipBotContract:               offchaintipbotcontract.NewPG(db),
		TradeOffer:                           tradeoffer.NewPG(db),
		OffchainTipBotUserBalances:           offchaintipbotuserbalances.NewPG(db),
		GuildBlacklistChannelRepostConfigs:   guildblacklistchannelrepostconfigs.NewPG(db),
		OffchainTipBotTokens:                 offchaintipbottokens.NewPG(db),
		OffchainTipBotActivityLogs:           offchaintipbotactivitylogs.NewPG(db),
		OffchainTipBotTransferHistories:      offchaintipbottransferhistories.NewPG(db),
		GuildConfigTwitterBlacklist:          guildconfigtwitterblacklist.NewPG(db),
		MonikerConfig:                        monikerconfig.NewPG(db),
		OffchainTipBotConfigNotify:           offchaintipbotconfignotify.NewPG(db),
		NftAddRequestHistory:                 nftaddrequesthistory.NewPG(db),
		OffchainTipBotDepositLog:             offchaintipbotdepositlog.NewPG(db),
	}
}

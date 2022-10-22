package repo

import (
	"github.com/defipod/mochi/pkg/repo/activity"
	"github.com/defipod/mochi/pkg/repo/chain"
	coingeckosupportedtokens "github.com/defipod/mochi/pkg/repo/coingecko_supported_tokens"
	configxplevel "github.com/defipod/mochi/pkg/repo/config_xp_level"
	conversationreposthistories "github.com/defipod/mochi/pkg/repo/conversation_repost_histories"
	discordguildstatchannels "github.com/defipod/mochi/pkg/repo/discord_guild_stat_channels"
	discordguildstats "github.com/defipod/mochi/pkg/repo/discord_guild_stats"
	discordguilds "github.com/defipod/mochi/pkg/repo/discord_guilds"
	discordusergmstreak "github.com/defipod/mochi/pkg/repo/discord_user_gm_streak"
	discorduserupvotelog "github.com/defipod/mochi/pkg/repo/discord_user_upvote_log"
	discorduserupvotestreak "github.com/defipod/mochi/pkg/repo/discord_user_upvote_streak"
	discordwalletverification "github.com/defipod/mochi/pkg/repo/discord_wallet_verification"
	gitbookclickcollector "github.com/defipod/mochi/pkg/repo/gitbook_click_collectors"
	guildblacklistchannelrepostconfigs "github.com/defipod/mochi/pkg/repo/guild_blacklist_channel_repost_configs"
	guildconfigactivity "github.com/defipod/mochi/pkg/repo/guild_config_activity"
	guildconfigdefaultcollection "github.com/defipod/mochi/pkg/repo/guild_config_default_collection"
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
	offchaintipbotactivitylogs "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_activity_logs"
	offchaintipbotchain "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_chain"
	offchaintipbotcontract "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_contract"
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
	token "github.com/defipod/mochi/pkg/repo/token"
	tradeoffer "github.com/defipod/mochi/pkg/repo/trade_offer"
	twitterpost "github.com/defipod/mochi/pkg/repo/twitter_post"
	upvotestreaktier "github.com/defipod/mochi/pkg/repo/upvote_streak_tiers"
	usernftbalance "github.com/defipod/mochi/pkg/repo/user_nft_balance"
	usernftwatchlistitem "github.com/defipod/mochi/pkg/repo/user_nft_watchlist_items"
	usertelegramdiscordassociation "github.com/defipod/mochi/pkg/repo/user_telegram_discord_association"
	userwallet "github.com/defipod/mochi/pkg/repo/user_wallet"
	userwatchlistitem "github.com/defipod/mochi/pkg/repo/user_watchlist_item"
	users "github.com/defipod/mochi/pkg/repo/users"
	whitelistcampaignusers "github.com/defipod/mochi/pkg/repo/whitelist_campaign_users"
	whitelistcampaigns "github.com/defipod/mochi/pkg/repo/whitelist_campaigns"
)

type Repo struct {
	DiscordUserGMStreak                  discordusergmstreak.Store
	DiscordUserUpvoteStreak              discorduserupvotestreak.Store
	GuildConfigGmGn                      guildconfiggmgn.Store
	GuildConfigSalesTracker              guildconfigsalestracker.Store
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
	GuildConfigWelcomeChannel            guildconfigwelcomechannel.Store
	GuildConfigReactionRole              guildconfigreactionrole.Store
	DiscordUserUpvoteLog                 discorduserupvotelog.Store
	GuildConfigDefaultRole               guildconfigdefaultrole.Store
	GuildConfigDefaultCollection         guildconfigdefaultcollection.Store
	GuildConfigPruneExclude              guildconfigpruneexclude.Store
	GuildConfigJoinLeaveChannel          guildconfigjoinleavechannel.Store
	GuildConfigRepostReaction            guildconfigrepostreaction.Store
	GuildConfigTwitterFeed               guildconfigtwitterfeed.Store
	GuildConfigVoteChannel               guildconfigvotechannel.Store
	GuildConfigTwitterHashtag            guildconfigtwitterhashtag.Store
	DiscordGuildStats                    discordguildstats.Store
	DiscordGuildStatChannels             discordguildstatchannels.Store
	UpvoteStreakTier                     upvotestreaktier.Store
	GuildConfigToken                     guildconfigtoken.Store
	WhitelistCampaigns                   whitelistcampaigns.Store
	WhitelistCampaignUsers               whitelistcampaignusers.Store
	NFTCollection                        nftcollection.Store
	Activity                             activity.Store
	TwitterPost                          twitterpost.Store
	GuildConfigActivity                  guildconfigactivity.Store
	ConfigXPLevel                        configxplevel.Store
	GuildUserActivityLog                 guilduseractivitylog.Store
	GuildUserXP                          guilduserxp.Store
	GuildConfigLevelRole                 guildconfiglevelrole.Store
	Chain                                chain.Store
	GuildConfigNFTRole                   guildconfignftrole.Store
	UserNFTBalance                       usernftbalance.Store
	MessageRepostHistory                 messagereposthistory.Store
	GuildScheduledEvent                  guildscheduledevent.Store
	NFTSalesTracker                      nftsalestracker.Store
	MochiNFTSales                        mochinftsales.Store
	GuildConfigDefaultTicker             guildconfigdefaultticker.Store
	UserWatchlistItem                    userwatchlistitem.Store
	GuildConfigGroupNFTRole              guildconfiggroupnftrole.Store
	CoingeckoSupportedTokens             coingeckosupportedtokens.Store
	UserTelegramDiscordAssociation       usertelegramdiscordassociation.Store
	ServersUsageStats                    serversusagestats.Store
	MessageReaction                      messagereaction.Store
	UserNftWatchlistItem                 usernftwatchlistitem.Store
	ConversationRepostHistories          conversationreposthistories.Store
	Quest                                quest.Store
	QuestRewardType                      questrewardtype.Store
	QuestUserLog                         questuserlog.Store
	QuestUserList                        questuserlist.Store
	QuestPass                            questpass.Store
	QuestUserPass                        questuserpass.Store
	QuestReward                          questreward.Store
	QuestUserReward                      questuserreward.Store
	GitbookClickCollector                gitbookclickcollector.Store
	OffchainTipBotChain                  offchaintipbotchain.Store
	OffchainTipBotContract               offchaintipbotcontract.Store
	TradeOffer                           tradeoffer.Store
	QuestStreak                          queststreak.Store
	OffchainTipBotUserBalances           offchaintipbotuserbalances.Store
	GuildBlacklistChannelRepostConfigs   guildblacklistchannelrepostconfigs.Store
	OffchainTipBotTokens                 offchaintipbottokens.Store
	OffchainTipBotActivityLogs           offchaintipbotactivitylogs.Store
	OffchainTipBotTransferHistories      offchaintipbottransferhistories.Store
}

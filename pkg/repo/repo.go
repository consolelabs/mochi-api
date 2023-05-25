package repo

import (
	"github.com/defipod/mochi/pkg/repo/activity"
	"github.com/defipod/mochi/pkg/repo/chain"
	coingeckosupportedtokens "github.com/defipod/mochi/pkg/repo/coingecko_supported_tokens"
	commonwealthdiscussionsubscription "github.com/defipod/mochi/pkg/repo/commonwealth_discussion_subscriptions"
	commonwealthlastestdata "github.com/defipod/mochi/pkg/repo/commonwealth_latest_data"
	configxplevel "github.com/defipod/mochi/pkg/repo/config_xp_level"
	conversationreposthistories "github.com/defipod/mochi/pkg/repo/conversation_repost_histories"
	daoguidelinemessages "github.com/defipod/mochi/pkg/repo/dao_guideline_messages"
	daoproposal "github.com/defipod/mochi/pkg/repo/dao_proposal"
	daoproposalvoteoption "github.com/defipod/mochi/pkg/repo/dao_proposal_vote_option"
	daovote "github.com/defipod/mochi/pkg/repo/dao_vote"
	daovoteoption "github.com/defipod/mochi/pkg/repo/dao_vote_option"
	discordguildstatchannels "github.com/defipod/mochi/pkg/repo/discord_guild_stat_channels"
	discordguildstats "github.com/defipod/mochi/pkg/repo/discord_guild_stats"
	discordguilds "github.com/defipod/mochi/pkg/repo/discord_guilds"
	discorduserdevice "github.com/defipod/mochi/pkg/repo/discord_user_device"
	discordusergmstreak "github.com/defipod/mochi/pkg/repo/discord_user_gm_streak"
	discordusertokenalert "github.com/defipod/mochi/pkg/repo/discord_user_token_alert"
	discorduserupvotelog "github.com/defipod/mochi/pkg/repo/discord_user_upvote_log"
	discorduserupvotestreak "github.com/defipod/mochi/pkg/repo/discord_user_upvote_streak"
	discordwalletverification "github.com/defipod/mochi/pkg/repo/discord_wallet_verification"
	"github.com/defipod/mochi/pkg/repo/envelop"
	guildblacklistchannelrepostconfigs "github.com/defipod/mochi/pkg/repo/guild_blacklist_channel_repost_configs"
	guildconfigactivity "github.com/defipod/mochi/pkg/repo/guild_config_activity"
	guildconfigadminrole "github.com/defipod/mochi/pkg/repo/guild_config_admin_role"
	guildconfigdaoproposal "github.com/defipod/mochi/pkg/repo/guild_config_dao_proposal"
	guildconfigdaotracker "github.com/defipod/mochi/pkg/repo/guild_config_dao_tracker"
	guildconfigdefaultcollection "github.com/defipod/mochi/pkg/repo/guild_config_default_collection"
	guildconfigdefaultcurrency "github.com/defipod/mochi/pkg/repo/guild_config_default_currency"
	guildconfigdefaultrole "github.com/defipod/mochi/pkg/repo/guild_config_default_roles"
	guildconfigdefaultticker "github.com/defipod/mochi/pkg/repo/guild_config_default_ticker"
	guildconfiggmgn "github.com/defipod/mochi/pkg/repo/guild_config_gm_gn"
	guildconfiggroupnftrole "github.com/defipod/mochi/pkg/repo/guild_config_group_nft_role"
	guildconfiginvitetracker "github.com/defipod/mochi/pkg/repo/guild_config_invite_tracker"
	guildconfigjoinleavechannel "github.com/defipod/mochi/pkg/repo/guild_config_join_leave_channel"
	guildconfiglevelrole "github.com/defipod/mochi/pkg/repo/guild_config_level_role"
	guildconfiglevelupmessage "github.com/defipod/mochi/pkg/repo/guild_config_levelup_message"
	guildconfigmixrole "github.com/defipod/mochi/pkg/repo/guild_config_mix_role"
	guildconfignftrole "github.com/defipod/mochi/pkg/repo/guild_config_nft_role"
	guildconfigpruneexclude "github.com/defipod/mochi/pkg/repo/guild_config_prune_exclude"
	guildconfigreactionrole "github.com/defipod/mochi/pkg/repo/guild_config_reaction_roles"
	guildconfigrepostreaction "github.com/defipod/mochi/pkg/repo/guild_config_repost_reaction"
	guildconfigsalestracker "github.com/defipod/mochi/pkg/repo/guild_config_sales_tracker"
	guildconfigtiprange "github.com/defipod/mochi/pkg/repo/guild_config_tip_range"
	guildconfigtoken "github.com/defipod/mochi/pkg/repo/guild_config_token"
	guildconfigtokenrole "github.com/defipod/mochi/pkg/repo/guild_config_token_role"
	guildconfigtwitterblacklist "github.com/defipod/mochi/pkg/repo/guild_config_twitter_blacklist"
	guildconfigtwitterfeed "github.com/defipod/mochi/pkg/repo/guild_config_twitter_feed"
	guildconfigtwitterhashtag "github.com/defipod/mochi/pkg/repo/guild_config_twitter_hashtag"
	guildconfigvotechannel "github.com/defipod/mochi/pkg/repo/guild_config_vote_channel"
	guildconfigwalletverificationmessage "github.com/defipod/mochi/pkg/repo/guild_config_wallet_verification_message"
	guildconfigwelcomechannel "github.com/defipod/mochi/pkg/repo/guild_config_welcome_channel"
	guildconfigxprole "github.com/defipod/mochi/pkg/repo/guild_config_xp_role"
	guildscheduledevent "github.com/defipod/mochi/pkg/repo/guild_scheduled_event"
	guilduseractivitylog "github.com/defipod/mochi/pkg/repo/guild_user_activity_log"
	guilduserxp "github.com/defipod/mochi/pkg/repo/guild_user_xp"
	guildusers "github.com/defipod/mochi/pkg/repo/guild_users"
	invitehistories "github.com/defipod/mochi/pkg/repo/invite_histories"
	kyberswapsupportedtokens "github.com/defipod/mochi/pkg/repo/kyberswap_supported_tokens"
	messagereaction "github.com/defipod/mochi/pkg/repo/message_reaction"
	messagereposthistory "github.com/defipod/mochi/pkg/repo/message_repost_history"
	mochinftsales "github.com/defipod/mochi/pkg/repo/mochi_nft_sales"
	monikerconfig "github.com/defipod/mochi/pkg/repo/moniker_config"
	nftaddrequesthistory "github.com/defipod/mochi/pkg/repo/nft_add_request_history"
	nftcollection "github.com/defipod/mochi/pkg/repo/nft_collection"
	nftsoulbound "github.com/defipod/mochi/pkg/repo/nft_soulbound"
	offchaintipbotactivitylogs "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_activity_logs"
	offchaintipbotchain "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_chain"
	offchaintipbotconfignotify "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_config_notify"
	offchaintipbotcontract "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_contract"
	offchaintipbotdepositlog "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_deposit_log"
	offchaintipbottokens "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_tokens"
	offchaintipbottransferhistories "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_transfer_histories"
	offchaintipbotuserbalancesnapshot "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_user_balance_snapshot"
	offchaintipbotuserbalances "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_user_balances"
	onchaintipbottransaction "github.com/defipod/mochi/pkg/repo/onchain_tip_bot_transaction"
	"github.com/defipod/mochi/pkg/repo/quest"
	questpass "github.com/defipod/mochi/pkg/repo/quest_pass"
	questreward "github.com/defipod/mochi/pkg/repo/quest_reward"
	questrewardtype "github.com/defipod/mochi/pkg/repo/quest_reward_type"
	queststreak "github.com/defipod/mochi/pkg/repo/quest_streak"
	questuserlist "github.com/defipod/mochi/pkg/repo/quest_user_list"
	questuserlog "github.com/defipod/mochi/pkg/repo/quest_user_log"
	questuserpass "github.com/defipod/mochi/pkg/repo/quest_user_pass"
	questuserreward "github.com/defipod/mochi/pkg/repo/quest_user_reward"
	salebotmarketplace "github.com/defipod/mochi/pkg/repo/sale_bot_marketplace"
	salebottwitterconfig "github.com/defipod/mochi/pkg/repo/sale_bot_twitter_config"
	token "github.com/defipod/mochi/pkg/repo/token"
	tradeoffer "github.com/defipod/mochi/pkg/repo/trade_offer"
	"github.com/defipod/mochi/pkg/repo/treasurer"
	treasurerrequest "github.com/defipod/mochi/pkg/repo/treasurer_request"
	treasurersubmission "github.com/defipod/mochi/pkg/repo/treasurer_submission"
	twitterpost "github.com/defipod/mochi/pkg/repo/twitter_post"
	twitterpoststreak "github.com/defipod/mochi/pkg/repo/twitter_post_streak"
	upvotestreaktier "github.com/defipod/mochi/pkg/repo/upvote_streak_tiers"
	userfeedback "github.com/defipod/mochi/pkg/repo/user_feedback"
	usernftbalance "github.com/defipod/mochi/pkg/repo/user_nft_balance"
	usernftwatchlistitem "github.com/defipod/mochi/pkg/repo/user_nft_watchlist_items"
	usersubmittedad "github.com/defipod/mochi/pkg/repo/user_submitted_ad"
	usertag "github.com/defipod/mochi/pkg/repo/user_tag"
	usertelegram "github.com/defipod/mochi/pkg/repo/user_telegram"
	usertelegramdiscordassociation "github.com/defipod/mochi/pkg/repo/user_telegram_discord_association"
	usertokenpricealert "github.com/defipod/mochi/pkg/repo/user_token_price_alert"
	usertokensupportrequest "github.com/defipod/mochi/pkg/repo/user_token_support_request"
	userwallet "github.com/defipod/mochi/pkg/repo/user_wallet"
	userwalletwatchlistitem "github.com/defipod/mochi/pkg/repo/user_wallet_watchlist_item"
	userwatchlistitem "github.com/defipod/mochi/pkg/repo/user_watchlist_item"
	users "github.com/defipod/mochi/pkg/repo/users"
	"github.com/defipod/mochi/pkg/repo/vault"
	vaultconfig "github.com/defipod/mochi/pkg/repo/vault_config"
	vaultinfo "github.com/defipod/mochi/pkg/repo/vault_info"
	vaulttransaction "github.com/defipod/mochi/pkg/repo/vault_transaction"
	walletsnapshot "github.com/defipod/mochi/pkg/repo/wallet_snapshot"
)

type Repo struct {
	DiscordUserGMStreak                  discordusergmstreak.Store
	DiscordUserUpvoteStreak              discorduserupvotestreak.Store
	GuildConfigGmGn                      guildconfiggmgn.Store
	GuildConfigSalesTracker              guildconfigsalestracker.Store
	GuildConfigWalletVerificationMessage guildconfigwalletverificationmessage.Store
	CommonwealthLatestData               commonwealthlastestdata.Store
	DiscordGuilds                        discordguilds.Store
	DiscordWalletVerification            discordwalletverification.Store
	InviteHistories                      invitehistories.Store
	Users                                users.Store
	UserWallet                           userwallet.Store
	GuildUsers                           guildusers.Store
	Token                                token.Store
	GuildConfigInviteTracker             guildconfiginvitetracker.Store
	UserSubmittedAd                      usersubmittedad.Store
	DiscordUserTokenAlert                discordusertokenalert.Store
	DiscordUserDevice                    discorduserdevice.Store
	GuildConfigWelcomeChannel            guildconfigwelcomechannel.Store
	GuildConfigReactionRole              guildconfigreactionrole.Store
	UserFeedback                         userfeedback.Store
	DiscordUserUpvoteLog                 discorduserupvotelog.Store
	GuildConfigDefaultRole               guildconfigdefaultrole.Store
	GuildConfigDefaultCollection         guildconfigdefaultcollection.Store
	GuildConfigPruneExclude              guildconfigpruneexclude.Store
	GuildConfigLevelUpMessage            guildconfiglevelupmessage.Store
	GuildConfigDaoTracker                guildconfigdaotracker.Store
	GuildConfigJoinLeaveChannel          guildconfigjoinleavechannel.Store
	GuildConfigRepostReaction            guildconfigrepostreaction.Store
	GuildConfigTwitterFeed               guildconfigtwitterfeed.Store
	GuildConfigVoteChannel               guildconfigvotechannel.Store
	GuildConfigDefaultCurrency           guildconfigdefaultcurrency.Store
	GuildConfigTwitterHashtag            guildconfigtwitterhashtag.Store
	GuildConfigTwitterBlacklist          guildconfigtwitterblacklist.Store
	DiscordGuildStats                    discordguildstats.Store
	DiscordGuildStatChannels             discordguildstatchannels.Store
	UpvoteStreakTier                     upvotestreaktier.Store
	GuildConfigToken                     guildconfigtoken.Store
	NFTCollection                        nftcollection.Store
	Activity                             activity.Store
	TwitterPost                          twitterpost.Store
	TwitterPostStreak                    twitterpoststreak.Store
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
	MochiNFTSales                        mochinftsales.Store
	GuildConfigDefaultTicker             guildconfigdefaultticker.Store
	UserWatchlistItem                    userwatchlistitem.Store
	GuildConfigGroupNFTRole              guildconfiggroupnftrole.Store
	CoingeckoSupportedTokens             coingeckosupportedtokens.Store
	UserTelegramDiscordAssociation       usertelegramdiscordassociation.Store
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
	OffchainTipBotChain                  offchaintipbotchain.Store
	OffchainTipBotContract               offchaintipbotcontract.Store
	TradeOffer                           tradeoffer.Store
	QuestStreak                          queststreak.Store
	OffchainTipBotUserBalances           offchaintipbotuserbalances.Store
	OffchainTipBotUserBalanceSnapshot    offchaintipbotuserbalancesnapshot.Store
	GuildBlacklistChannelRepostConfigs   guildblacklistchannelrepostconfigs.Store
	OffchainTipBotTokens                 offchaintipbottokens.Store
	OffchainTipBotActivityLogs           offchaintipbotactivitylogs.Store
	OffchainTipBotTransferHistories      offchaintipbottransferhistories.Store
	MonikerConfig                        monikerconfig.Store
	OffchainTipBotConfigNotify           offchaintipbotconfignotify.Store
	NftAddRequestHistory                 nftaddrequesthistory.Store
	OffchainTipBotDepositLog             offchaintipbotdepositlog.Store
	GuildConfigDaoProposal               guildconfigdaoproposal.Store
	DaoProposal                          daoproposal.Store
	DaoVote                              daovote.Store
	DaoProposalVoteOption                daoproposalvoteoption.Store
	DaoVoteOption                        daovoteoption.Store
	DaoGuidelineMessages                 daoguidelinemessages.Store
	OnchainTipBotTransaction             onchaintipbottransaction.Store
	GuildConfigTokenRole                 guildconfigtokenrole.Store
	Envelop                              envelop.Store
	GuildConfigXPRole                    guildconfigxprole.Store
	NftSoulbound                         nftsoulbound.Store
	SaleBotMarketplace                   salebotmarketplace.Store
	SaleBotTwitterConfig                 salebottwitterconfig.Store
	GuildConfigMixRole                   guildconfigmixrole.Store
	GuildConfigAdminRole                 guildconfigadminrole.Store
	UserWalletWatchlistItem              userwalletwatchlistitem.Store
	UserTokenPriceAlert                  usertokenpricealert.Store
	CommonwealthDiscussionSubscription   commonwealthdiscussionsubscription.Store
	UserTokenSupportRequest              usertokensupportrequest.Store
	UserTelegram                         usertelegram.Store
	Vault                                vault.Store
	VaultInfo                            vaultinfo.Store
	VaultConfig                          vaultconfig.Store
	Treasurer                            treasurer.Store
	TreasurerRequest                     treasurerrequest.Store
	KyberswapSupportedToken              kyberswapsupportedtokens.Store
	TreasurerSubmission                  treasurersubmission.Store
	VaultTransaction                     vaulttransaction.Store
	UserTag                              usertag.Store
	GuildConfigTipRange                  guildconfigtiprange.Store
	WalletSnapshot                       walletsnapshot.Store
}

package pg

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/repo"
	"github.com/defipod/mochi/pkg/repo/activity"
	ac "github.com/defipod/mochi/pkg/repo/airdrop_campaign"
	autoActionHistory "github.com/defipod/mochi/pkg/repo/auto_action_history"
	autoTrigger "github.com/defipod/mochi/pkg/repo/auto_trigger"
	"github.com/defipod/mochi/pkg/repo/chain"
	coingeckoinfo "github.com/defipod/mochi/pkg/repo/coingecko_info"
	coingeckosupportedtokens "github.com/defipod/mochi/pkg/repo/coingecko_supported_tokens"
	commonwealthdiscussionsubscription "github.com/defipod/mochi/pkg/repo/commonwealth_discussion_subscriptions"
	commonwealthlastestdata "github.com/defipod/mochi/pkg/repo/commonwealth_latest_data"
	configxplevel "github.com/defipod/mochi/pkg/repo/config_xp_level"
	"github.com/defipod/mochi/pkg/repo/content"
	daoguidelinemessages "github.com/defipod/mochi/pkg/repo/dao_guideline_messages"
	daoproposal "github.com/defipod/mochi/pkg/repo/dao_proposal"
	daoproposalvoteoption "github.com/defipod/mochi/pkg/repo/dao_proposal_vote_option"
	daovote "github.com/defipod/mochi/pkg/repo/dao_vote"
	daovoteoption "github.com/defipod/mochi/pkg/repo/dao_vote_option"
	discordguilds "github.com/defipod/mochi/pkg/repo/discord_guilds"
	discordusergmstreak "github.com/defipod/mochi/pkg/repo/discord_user_gm_streak"
	"github.com/defipod/mochi/pkg/repo/emojis"
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
	guildconfigjoinleavechannel "github.com/defipod/mochi/pkg/repo/guild_config_join_leave_channel"
	guildconfiglevelrole "github.com/defipod/mochi/pkg/repo/guild_config_level_role"
	guildconfiglevelupmessage "github.com/defipod/mochi/pkg/repo/guild_config_levelup_message"
	guildconfignftrole "github.com/defipod/mochi/pkg/repo/guild_config_nft_role"
	guildconfigreactionrole "github.com/defipod/mochi/pkg/repo/guild_config_reaction_roles"
	guildconfigsalestracker "github.com/defipod/mochi/pkg/repo/guild_config_sales_tracker"
	guildconfigtiprange "github.com/defipod/mochi/pkg/repo/guild_config_tip_range"
	guildconfigtoken "github.com/defipod/mochi/pkg/repo/guild_config_token"
	guildconfigtokenrole "github.com/defipod/mochi/pkg/repo/guild_config_token_role"
	guildconfigwalletverificationmessage "github.com/defipod/mochi/pkg/repo/guild_config_wallet_verification_message"
	guildconfigwelcomechannel "github.com/defipod/mochi/pkg/repo/guild_config_welcome_channel"
	guilduseractivitylog "github.com/defipod/mochi/pkg/repo/guild_user_activity_log"
	guilduserxp "github.com/defipod/mochi/pkg/repo/guild_user_xp"
	guildusers "github.com/defipod/mochi/pkg/repo/guild_users"
	kyberswapsupportedtokens "github.com/defipod/mochi/pkg/repo/kyberswap_supported_tokens"
	messagereaction "github.com/defipod/mochi/pkg/repo/message_reaction"
	messagereposthistory "github.com/defipod/mochi/pkg/repo/message_repost_history"
	mochinftsales "github.com/defipod/mochi/pkg/repo/mochi_nft_sales"
	monikerconfig "github.com/defipod/mochi/pkg/repo/moniker_config"
	nftaddrequesthistory "github.com/defipod/mochi/pkg/repo/nft_add_request_history"
	nftcollection "github.com/defipod/mochi/pkg/repo/nft_collection"
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
	pac "github.com/defipod/mochi/pkg/repo/profile_airdrop_campaign"
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
	"github.com/defipod/mochi/pkg/repo/token"
	"github.com/defipod/mochi/pkg/repo/treasurer"
	treasurerrequest "github.com/defipod/mochi/pkg/repo/treasurer_request"
	treasurersubmission "github.com/defipod/mochi/pkg/repo/treasurer_submission"
	userfeedback "github.com/defipod/mochi/pkg/repo/user_feedback"
	usernftbalance "github.com/defipod/mochi/pkg/repo/user_nft_balance"
	usernftwatchlistitem "github.com/defipod/mochi/pkg/repo/user_nft_watchlist_items"
	usersubmittedad "github.com/defipod/mochi/pkg/repo/user_submitted_ad"
	usertag "github.com/defipod/mochi/pkg/repo/user_tag"
	usertokenpricealert "github.com/defipod/mochi/pkg/repo/user_token_price_alert"
	usertokensupportrequest "github.com/defipod/mochi/pkg/repo/user_token_support_request"
	userwalletwatchlistitem "github.com/defipod/mochi/pkg/repo/user_wallet_watchlist_item"
	userwatchlistitem "github.com/defipod/mochi/pkg/repo/user_watchlist_item"
	"github.com/defipod/mochi/pkg/repo/users"
	"github.com/defipod/mochi/pkg/repo/vault"
	vaultconfig "github.com/defipod/mochi/pkg/repo/vault_config"
	vaultinfo "github.com/defipod/mochi/pkg/repo/vault_info"
	vaulttransaction "github.com/defipod/mochi/pkg/repo/vault_transaction"
	walletsnapshot "github.com/defipod/mochi/pkg/repo/wallet_snapshot"
)

// NewRepo new pg repo implementation
func NewRepo(db *gorm.DB) *repo.Repo {
	return &repo.Repo{
		Store:                                NewStore(db),
		DiscordGuilds:                        discordguilds.NewPG(db),
		Users:                                users.NewPG(db),
		GuildUsers:                           guildusers.NewPG(db),
		Token:                                token.NewPG(db),
		DiscordUserGMStreak:                  discordusergmstreak.NewPG(db),
		GuildConfigWelcomeChannel:            guildconfigwelcomechannel.NewPG(db),
		GuildConfigGmGn:                      guildconfiggmgn.NewPG(db),
		GuildConfigSalesTracker:              guildconfigsalestracker.NewPG(db),
		CommonwealthLatestData:               commonwealthlastestdata.NewPG(db),
		GuildConfigReactionRole:              guildconfigreactionrole.NewPG(db),
		GuildConfigDefaultCurrency:           guildconfigdefaultcurrency.NewPG(db),
		GuildConfigDaoTracker:                guildconfigdaotracker.NewPG(db),
		GuildConfigDefaultRole:               guildconfigdefaultrole.NewPG(db),
		GuildConfigJoinLeaveChannel:          guildconfigjoinleavechannel.NewPG(db),
		GuildConfigDefaultCollection:         guildconfigdefaultcollection.NewPG(db),
		GuildConfigWalletVerificationMessage: guildconfigwalletverificationmessage.NewPG(db),
		GuildConfigToken:                     guildconfigtoken.NewPG(db),
		NFTCollection:                        nftcollection.NewPG(db),
		UserSubmittedAd:                      usersubmittedad.NewPG(db),
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
		MochiNFTSales:                        mochinftsales.NewPG(db),
		GuildConfigDefaultTicker:             guildconfigdefaultticker.NewPG(db),
		UserWatchlistItem:                    userwatchlistitem.NewPG(db),
		GuildConfigGroupNFTRole:              guildconfiggroupnftrole.NewPG(db),
		CoingeckoSupportedTokens:             coingeckosupportedtokens.NewPG(db),
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
		OffchainTipBotChain:                  offchaintipbotchain.NewPG(db),
		OffchainTipBotContract:               offchaintipbotcontract.NewPG(db),
		OffchainTipBotUserBalances:           offchaintipbotuserbalances.NewPG(db),
		OffchainTipBotUserBalanceSnapshot:    offchaintipbotuserbalancesnapshot.NewPG(db),
		OffchainTipBotTokens:                 offchaintipbottokens.NewPG(db),
		OffchainTipBotActivityLogs:           offchaintipbotactivitylogs.NewPG(db),
		OffchainTipBotTransferHistories:      offchaintipbottransferhistories.NewPG(db),
		MonikerConfig:                        monikerconfig.NewPG(db),
		OffchainTipBotConfigNotify:           offchaintipbotconfignotify.NewPG(db),
		NftAddRequestHistory:                 nftaddrequesthistory.NewPG(db),
		OffchainTipBotDepositLog:             offchaintipbotdepositlog.NewPG(db),
		GuildConfigDaoProposal:               guildconfigdaoproposal.NewPG(db),
		DaoProposal:                          daoproposal.NewPG(db),
		DaoVote:                              daovote.NewPG(db),
		DaoProposalVoteOption:                daoproposalvoteoption.NewPG(db),
		DaoVoteOption:                        daovoteoption.NewPG(db),
		DaoGuidelineMessages:                 daoguidelinemessages.NewPG(db),
		OnchainTipBotTransaction:             onchaintipbottransaction.NewPG(db),
		GuildConfigTokenRole:                 guildconfigtokenrole.NewPG(db),
		GuildConfigLevelUpMessage:            guildconfiglevelupmessage.NewPG(db),
		Emojis:                               emojis.NewPG(db),
		SaleBotMarketplace:                   salebotmarketplace.NewPG(db),
		SaleBotTwitterConfig:                 salebottwitterconfig.NewPG(db),
		UserWalletWatchlistItem:              userwalletwatchlistitem.NewPG(db),
		UserTokenPriceAlert:                  usertokenpricealert.NewPG(db),
		CommonwealthDiscussionSubscription:   commonwealthdiscussionsubscription.NewPG(db),
		UserTokenSupportRequest:              usertokensupportrequest.NewPG(db),
		Vault:                                vault.NewPG(db),
		VaultInfo:                            vaultinfo.NewPG(db),
		VaultConfig:                          vaultconfig.NewPG(db),
		Treasurer:                            treasurer.NewPG(db),
		TreasurerRequest:                     treasurerrequest.NewPG(db),
		KyberswapSupportedToken:              kyberswapsupportedtokens.NewPG(db),
		TreasurerSubmission:                  treasurersubmission.NewPG(db),
		VaultTransaction:                     vaulttransaction.NewPG(db),
		UserTag:                              usertag.NewPG(db),
		GuildConfigTipRange:                  guildconfigtiprange.NewPG(db),
		GuildConfigAdminRole:                 guildconfigadminrole.NewPG(db),
		WalletSnapshot:                       walletsnapshot.NewPG(db),
		Content:                              content.NewPG(db),
		AirdropCampaign:                      ac.NewPG(db),
		ProfileAirdropCampaign:               pac.NewPG(db),
		AutoTrigger:                          autoTrigger.NewPG(db),
		AutoActionHistory:                    autoActionHistory.NewPG(db),
		CoingeckoInfo:                        coingeckoinfo.NewPG(db),
	}
}

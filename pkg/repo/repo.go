package repo

import (
	"github.com/defipod/mochi/pkg/repo/activity"
	ac "github.com/defipod/mochi/pkg/repo/airdrop_campaign"
	autoActionHistory "github.com/defipod/mochi/pkg/repo/auto_action_history"
	autoTrigger "github.com/defipod/mochi/pkg/repo/auto_trigger"
	"github.com/defipod/mochi/pkg/repo/chain"
	coingeckoinfo "github.com/defipod/mochi/pkg/repo/coingecko_info"
	coingeckosupportedtokens "github.com/defipod/mochi/pkg/repo/coingecko_supported_tokens"
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
	guildconfiglevelrole "github.com/defipod/mochi/pkg/repo/guild_config_level_role"
	guildconfiglogchannel "github.com/defipod/mochi/pkg/repo/guild_config_log_channel"
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
	offchaintipbotchain "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_chain"
	offchaintipbotconfignotify "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_config_notify"
	offchaintipbottokens "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_tokens"
	productbotcommand "github.com/defipod/mochi/pkg/repo/product_bot_command"
	productchangelogs "github.com/defipod/mochi/pkg/repo/product_changelogs"
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
	token "github.com/defipod/mochi/pkg/repo/token"
	tokeninfo "github.com/defipod/mochi/pkg/repo/token_info"
	userfeedback "github.com/defipod/mochi/pkg/repo/user_feedback"
	usernftbalance "github.com/defipod/mochi/pkg/repo/user_nft_balance"
	usernftwatchlistitem "github.com/defipod/mochi/pkg/repo/user_nft_watchlist_items"
	usertag "github.com/defipod/mochi/pkg/repo/user_tag"
	usertokenpricealert "github.com/defipod/mochi/pkg/repo/user_token_price_alert"
	usertokensupportrequest "github.com/defipod/mochi/pkg/repo/user_token_support_request"
	usertokenwatchlistitem "github.com/defipod/mochi/pkg/repo/user_token_watchlist_item"
	userwalletwatchlistitem "github.com/defipod/mochi/pkg/repo/user_wallet_watchlist_item"
	users "github.com/defipod/mochi/pkg/repo/users"
	"github.com/defipod/mochi/pkg/repo/vault"
	vaultconfig "github.com/defipod/mochi/pkg/repo/vault_config"
	vaultrequest "github.com/defipod/mochi/pkg/repo/vault_request"
	vaultsubmission "github.com/defipod/mochi/pkg/repo/vault_submission"
	vaulttransaction "github.com/defipod/mochi/pkg/repo/vault_transaction"
	vaulttreasurer "github.com/defipod/mochi/pkg/repo/vault_treasurer"
	walletsnapshot "github.com/defipod/mochi/pkg/repo/wallet_snapshot"
)

type Repo struct {
	Store                                Store
	DiscordUserGMStreak                  discordusergmstreak.Store
	GuildConfigGmGn                      guildconfiggmgn.Store
	GuildConfigSalesTracker              guildconfigsalestracker.Store
	GuildConfigWalletVerificationMessage guildconfigwalletverificationmessage.Store
	DiscordGuilds                        discordguilds.Store
	Users                                users.Store
	GuildUsers                           guildusers.Store
	Token                                token.Store
	GuildConfigWelcomeChannel            guildconfigwelcomechannel.Store
	GuildConfigReactionRole              guildconfigreactionrole.Store
	UserFeedback                         userfeedback.Store
	GuildConfigDefaultRole               guildconfigdefaultrole.Store
	GuildConfigDefaultCollection         guildconfigdefaultcollection.Store
	GuildConfigDefaultCurrency           guildconfigdefaultcurrency.Store
	GuildConfigToken                     guildconfigtoken.Store
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
	MochiNFTSales                        mochinftsales.Store
	GuildConfigDefaultTicker             guildconfigdefaultticker.Store
	UserTokenWatchlistItem               usertokenwatchlistitem.Store
	GuildConfigGroupNFTRole              guildconfiggroupnftrole.Store
	CoingeckoSupportedTokens             coingeckosupportedtokens.Store
	MessageReaction                      messagereaction.Store
	UserNftWatchlistItem                 usernftwatchlistitem.Store
	Quest                                quest.Store
	QuestRewardType                      questrewardtype.Store
	QuestUserLog                         questuserlog.Store
	QuestUserList                        questuserlist.Store
	QuestPass                            questpass.Store
	QuestUserPass                        questuserpass.Store
	QuestReward                          questreward.Store
	QuestUserReward                      questuserreward.Store
	OffchainTipBotChain                  offchaintipbotchain.Store
	QuestStreak                          queststreak.Store
	OffchainTipBotTokens                 offchaintipbottokens.Store
	MonikerConfig                        monikerconfig.Store
	OffchainTipBotConfigNotify           offchaintipbotconfignotify.Store
	NftAddRequestHistory                 nftaddrequesthistory.Store
	GuildConfigTokenRole                 guildconfigtokenrole.Store
	Emojis                               emojis.Store
	GuildConfigAdminRole                 guildconfigadminrole.Store
	UserWalletWatchlistItem              userwalletwatchlistitem.Store
	UserTokenPriceAlert                  usertokenpricealert.Store
	UserTokenSupportRequest              usertokensupportrequest.Store
	Vault                                vault.Store
	VaultConfig                          vaultconfig.Store
	VaultTreasurer                       vaulttreasurer.Store
	VaultRequest                         vaultrequest.Store
	KyberswapSupportedToken              kyberswapsupportedtokens.Store
	VaultSubmission                      vaultsubmission.Store
	VaultTransaction                     vaulttransaction.Store
	UserTag                              usertag.Store
	GuildConfigTipRange                  guildconfigtiprange.Store
	WalletSnapshot                       walletsnapshot.Store
	Content                              content.Store
	AirdropCampaign                      ac.Store
	ProfileAirdropCampaign               pac.Store
	AutoTrigger                          autoTrigger.Store
	AutoActionHistory                    autoActionHistory.Store
	CoingeckoInfo                        coingeckoinfo.Store
	GuildConfigLogChannel                guildconfiglogchannel.Store
	GuildConfigDaoTracker                guildconfigdaotracker.Store
	GuildConfigDaoProposal               guildconfigdaoproposal.Store
	DaoProposal                          daoproposal.Store
	DaoVote                              daovote.Store
	DaoProposalVoteOption                daoproposalvoteoption.Store
	DaoVoteOption                        daovoteoption.Store
	DaoGuidelineMessages                 daoguidelinemessages.Store
	TokenInfo                            tokeninfo.Store
	ProductBotCommand                    productbotcommand.Store
	ProductChangelogs                    productchangelogs.Store
}

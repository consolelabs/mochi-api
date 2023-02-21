package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/handler"
	"github.com/defipod/mochi/pkg/middleware"
)

// NewRoutes ...
func NewRoutes(r *gin.Engine, h *handler.Handler, cfg config.Config) {
	// asdsda
	v1 := r.Group("/api/v1")
	v1.Use(middleware.WithAuthContext(cfg))
	authGroup := v1.Group("/auth")
	{
		authGroup.POST("/login", h.Auth.Login)
		authGroup.POST("/logout", h.Auth.Logout)
	}
	cacheGroup := v1.Group("/cache")
	{
		cacheGroup.POST("/upvote", h.Cache.SetUpvoteMessageCache)
	}

	daoVoting := v1.Group("/dao-voting")
	{
		tokenHolderGroup := daoVoting.Group("/token-holder")
		{
			tokenHolderGroup.GET("status", h.DaoVoting.TokenHolderStatus)
		}
		proposalGroup := daoVoting.Group("/proposals")
		{
			proposalGroup.POST("/", h.DaoVoting.CreateProposal)
			proposalGroup.GET("", h.DaoVoting.GetProposals)
			proposalGroup.GET("/:proposal_id", h.DaoVoting.GetUserVotes)
			proposalGroup.DELETE("/:proposal_id", h.DaoVoting.DeteteProposal)
			voteGroup := proposalGroup.Group("/votes")
			{
				voteGroup.GET("", h.DaoVoting.GetVote)
				voteGroup.POST("", h.DaoVoting.CreateDaoVote)
				voteGroup.PUT("/:vote_id", h.DaoVoting.UpdateDaoVote)
			}
		}

	}
	dataGroup := v1.Group("/data")
	{
		dataGroup.GET("/metrics", h.Data.MetricByProperties)
		usageGroup := dataGroup.Group("/usage-stats")
		{
			usageGroup.GET("/gitbook", h.Data.AddGitbookClick)
			usageGroup.GET("/proposal", h.Data.MetricProposalUsage)
			usageGroup.GET("/dao-tracker", h.Data.MetricDaoTracker)
		}
		activitygroup := dataGroup.Group("/activities")
		{
			activitygroup.POST("/:activity", h.Config.ToggleActivityConfig)
		}
	}

	tipBotGroup := v1.Group("/tip")
	{
		// watch total balances
		tipBotGroup.GET("/total-balances", h.Tip.TotalBalances)
		tipBotGroup.GET("/total-offchain-balances", h.Tip.TotalOffchainBalances)
		tipBotGroup.GET("/total-fees", h.Tip.TotalFee)
		offchainTipBotTokensGroup := tipBotGroup.Group("/tokens")
		{
			offchainTipBotTokensGroup.GET("", h.Tip.GetAllTipBotTokens)
			offchainTipBotTokensGroup.PUT("", h.Tip.UpdateTokenFee)
		}
		// offchain tip bot
		tipBotGroup.GET("/chains", h.Tip.OffchainTipBotListAllChains)
		tipBotGroup.POST("/assign-contract", h.Tip.OffchainTipBotCreateAssignContract)
		tipBotGroup.GET("/balances", h.Tip.GetUserBalances)
		tipBotGroup.POST("/withdraw", h.Tip.OffchainTipBotWithdraw)
		tipBotGroup.POST("/transfer", h.Tip.TransferToken)
		tipBotGroup.GET("/transactions", h.User.GetTransactionsByQuery)
		tipBotGroup.GET("/history", h.Tip.GetTransactionHistoryByQuery)

		onchainGroup := tipBotGroup.Group("/onchain")
		{
			onchainGroup.POST("/submit", h.Tip.SubmitOnchainTransfer)
			onchainGroup.POST("/claim", h.Tip.ClaimOnchainTransfer)
			onchainGroup.GET("/:user_id/transfers", h.Tip.GetOnchainTransfers)
			onchainGroup.GET("/:user_id/balances", h.Tip.GetOnchainBalances)
		}
	}

	guildGroup := v1.Group("/guilds")
	{
		guildGroup.POST("", h.Guild.CreateGuild)
		guildGroup.GET("", h.Guild.GetGuilds)
		guildGroup.GET("/:guild_id", h.Guild.GetGuild)
		guildGroup.GET("/:guild_id/custom-tokens", h.ConfigDefi.ListAllCustomToken)
		guildGroup.GET("/user-managed", middleware.AuthGuard(cfg), h.Guild.ListMyGuilds)
		guildGroup.PUT("/:guild_id", h.Guild.UpdateGuild)
		countStatsGroup := guildGroup.Group("/:guild_id/stats")
		{
			countStatsGroup.GET("", h.Guild.GetGuildStatsHandler)
		}
		// api to contact with discord
		guildGroup.POST("/:guild_id/channels", h.Guild.CreateGuildChannel)
	}

	userGroup := v1.Group("/users")
	{
		userGroup.GET("me", middleware.AuthGuard(cfg), h.User.GetMyInfo)
		userGroup.POST("", h.User.IndexUsers)
		userGroup.GET("/:id", h.User.GetUser)
		userGroup.GET("/wallets/:address", h.User.GetUserWalletByGuildIDAddress)
		userGroup.GET("/gmstreak", h.User.GetUserCurrentGMStreak)
		userGroup.GET("/upvote-streak", h.User.GetUserCurrentUpvoteStreak) // get users upvote streak
		userGroup.GET("/upvote-leaderboard", h.User.GetUserUpvoteLeaderboard)
		userGroup.GET("/:id/transactions", h.User.GetUserTransaction)
		userGroup.GET("/top", h.User.GetTopUsers)
		userGroup.GET("/profiles", h.User.GetUserProfile)
		// moved to /widget/device, to be removed
		deviceGroup := userGroup.Group("/device")
		{
			deviceGroup.GET("", h.Widget.GetUserDevice)
			deviceGroup.POST("", h.Widget.UpsertUserDevice)
			deviceGroup.DELETE("", h.Widget.DeleteUserDevice)
		}
		userGroup.POST("/xp", h.User.SendUserXP)
		userGroup.POST("/envelop", h.User.CreateEnvelop)
		userGroup.GET("/:id/envelop-streak", h.User.GetUserEnvelopStreak) // get users upvote streak

		walletsGroup := userGroup.Group("/:id/wallets")
		{
			walletsGroup.GET("", h.Wallet.List)
			walletsGroup.POST("/track", h.Wallet.Track)
			walletsGroup.POST("/untrack", h.Wallet.Untrack)
			walletsGroup.GET("/:address", h.Wallet.GetOne)
			walletsGroup.GET("/:address/:type/assets", h.Wallet.ListAssets)
			walletsGroup.GET("/:address/:type/txns", h.Wallet.ListTransactions)
		}
	}

	communityGroup := v1.Group("/community")
	{
		invitesGroup := communityGroup.Group("/invites")
		{
			invitesGroup.GET("/", h.User.GetInvites)
			invitesGroup.GET("/config", h.ConfigChannel.GetInviteTrackerConfig)
			invitesGroup.POST("/config", h.ConfigChannel.ConfigureInvites)
			invitesGroup.GET("/leaderboard/:id", h.User.GetInvitesLeaderboard)
			invitesGroup.GET("/aggregation", h.User.InvitesAggregation)
		}
		feedbackGroup := communityGroup.Group("/feedback")
		{
			feedbackGroup.POST("", h.Community.HandleUserFeedback)
			feedbackGroup.PUT("", h.Community.UpdateUserFeedback)
			feedbackGroup.GET("", h.Community.GetAllUserFeedback)
		}
		questGroup := communityGroup.Group("/quests")
		{
			questGroup.GET("", h.Community.GetUserQuestList)
			questGroup.POST("/progress", h.Community.UpdateQuestProgress)
			questGroup.POST("/claim", h.Community.ClaimQuestsRewards)
		}
		twitterGroup := communityGroup.Group("/twitter")
		{
			twitterGroup.POST("", h.Community.CreateTwitterPost)
			twitterGroup.GET("/top", h.Community.GetTwitterLeaderboard)
		}
		// starboard
		repostReactionGroup := communityGroup.Group("/repost-reactions")
		{
			repostReactionGroup.GET("/:guild_id", h.Community.GetRepostReactionConfigs)
			repostReactionGroup.POST("", h.Community.ConfigRepostReaction)
			repostReactionGroup.DELETE("", h.Community.RemoveRepostReactionConfig)
			repostReactionGroup.POST("/conversation", h.Community.CreateConfigRepostReactionConversation)
			repostReactionGroup.DELETE("/conversation", h.Community.RemoveConfigRepostReactionConversation)
			repostReactionGroup.PUT("/message-repost", h.Community.EditMessageRepost)
			repostReactionGroup.POST("/blacklist-channel", h.Community.CreateBlacklistChannelRepostConfig)
			repostReactionGroup.GET("/blacklist-channel", h.Community.GetGuildBlacklistChannelRepostConfig)
			repostReactionGroup.DELETE("/blacklist-channel", h.Community.DeleteBlacklistChannelRepostConfig)
		}
		levelupGroup := communityGroup.Group("/levelup")
		{
			levelupGroup.GET("", h.Community.GetLevelUpMessage)
			levelupGroup.POST("", h.Community.UpsertLevelUpMessage)
			levelupGroup.DELETE("", h.Community.DeleteLevelUpMessage)
		}
	}

	configGroup := v1.Group("/configs")
	{
		//
		configGroup.GET("/upvote-tiers", h.ConfigChannel.GetUpvoteTiersConfig)
		configGroup.GET("/sales-tracker", h.ConfigChannel.GetSalesTrackerConfig)
		configGroup.POST("/sales-tracker", h.ConfigChannel.CreateSalesTrackerConfig)
		// prune exclude
		configGroup.GET("/whitelist-prune", h.Config.GetGuildPruneExclude)
		configGroup.POST("/whitelist-prune", h.Config.UpsertGuildPruneExclude)
		configGroup.DELETE("/whitelist-prune", h.Config.DeleteGuildPruneExclude)
		// moved to /widget/token-alert, to be removed
		tokenAlertGroup := configGroup.Group("/token-alert")
		{
			tokenAlertGroup.GET("", h.Widget.GetUserTokenAlert)
			tokenAlertGroup.POST("", h.Widget.UpsertUserTokenAlert)
			tokenAlertGroup.DELETE("", h.Widget.DeleteUserTokenAlert)
		}
		configTwitterSaleGroup := configGroup.Group("/twitter-sales")
		{
			configTwitterSaleGroup.GET("", h.ConfigTwitterSale.Get)
			configTwitterSaleGroup.POST("", h.ConfigTwitterSale.Create)
		}
	}

	configChannelGroup := v1.Group("/config-channels")
	{
		configChannelGroup.GET("/gm", h.ConfigChannel.GetGmConfig)
		configChannelGroup.POST("/gm", h.ConfigChannel.UpsertGmConfig)
		// config welcome channel
		configChannelGroup.GET("/welcome", h.ConfigChannel.GetWelcomeChannelConfig)
		configChannelGroup.POST("/welcome", h.ConfigChannel.UpsertWelcomeChannelConfig)
		configChannelGroup.DELETE("/welcome", h.ConfigChannel.DeleteWelcomeChannelConfig)
		// config vote channel
		configChannelGroup.GET("/upvote", h.ConfigChannel.GetVoteChannelConfig)
		configChannelGroup.POST("/upvote", h.ConfigChannel.UpsertVoteChannelConfig)
		configChannelGroup.DELETE("/upvote", h.ConfigChannel.DeleteVoteChannelConfig)
		// config tip notify channel
		configChannelGroup.POST("/tip-notify", h.ConfigChannel.CreateConfigNotify)
		configChannelGroup.GET("/tip-notify", h.ConfigChannel.ListConfigNotify)
		configChannelGroup.DELETE("/tip-notify/:id", h.ConfigChannel.DeleteConfigNotify)
		// config join-leave channel
		configChannelGroup.GET("/join-leave", h.ConfigChannel.GetJoinLeaveChannelConfig)
		configChannelGroup.POST("/join-leave", h.ConfigChannel.UpsertJoinLeaveChannelConfig)
		configChannelGroup.DELETE("/join-leave", h.ConfigChannel.DeleteJoinLeaveChannelConfig)
		// config dao proposal channel
		configChannelGroup.GET("/:guild_id/proposal", h.ConfigChannel.GetGuildConfigDaoProposal)
		configChannelGroup.DELETE("/proposal", h.ConfigChannel.DeleteGuildConfigDaoProposal)
		configChannelGroup.POST("/proposal", h.ConfigChannel.CreateProposalChannelConfig)
		// config dao tracker channel
		configChannelGroup.GET("/dao-tracker/:guild_id", h.ConfigChannel.GetGuildConfigDaoTracker)
		configChannelGroup.DELETE("/dao-tracker", h.ConfigChannel.DeleteGuildConfigDaoTracker)
		configChannelGroup.POST("/dao-tracker", h.ConfigChannel.UpsertGuildConfigDaoTracker)
	}

	configRoleGroup := v1.Group("/config-roles")
	{
		roleReactionGroup := configRoleGroup.Group("/reaction-roles")
		{
			roleReactionGroup.GET("", h.ConfigRoles.GetAllRoleReactionConfigs)
			roleReactionGroup.POST("", h.ConfigRoles.AddReactionRoleConfig)
			roleReactionGroup.DELETE("", h.ConfigRoles.RemoveReactionRoleConfig)
			roleReactionGroup.POST("/filter", h.ConfigRoles.FilterConfigByReaction)

		}
		defaultRoleGroup := configRoleGroup.Group("/default-roles")
		{
			defaultRoleGroup.GET("", h.ConfigRoles.GetDefaultRolesByGuildID)
			defaultRoleGroup.POST("", h.ConfigRoles.CreateDefaultRole)
			defaultRoleGroup.DELETE("", h.ConfigRoles.DeleteDefaultRoleByGuildID)
		}
		levelRoleGroup := configRoleGroup.Group("/level-roles")
		{
			levelRoleGroup.POST("", h.ConfigRoles.ConfigLevelRole)
			levelRoleGroup.GET("/:guild_id", h.ConfigRoles.GetLevelRoleConfigs)
			levelRoleGroup.DELETE("/:guild_id", h.ConfigRoles.RemoveLevelRoleConfig)
		}
		nftRoleGroup := configRoleGroup.Group("/nft-roles")
		{
			nftRoleGroup.GET("", h.ConfigRoles.ListGuildGroupNFTRoles)
			nftRoleGroup.POST("", h.ConfigRoles.NewGuildGroupNFTRole)
			nftRoleGroup.DELETE("/group", h.ConfigRoles.RemoveGuildGroupNFTRole)
			nftRoleGroup.DELETE("/", h.ConfigRoles.RemoveGuildNFTRole)
		}
		tokenRoleGroup := configRoleGroup.Group("/token-roles")
		{
			tokenRoleGroup.POST("", h.ConfigRoles.CreateGuildTokenRole)
			tokenRoleGroup.GET(":guild_id", h.ConfigRoles.ListGuildTokenRoles)
			tokenRoleGroup.PUT("/:id", h.ConfigRoles.UpdateGuildTokenRole)
			tokenRoleGroup.DELETE("/:id", h.ConfigRoles.RemoveGuildTokenRole)
		}
		xpRoleGroup := configRoleGroup.Group("/xp-roles")
		{
			xpRoleGroup.POST("", h.ConfigRoles.CreateGuildXPRole)
			xpRoleGroup.GET("", h.ConfigRoles.ListGuildXPRoles)
			xpRoleGroup.DELETE("/:id", h.ConfigRoles.RemoveGuildXPRole)
		}
		mixRoleGroup := configRoleGroup.Group("/mix-roles")
		{
			mixRoleGroup.POST("", h.ConfigRoles.CreateGuildMixRole)
			mixRoleGroup.GET("", h.ConfigRoles.ListGuildMixRoles)
			mixRoleGroup.DELETE("/:id", h.ConfigRoles.RemoveGuildMixRole)
		}
	}

	configCommunityGroup := v1.Group("/config-community")
	{
		telegramGroup := configCommunityGroup.Group("/telegram")
		{
			telegramGroup.GET("", h.ConfigCommunity.GetLinkedTelegram)
			telegramGroup.POST("", h.ConfigCommunity.LinkUserTelegramWithDiscord)
		}
		twitterGroup := configCommunityGroup.Group("/twitter")
		{
			twitterGroup.POST("", h.ConfigCommunity.CreateTwitterConfig)
			twitterGroup.GET("", h.ConfigCommunity.GetAllTwitterConfig)
			twitterGroup.GET("/hashtag/:guild_id", h.ConfigCommunity.GetTwitterHashtagConfig)
			twitterGroup.DELETE("/hashtag/:guild_id", h.ConfigCommunity.DeleteTwitterHashtagConfig)
			twitterGroup.POST("/hashtag", h.ConfigCommunity.CreateTwitterHashtagConfig)
			twitterGroup.GET("/hashtag", h.ConfigCommunity.GetAllTwitterHashtagConfig)
			twitterGroup.POST("/blacklist", h.ConfigCommunity.AddToTwitterBlackList)
			twitterGroup.GET("/blacklist", h.ConfigCommunity.GetTwitterBlackList)
			twitterGroup.DELETE("/blacklist", h.ConfigCommunity.DeleteFromTwitterBlackList)
		}
	}

	configDefiGroup := v1.Group("/config-defi")
	{
		defaultCurrencyGroup := configDefiGroup.Group("/default-currency")
		{
			defaultCurrencyGroup.GET("", h.ConfigDefi.GetGuildDefaultCurrency)
			defaultCurrencyGroup.POST("", h.ConfigDefi.UpsertGuildDefaultCurrency)
			defaultCurrencyGroup.DELETE("", h.ConfigDefi.DeleteGuildDefaultCurrency)
		}
		defaultSymbolGroup := configDefiGroup.Group("/default-symbol")
		{
			defaultSymbolGroup.POST("", h.ConfigDefi.CreateDefaultCollectionSymbol)
		}
		tokenGroup := configDefiGroup.Group("/tokens")
		{
			tokenGroup.GET("", h.ConfigDefi.GetGuildTokens)
			tokenGroup.POST("", h.ConfigDefi.UpsertGuildTokenConfig)
			tokenGroup.GET("/default", h.ConfigDefi.GetDefaultToken)
			tokenGroup.POST("/default", h.ConfigDefi.ConfigDefaultToken)
			tokenGroup.DELETE("/default", h.ConfigDefi.RemoveDefaultToken)
		}
		customTokenGroup := configDefiGroup.Group("/custom-tokens")
		{
			customTokenGroup.POST("", h.ConfigDefi.HandlerGuildCustomTokenConfig)
		}

		defaultTickerGroup := configDefiGroup.Group("/default-ticker")
		{
			defaultTickerGroup.GET("", h.ConfigDefi.GetGuildDefaultTicker)
			defaultTickerGroup.POST("", h.ConfigDefi.SetGuildDefaultTicker)
		}
		monikerGroup := configDefiGroup.Group("/monikers")
		{
			monikerGroup.POST("", h.ConfigDefi.UpsertMonikerConfig)
			monikerGroup.GET("/:guild_id", h.ConfigDefi.GetMonikerByGuildID)
			monikerGroup.DELETE("", h.ConfigDefi.DeleteMonikerConfig)
			monikerGroup.GET("/default", h.ConfigDefi.GetDefaultMoniker)
		}
	}

	defiGroup := v1.Group("/defi")
	{
		defiGroup.GET("")
		defiGroup.GET("/tokens", h.Defi.GetSupportedTokens)
		defiGroup.GET("/token", h.Defi.GetSupportedToken)

		// Data from CoinGecko
		defiGroup.GET("/market-chart", h.Defi.GetHistoricalMarketChart)
		defiGroup.GET("/coins/:id", h.Defi.GetCoin)
		defiGroup.GET("/coins", h.Defi.SearchCoins)
		defiGroup.GET("/coins/compare", h.Defi.CompareToken)
		defiGroup.GET("/chains", h.Defi.ListAllChain)

		watchlistGroup := defiGroup.Group("/watchlist")
		{
			watchlistGroup.GET("", h.Defi.GetUserWatchlist)
			watchlistGroup.POST("", h.Defi.AddToWatchlist)
			watchlistGroup.DELETE("", h.Defi.RemoveFromWatchlist)
		}

		priceAlertGroup := defiGroup.Group("/price-alert")
		{
			priceAlertGroup.GET("", h.Defi.GetUserListPriceAlert)
			priceAlertGroup.POST("", h.Defi.AddTokenPriceAlert)
			priceAlertGroup.DELETE("", h.Defi.RemoveTokenPriceAlert)
		}
	}

	verifyGroup := v1.Group("/verify")
	{
		verifyGroup.POST("/config", h.Verify.NewGuildConfigWalletVerificationMessage)
		verifyGroup.GET("/config/:guild_id", h.Verify.GetGuildConfigWalletVerificationMessage)
		verifyGroup.PUT("/config", h.Verify.UpdateGuildConfigWalletVerificationMessage)
		verifyGroup.DELETE("/config", h.Verify.DeleteGuildConfigWalletVerificationMessage)
		verifyGroup.POST("/generate", h.Verify.GenerateVerification)
		verifyGroup.POST("", h.Verify.VerifyWalletAddress)
	}

	nftsGroup := v1.Group("/nfts")
	{
		nftsGroup.GET("", h.Nft.ListAllNFTCollections)
		nftsGroup.GET("/tickers", h.Nft.GetNftTokenTickers)
		nftsGroup.GET("/:symbol/:id", h.Nft.GetNFTDetail)
		nftsGroup.GET("/:symbol/:id/activity", h.Nft.GetNFTActivity)
		nftsGroup.GET("/supported-chains", h.Nft.GetSupportedChains)
		nftsGroup.GET("/trading-volume", h.Nft.GetNFTTradingVolume)
		nftsGroup.GET("/sales", h.Nft.GetNftSalesHandler)
		nftsGroup.GET("/new-listed", h.Nft.GetNewListedNFTCollection)
		nftsGroup.GET("/icons", h.Nft.GetNftMetadataAttrIcon)
		collectionsGroup := nftsGroup.Group("/collections")
		{
			collectionsGroup.GET("/:symbol/detail", h.Nft.GetDetailNftCollection)
			collectionsGroup.GET("/stats", h.Nft.GetCollectionCount)
			collectionsGroup.GET("", h.Nft.GetNFTCollections)
			collectionsGroup.GET("/suggestion", h.Nft.GetSuggestionNFTCollections)
			collectionsGroup.POST("", h.Nft.CreateNFTCollection)
			collectionsGroup.PATCH("/:address", h.Nft.UpdateNFTCollection) //to update collection images, delete after use
			collectionsGroup.GET("/:symbol", h.Nft.GetNFTTokens)
			collectionsGroup.GET("/tickers", h.Nft.GetNFTCollectionTickers)
			collectionsGroup.GET("/address/:address", h.Nft.GetNFTCollectionByAddressChain)
		}
		nftWatchlistGroup := nftsGroup.Group("/watchlist")
		{
			nftWatchlistGroup.GET("", h.Nft.GetNftWatchlist)
			nftWatchlistGroup.POST("", h.Nft.AddNftWatchlist)
			nftWatchlistGroup.DELETE("", h.Nft.DeleteNftWatchlist)
		}
		trade := nftsGroup.Group("/trades")
		{
			trade.GET("/:id", h.Nft.GetTradeOffer)
			trade.POST("", h.Nft.CreateTradeOffer)
		}
		defaultNftTickerGroup := nftsGroup.Group("/default-nft-ticker")
		{
			defaultNftTickerGroup.GET("", h.Nft.GetGuildDefaultNftTicker)
			defaultNftTickerGroup.POST("", h.Nft.SetGuildDefaultNftTicker)
		}
		soulbound := nftsGroup.Group("/soulbound")
		{
			soulbound.GET("", h.Nft.GetSoulboundNFT)
			soulbound.POST("", h.Nft.EnrichSoulboundNFT)
		}
	}

	fiatGroup := v1.Group("/fiat")
	{
		fiatGroup.GET("/historical-exchange-rates", h.Defi.GetFiatHistoricalExchangeRates)
	}

	widgetGroup := v1.Group("/widget")
	{
		tokenAlertGroup := widgetGroup.Group("/token-alert")
		{
			tokenAlertGroup.GET("", h.Widget.GetUserTokenAlert)
			tokenAlertGroup.POST("", h.Widget.UpsertUserTokenAlert)
			tokenAlertGroup.DELETE("", h.Widget.DeleteUserTokenAlert)
		}
		deviceGroup := widgetGroup.Group("/device")
		{
			deviceGroup.GET("", h.Widget.GetUserDevice)
			deviceGroup.POST("", h.Widget.UpsertUserDevice)
			deviceGroup.DELETE("", h.Widget.DeleteUserDevice)
		}
	}
	webhook := v1.Group("/webhook")
	{
		webhook.POST("/discord", h.Webhook.HandleDiscordWebhook)
		webhook.POST("/nft", h.Webhook.WebhookNftHandler)
		webhook.POST("/topgg", h.Webhook.WebhookUpvoteTopGG)
		webhook.POST("/discordbotlist", h.Webhook.WebhookUpvoteDiscordBot)
		webhook.POST("/snapshot", h.Webhook.WebhookSnapshotProposal)
	}
	dataWebhookGroup := v1.Group("/data-webhook")
	{
		dataWebhookGroup.POST("/notify-nft-integration", h.Webhook.NotifyNftCollectionIntegration)
		dataWebhookGroup.POST("/notify-nft-add", h.Webhook.NotifyNftCollectionAdd)
		dataWebhookGroup.POST("/notify-nft-sync", h.Webhook.NotifyNftCollectionSync)
		dataWebhookGroup.POST("/notify-sale-marketplace", h.Webhook.NotifySaleMarketplace)
	}
}

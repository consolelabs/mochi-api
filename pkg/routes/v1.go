package routes

import (
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/handler"
	"github.com/defipod/mochi/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// NewRoutes ...
func NewRoutes(r *gin.Engine, h *handler.Handler, cfg config.Config) {
	// asdsda
	v1 := r.Group("/api/v1")
	v1.Use(middleware.WithAuthContext(cfg))

	chainGroup := v1.Group("/chains")
	{
		chainGroup.GET("", h.ListAllChain)
	}

	offchainTipBotGroup := v1.Group("/offchain-tip-bot")
	{
		offchainTipBotGroup.GET("/chains", h.OffchainTipBotListAllChains)
		offchainTipBotGroup.POST("/assign-contract", h.OffchainTipBotCreateAssignContract)
		offchainTipBotGroup.GET("/balances", h.GetUserBalances)
		offchainTipBotGroup.POST("/withdraw", h.OffchainTipBotWithdraw)
		offchainTipBotGroup.POST("/transfer", h.TransferToken)
	}

	trade := v1.Group("/trades")
	{
		trade.GET("/:id", h.GetTradeOffer)
		trade.POST("", h.CreateTradeOffer)
	}

	authGroup := v1.Group("/auth")
	{
		authGroup.POST("/login", h.Login)
		authGroup.POST("/logout", h.Logout)
	}

	guildGroup := v1.Group("/guilds")
	{
		guildGroup.POST("", h.CreateGuild)
		guildGroup.GET("", h.GetGuilds)
		guildGroup.GET("/:guild_id", h.GetGuild)
		guildGroup.GET("/:guild_id/custom-tokens", h.ListAllCustomToken)
		guildGroup.GET("/user-managed", middleware.AuthGuard(cfg), h.ListMyGuilds)
		guildGroup.PUT("/:guild_id", h.UpdateGuild)

		customCommandGroup := guildGroup.Group("/:guild_id/custom-commands")
		{
			customCommandGroup.POST("", h.CreateCustomCommand)
			customCommandGroup.GET("", h.ListCustomCommands)
			customCommandGroup.GET("/:command_id", h.GetCustomCommand)
			customCommandGroup.PUT("/:command_id", h.UpdateCustomCommand)
			customCommandGroup.DELETE("/:command_id", h.DeleteCustomCommand)
		}

		countStatsGroup := guildGroup.Group("/:guild_id/stats")
		{
			countStatsGroup.GET("", h.GetGuildStatsHandler)
		}

		// api to contact with discord
		guildGroup.POST("/:guild_id/channels", h.CreateGuildChannel)
	}

	userGroup := v1.Group("/users")
	{
		userGroup.GET("me", middleware.AuthGuard(cfg), h.GetMyInfo)
		userGroup.POST("", h.IndexUsers)
		userGroup.GET("/:id", h.GetUser)
		userGroup.GET("/wallets/:address", h.GetUserWalletByGuildIDAddress)
		userGroup.GET("/gmstreak", h.GetUserCurrentGMStreak)
		userGroup.GET("/upvote-streak", h.GetUserCurrentUpvoteStreak) // get users upvote streak
		userGroup.GET("/upvote-leaderboard", h.GetUserUpvoteLeaderboard)
		userGroup.GET("/:id/transactions", h.GetUserTransaction)
		userGroup.GET("/top", h.GetTopUsers)
		deviceGroup := userGroup.Group("/device")
		{
			deviceGroup.GET("", h.GetUserDevice)
			deviceGroup.POST("", h.UpsertUserDevice)
			deviceGroup.DELETE("", h.DeleteUserDevice)
		}
	}

	communityGroup := v1.Group("/community")
	{
		invitesGroup := communityGroup.Group("/invites")
		{
			invitesGroup.GET("/", h.GetInvites)
			invitesGroup.GET("/config", h.GetInviteTrackerConfig)
			invitesGroup.POST("/config", h.ConfigureInvites)
			invitesGroup.GET("/leaderboard/:id", h.GetInvitesLeaderboard)
			invitesGroup.GET("/aggregation", h.InvitesAggregation)
		}
	}

	profleGroup := v1.Group("/profiles")
	{
		profleGroup.GET("", h.GetUserProfile)
	}

	configGroup := v1.Group("/configs")
	{
		configGroup.GET("")
		configGroup.GET("/gm", h.GetGmConfig)
		configGroup.POST("/gm", h.UpsertGmConfig)
		// config welcome channel
		configGroup.GET("/welcome", h.GetWelcomeChannelConfig)
		configGroup.POST("/welcome", h.UpsertWelcomeChannelConfig)
		configGroup.DELETE("/welcome", h.DeleteWelcomeChannelConfig)
		// config vote channel
		configGroup.GET("/upvote", h.GetVoteChannelConfig)
		configGroup.POST("/upvote", h.UpsertVoteChannelConfig)
		configGroup.DELETE("/upvote", h.DeleteVoteChannelConfig)
		//
		configGroup.GET("/upvote-tiers", h.GetUpvoteTiersConfig)
		configGroup.GET("/sales-tracker", h.GetSalesTrackerConfig)
		// prune exclude
		configGroup.GET("/whitelist-prune", h.GetGuildPruneExclude)
		configGroup.POST("/whitelist-prune", h.UpsertGuildPruneExclude)
		configGroup.DELETE("/whitelist-prune", h.DeleteGuildPruneExclude)
		// config join-leave channel
		configGroup.GET("/join-leave", h.GetJoinLeaveChannelConfig)
		configGroup.POST("/join-leave", h.UpsertJoinLeaveChannelConfig)
		configGroup.DELETE("/join-leave", h.DeleteJoinLeaveChannelConfig)
		roleReactionGroup := configGroup.Group("/reaction-roles")
		{
			roleReactionGroup.GET("", h.GetAllRoleReactionConfigs)
			roleReactionGroup.POST("", h.AddReactionRoleConfig)
			roleReactionGroup.DELETE("", h.RemoveReactionRoleConfig)
			roleReactionGroup.POST("/filter", h.FilterConfigByReaction)

		}
		defaultRoleGroup := configGroup.Group("/default-roles")
		{
			defaultRoleGroup.GET("", h.GetDefaultRolesByGuildID)
			defaultRoleGroup.POST("", h.CreateDefaultRole)
			defaultRoleGroup.DELETE("", h.DeleteDefaultRoleByGuildID)
		}
		defaultSymbolGroup := configGroup.Group("/default-symbol")
		{
			defaultSymbolGroup.POST("", h.CreateDefaultCollectionSymbol)
		}
		tokenGroup := configGroup.Group("/tokens")
		{
			tokenGroup.GET("", h.GetGuildTokens)
			tokenGroup.POST("", h.UpsertGuildTokenConfig)
			tokenGroup.GET("/default", h.GetDefaultToken)
			tokenGroup.POST("/default", h.ConfigDefaultToken)
			tokenGroup.DELETE("/default", h.RemoveDefaultToken)
		}
		customTokenGroup := configGroup.Group("/custom-tokens")
		{
			customTokenGroup.POST("", h.HandlerGuildCustomTokenConfig)
		}
		levelRoleGroup := configGroup.Group("/level-roles")
		{
			levelRoleGroup.POST("", h.ConfigLevelRole)
			levelRoleGroup.GET("/:guild_id", h.GetLevelRoleConfigs)
			levelRoleGroup.DELETE("/:guild_id", h.RemoveLevelRoleConfig)
		}
		nftRoleGroup := configGroup.Group("/nft-roles")
		{
			nftRoleGroup.GET("", h.ListGuildGroupNFTRoles)
			nftRoleGroup.POST("", h.NewGuildGroupNFTRole)
			nftRoleGroup.DELETE("/group", h.RemoveGuildGroupNFTRole)
			nftRoleGroup.DELETE("/", h.RemoveGuildNFTRole)
		}
		repostReactionGroup := configGroup.Group("/repost-reactions")
		{
			repostReactionGroup.GET("/:guild_id", h.GetRepostReactionConfigs)
			repostReactionGroup.POST("", h.ConfigRepostReaction)
			repostReactionGroup.DELETE("", h.RemoveRepostReactionConfig)
			repostReactionGroup.POST("/conversation", h.CreateConfigRepostReactionConversation)
			repostReactionGroup.DELETE("/conversation", h.RemoveConfigRepostReactionConversation)
			repostReactionGroup.PUT("/message-repost", h.EditMessageRepost)
			repostReactionGroup.POST("/blacklist-channel", h.CreateBlacklistChannelRepostConfig)
			repostReactionGroup.GET("/blacklist-channel", h.GetGuildBlacklistChannelRepostConfig)
			repostReactionGroup.DELETE("/blacklist-channel", h.DeleteBlacklistChannelRepostConfig)
		}
		activitygroup := configGroup.Group("/activities")
		{
			activitygroup.POST("/:activity", h.ToggleActivityConfig)
		}
		twitterGroup := configGroup.Group("/twitter")
		{
			twitterGroup.POST("", h.CreateTwitterConfig)
			twitterGroup.GET("", h.GetAllTwitterConfig)
			twitterGroup.GET("/hashtag/:guild_id", h.GetTwitterHashtagConfig)
			twitterGroup.DELETE("/hashtag/:guild_id", h.DeleteTwitterHashtagConfig)
			twitterGroup.POST("/hashtag", h.CreateTwitterHashtagConfig)
			twitterGroup.GET("/hashtag", h.GetAllTwitterHashtagConfig)
			twitterGroup.POST("/blacklist", h.AddToTwitterBlackList)
			twitterGroup.GET("/blacklist", h.GetTwitterBlackList)
			twitterGroup.DELETE("/blacklist", h.DeleteFromTwitterBlackList)
		}
		defaultTickerGroup := configGroup.Group("/default-ticker")
		{
			defaultTickerGroup.GET("", h.GetGuildDefaultTicker)
			defaultTickerGroup.POST("", h.SetGuildDefaultTicker)
		}

		defaultNftTickerGroup := configGroup.Group("/default-nft-ticker")
		{
			defaultNftTickerGroup.GET("", h.GetGuildDefaultNftTicker)
			defaultNftTickerGroup.POST("", h.SetGuildDefaultNftTicker)
		}

		telegramGroup := configGroup.Group("/telegram")
		{
			telegramGroup.GET("", h.GetLinkedTelegram)
			telegramGroup.POST("", h.LinkUserTelegramWithDiscord)
		}
		tokenAlertGroup := configGroup.Group("/token-alert")
		{
			tokenAlertGroup.GET("", h.GetUserTokenAlert)
			tokenAlertGroup.POST("", h.UpsertUserTokenAlert)
			tokenAlertGroup.DELETE("", h.DeleteUserTokenAlert)
		}
	}

	defiGroup := v1.Group("/defi")
	{
		defiGroup.GET("")
		defiGroup.POST("/transfer", h.InDiscordWalletTransfer)
		defiGroup.POST("/withdraw", h.InDiscordWalletWithdraw)
		defiGroup.GET("/balances", h.InDiscordWalletBalances)
		defiGroup.GET("/tokens", h.GetSupportedTokens)

		// Data from CoinGecko
		defiGroup.GET("/market-chart", h.GetHistoricalMarketChart)
		defiGroup.GET("/coins/:id", h.GetCoin)
		defiGroup.GET("/coins", h.SearchCoins)
		defiGroup.GET("/coins/compare", h.CompareToken)

		watchlistGroup := defiGroup.Group("/watchlist")
		{
			watchlistGroup.GET("", h.GetUserWatchlist)
			watchlistGroup.POST("", h.AddToWatchlist)
			watchlistGroup.DELETE("", h.RemoveFromWatchlist)
		}
	}

	webhook := v1.Group("/webhook")
	{
		webhook.POST("/discord", h.HandleDiscordWebhook)
		webhook.POST("/nft", h.WebhookNftHandler)
		webhook.POST("/topgg", h.WebhookUpvoteTopGG)
		webhook.POST("/discordbotlist", h.WebhookUpvoteDiscordBot)
	}

	verifyGroup := v1.Group("/verify")
	{
		verifyGroup.POST("/config", h.NewGuildConfigWalletVerificationMessage)
		verifyGroup.GET("/config/:guild_id", h.GetGuildConfigWalletVerificationMessage)
		verifyGroup.PUT("/config", h.UpdateGuildConfigWalletVerificationMessage)
		verifyGroup.DELETE("/config", h.DeleteGuildConfigWalletVerificationMessage)
		verifyGroup.POST("/generate", h.GenerateVerification)
		verifyGroup.POST("", h.VerifyWalletAddress)
	}

	whitelistCampaignGroup := v1.Group("/whitelist-campaigns")
	{
		whitelistCampaignGroup.POST("", h.CreateWhitelistCampaign)
		whitelistCampaignGroup.GET("", h.GetWhitelistCampaigns)
		whitelistCampaignGroup.GET("/:campaignId", h.GetWhitelistCampaignById)
		whitelistCampaignUserGroup := whitelistCampaignGroup.Group("/users")
		{
			whitelistCampaignUserGroup.POST("", h.AddWhitelistCampaignUsers)
			whitelistCampaignUserGroup.GET("", h.GetWhitelistCampaignUsers)
			whitelistCampaignUserGroup.GET("/:discordId", h.GetWhitelistCampaignUserByDiscordId)
			whitelistCampaignUserGroup.GET("/csv", h.GetWhitelistCampaignUsersCSV)
		}
	}
	nftsGroup := v1.Group("/nfts")
	{
		nftsGroup.GET("", h.ListAllNFTCollections)
		nftsGroup.GET("/tickers", h.GetNftTokenTickers)
		nftsGroup.GET("/:symbol/:id", h.GetNFTDetail)
		nftsGroup.GET("/:symbol/:id/activity", h.GetNFTActivity)
		nftsGroup.GET("/supported-chains", h.GetSupportedChains)
		nftsGroup.GET("/trading-volume", h.GetNFTTradingVolume)
		nftsGroup.POST("/sales-tracker", h.CreateNFTSalesTracker)
		nftsGroup.DELETE("/sales-tracker", h.DeleteNFTSalesTracker)
		nftsGroup.GET("/sales-tracker", h.GetAllNFTSalesTracker)
		nftsGroup.GET("/sales", h.GetNftSalesHandler)
		nftsGroup.GET("/new-listed", h.GetNewListedNFTCollection)
		nftsGroup.GET("/icons", h.GetNftMetadataAttrIcon)
		collectionsGroup := nftsGroup.Group("/collections")
		{
			collectionsGroup.GET("/:symbol/detail", h.GetDetailNftCollection)
			collectionsGroup.GET("/stats", h.GetCollectionCount)
			collectionsGroup.GET("", h.GetNFTCollections)
			collectionsGroup.GET("/suggestion", h.GetSuggestionNFTCollections)
			collectionsGroup.POST("", h.CreateNFTCollection)
			collectionsGroup.PATCH("/:address", h.UpdateNFTCollection) //to update collection images, delete after use
			collectionsGroup.GET("/:symbol", h.GetNFTTokens)
			collectionsGroup.GET("/tickers", h.GetNFTCollectionTickers)
			collectionsGroup.GET("/address/:address", h.GetNFTCollectionByAddressChain)
		}

		nftWatchlistGroup := nftsGroup.Group("/watchlist")
		{
			nftWatchlistGroup.GET("", h.GetNftWatchlist)
			nftWatchlistGroup.POST("", h.AddNftWatchlist)
			nftWatchlistGroup.DELETE("", h.DeleteNftWatchlist)
		}
	}
	giftGroup := v1.Group("/gift")
	{
		giftGroup.POST("/xp", h.GiftXpHandler)
	}
	twitterGroup := v1.Group("/twitter")
	{
		twitterGroup.POST("", h.CreateTwitterPost)
	}
	cacheGroup := v1.Group("/cache")
	{
		cacheGroup.POST("/upvote", h.SetUpvoteMessageCache)
	}
	usageGroup := v1.Group("/usage-stats")
	{
		usageGroup.POST("", h.AddServersUsageStat)
		usageGroup.GET("/gitbook", h.AddGitbookClick)
	}
	feedbackGroup := v1.Group("/feedback")
	{
		feedbackGroup.POST("", h.HandleUserFeedback)
	}
	// quests
	questGroup := v1.Group("/quests")
	{
		questGroup.GET("", h.GetUserQuestList)
		questGroup.POST("/progress", h.UpdateQuestProgress)
		questGroup.POST("/claim", h.ClaimQuestsRewards)
	}
}

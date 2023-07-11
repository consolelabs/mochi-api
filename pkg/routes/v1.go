package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/handler"
	"github.com/defipod/mochi/pkg/middleware"
)

// NewRoutes ...
func NewRoutes(r *gin.Engine, h *handler.Handler, cfg config.Config) {
	// API for Mpchi interface
	v1 := r.Group("/api/v1")
	v1.Use(middleware.WithAuthContext(cfg))

	tipBotGroup := v1.Group("/tip")
	{
		tipBotGroup.POST("/transfer", h.Tip.TransferToken)
	}

	guildGroup := v1.Group("/guilds")
	{
		guildGroup.GET("", h.Guild.GetGuilds)
		guildGroup.GET("/:guild_id", h.Guild.GetGuild)
		guildGroup.GET("/:guild_id/custom-tokens", h.ConfigDefi.ListAllCustomToken)
		guildGroup.GET("/user-managed", middleware.AuthGuard(cfg), h.Guild.ListMyGuilds)
		guildGroup.PUT("/:guild_id", h.Guild.UpdateGuild)
	}

	userGroup := v1.Group("/users")
	{
		userGroup.POST("", h.User.IndexUsers)
		userGroup.GET("/:id", h.User.GetUser)
		userGroup.GET("/gmstreak", h.User.GetUserCurrentGMStreak)
		// TODO
		userGroup.GET("/top", h.User.GetTopUsers)
		// TODO
		userGroup.GET("/profiles", h.User.GetUserProfile)

		// users/:id/wallet
		walletsGroup := userGroup.Group("/:id/wallets")
		{
			walletsGroup.GET("", h.Wallet.ListOwnedWallets)
			walletsGroup.GET("/:address", h.Wallet.GetOne)
			walletsGroup.GET("/:address/:type/assets", h.Wallet.ListAssets)
			walletsGroup.GET("/:address/:type/txns", h.Wallet.ListTransactions)
		}

		cexGroup := userGroup.Group("/:id/cexs") // this is profile id
		{
			cexGroup.GET("/binance", h.Dex.SumarizeBinanceAsset)
			cexGroup.GET("/:platform/assets", h.Dex.GetBinanceAssets)
		}

		// TODO: remove after migrate
		dexGroup := userGroup.Group("/:id/dexs") // this is profile id
		{
			dexGroup.GET("/binance", h.Dex.SumarizeBinanceAsset)
			dexGroup.GET("/:platform/assets", h.Dex.GetBinanceAssets)
		}

		// TODO: migrate wl apis to this group, add handler to handle Watchlist instead of using Wallet handler
		// For example: Token watchlist should be /users/:id/watchlists/tokens
		// Wallet watchlist should be /users/:id/watchlists/wallets
		// Action should be /users/:id/watchlists/wallets/action (POST)
		// It means:
		// 		What is the main entity? => User that have proper table in the DB, and can be extended by /:id
		//    What is the virtual entity? => Watchlist that is not a table in the DB, but can be extended by /watchlists to point to user's watchlist
		// 		Wallets is also the virtual entity here, it is different from the wallet in the DB, but it is the wallet that user want to watch
		// 		What is the action? => track/untrack
		userWatchListGroup := userGroup.Group("/:id/watchlists") //:id is profile_id
		{
			// wallets
			userWatchListGroup.POST("/wallets/track", h.Watchlist.TrackWallet)
			walletsGroup.POST("/wallets/untrack", h.Watchlist.UntrackWallet)
			userWatchListGroup.GET("/wallets", h.Watchlist.ListUserTrackingWallets)
			userWatchListGroup.PUT("/wallets/:address", h.Watchlist.UpdateTrackingWalletInfo)

			// tokens
			userWatchListGroup.POST("/tokens/track", h.Watchlist.TrackToken)
			userWatchListGroup.POST("/tokens/untrack", h.Watchlist.UntrackToken)
			userWatchListGroup.GET("/tokens", h.Watchlist.ListTrackingTokens)

			// nfts
			userWatchListGroup.POST("/nfts/track", h.Watchlist.TrackNft)
			userWatchListGroup.POST("/nfts/untrack", h.Watchlist.UntrackNft)
			userWatchListGroup.GET("/nfts", h.Watchlist.ListTrackingNfts)
		}

		watchListGroup := v1.Group("/watchlists")
		{
			watchListGroup.GET("/wallets", h.Watchlist.ListTrackingWallets)
		}

		userEarnGroup := userGroup.Group("/:id/earns") //:id is profile_id
		{
			userEarnGroup.POST("/airdrop-campaigns", h.AirdropCampaign.CreateProfileAirdropCampaign)
			userEarnGroup.GET("/airdrop-campaigns", h.AirdropCampaign.GetProfileAirdropCampaigns)
			userEarnGroup.DELETE("/airdrop-campaigns/:airdrop_campaign_id", h.AirdropCampaign.DeleteProfileAirdropCampaign)
		}

		// get offchain and onchain balances of user
		userGroup.GET("/:id/balances", h.User.GetUserBalance) //:id is profile_id
	}

	communityGroup := v1.Group("/community")
	{
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
		levelupGroup := communityGroup.Group("/levelup")
		{
			levelupGroup.GET("", h.Community.GetLevelUpMessage)
			levelupGroup.POST("", h.Community.UpsertLevelUpMessage)
			levelupGroup.DELETE("", h.Community.DeleteLevelUpMessage)
		}
		tagmeGroup := communityGroup.Group("/tagme")
		{
			tagmeGroup.GET("", h.Community.GetUserTag)
			tagmeGroup.POST("", h.Community.UpsertUserTag)
		}

	}

	configGroup := v1.Group("/config")
	{
		configGroup.GET("/sales-tracker", h.ConfigChannel.GetSalesTrackerConfig)
		configGroup.POST("/sales-tracker", h.ConfigChannel.CreateSalesTrackerConfig)

		configTwitterSaleGroup := configGroup.Group("/twitter-sales")
		{
			configTwitterSaleGroup.GET("", h.ConfigTwitterSale.Get)
			configTwitterSaleGroup.POST("", h.ConfigTwitterSale.Create)
		}
	}

	// v1/config-channels/
	configChannelGroup := v1.Group("/config-channels")
	{
		configChannelGroup.GET("/gm", h.ConfigChannel.GetGmConfig)
		configChannelGroup.POST("/gm", h.ConfigChannel.UpsertGmConfig)
		// config welcome channel
		configChannelGroup.GET("/welcome", h.ConfigChannel.GetWelcomeChannelConfig)
		configChannelGroup.POST("/welcome", h.ConfigChannel.UpsertWelcomeChannelConfig)
		configChannelGroup.DELETE("/welcome", h.ConfigChannel.DeleteWelcomeChannelConfig)
		// config tip notify channel
		configChannelGroup.POST("/tip-notify", h.ConfigChannel.CreateConfigNotify)
		configChannelGroup.GET("/tip-notify", h.ConfigChannel.ListConfigNotify)
		configChannelGroup.DELETE("/tip-notify/:id", h.ConfigChannel.DeleteConfigNotify)
	}

	// TODO:
	// v1/config/role/{guild-id}
	// v1/config/role/{guild-id}/reaction
	// // GET
	// // POST
	// // DELETE
	// v1/config/role/{guild-id}/default
	// v1/config/role/{guild-id}/level
	// v1/config/role/{guild-id}/nft
	// v1/config/role/{guild-id}/token
	// v1/config/role/{guild-id}/bot-manager
	configRoleGroup := configGroup.Group("/role/:guild_id")
	{
		roleReactionGroup := configRoleGroup.Group("/reaction")
		{
			roleReactionGroup.GET("", h.ConfigRoles.GetAllRoleReactionConfigs)
			roleReactionGroup.POST("", h.ConfigRoles.AddReactionRoleConfig)
			roleReactionGroup.DELETE("", h.ConfigRoles.RemoveReactionRoleConfig)
			roleReactionGroup.POST("/filter", h.ConfigRoles.FilterConfigByReaction)

		}
		roleDefaultGroup := configRoleGroup.Group("/default")
		{
			roleDefaultGroup.GET("", h.ConfigRoles.GetDefaultRolesByGuildID)
			roleDefaultGroup.POST("", h.ConfigRoles.CreateDefaultRole)
			roleDefaultGroup.DELETE("", h.ConfigRoles.DeleteDefaultRoleByGuildID)
		}
		levelRoleGroup := configRoleGroup.Group("/level")
		{
			levelRoleGroup.POST("", h.ConfigRoles.ConfigLevelRole)
			levelRoleGroup.GET("", h.ConfigRoles.GetLevelRoleConfigs)
			levelRoleGroup.DELETE("", h.ConfigRoles.RemoveLevelRoleConfig)
		}
		nftRoleGroup := configRoleGroup.Group("/nft")
		{
			nftRoleGroup.GET("", h.ConfigRoles.ListGuildGroupNFTRoles)
			nftRoleGroup.POST("", h.ConfigRoles.NewGuildGroupNFTRole)
			nftRoleGroup.DELETE("/group", h.ConfigRoles.RemoveGuildGroupNFTRole)
			nftRoleGroup.DELETE("", h.ConfigRoles.RemoveGuildNFTRole)
		}
		tokenRoleGroup := configRoleGroup.Group("/token")
		{
			tokenRoleGroup.POST("", h.ConfigRoles.CreateGuildTokenRole)
			tokenRoleGroup.GET("", h.ConfigRoles.ListGuildTokenRoles)
			tokenRoleGroup.DELETE("/:id", h.ConfigRoles.RemoveGuildTokenRole)
		}
		adminRoleGroup := configRoleGroup.Group("/bot-manager")
		{
			adminRoleGroup.POST("", h.ConfigRoles.CreateGuildAdminRoles)
			adminRoleGroup.GET("", h.ConfigRoles.ListGuildAdminRoles)
			adminRoleGroup.DELETE(":id", h.ConfigRoles.RemoveGuildAdminRole)
		}
	}

	// v1/config-defi
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
			defaultTickerGroup.GET("/:guild_id", h.ConfigDefi.GetListGuildDefaultTicker)
			defaultTickerGroup.POST("", h.ConfigDefi.SetGuildDefaultTicker)
		}
		monikerGroup := configDefiGroup.Group("/monikers")
		{
			monikerGroup.POST("", h.ConfigDefi.UpsertMonikerConfig)
			monikerGroup.GET("/:guild_id", h.ConfigDefi.GetMonikerByGuildID)
			monikerGroup.DELETE("", h.ConfigDefi.DeleteMonikerConfig)
			monikerGroup.GET("/default", h.ConfigDefi.GetDefaultMoniker)
		}
		tipRangeGroup := configDefiGroup.Group("/tip-range")
		{
			tipRangeGroup.POST("", h.ConfigDefi.UpsertGuildConfigTipRange)
			tipRangeGroup.GET("/:guild_id", h.ConfigDefi.GetGuildConfigTipRangeByGuildId)
			tipRangeGroup.DELETE("/:guild_id", h.ConfigDefi.RemoveGuildConfigTipRange)
		}
	}

	defiGroup := v1.Group("/defi")
	{
		defiGroup.GET("")
		defiGroup.GET("/tokens", h.Defi.GetSupportedTokens)
		defiGroup.GET("/token", h.Defi.GetSupportedToken)

		tokenSupportReqGroup := defiGroup.Group("/token-support")
		{
			tokenSupportReqGroup.GET("", h.Defi.GetUserRequestTokens)
			tokenSupportReqGroup.POST("", h.Defi.CreateUserTokenSupportRequest)
			tokenSupportReqGroup.PUT("/:id/approve", h.Defi.ApproveUserTokenSupportRequest)
			tokenSupportReqGroup.PUT("/:id/reject", h.Defi.RejectUserTokenSupportRequest)
		}

		// Data from CoinGecko
		defiGroup.GET("/market-chart", h.Defi.GetHistoricalMarketChart)
		defiGroup.GET("/coins/:id", h.Defi.GetCoin)
		defiGroup.GET("/coins/binance/:symbol", h.Defi.GetBinanceCoinData)
		defiGroup.GET("/coins", h.Defi.SearchCoins)
		defiGroup.GET("/coins/compare", h.Defi.CompareToken)
		defiGroup.GET("/chains", h.Defi.ListAllChain)
		defiGroup.GET("/market-data", h.Defi.GetCoinsMarketData)
		defiGroup.GET("/trending", h.Defi.GetTrendingSearch)
		defiGroup.GET("/top-gainer-loser", h.Defi.TopGainerLoser)

		priceAlertGroup := defiGroup.Group("/price-alert")
		{
			priceAlertGroup.GET("", h.Defi.GetUserListPriceAlert)
			priceAlertGroup.POST("", h.Defi.AddTokenPriceAlert)
			priceAlertGroup.DELETE("", h.Defi.RemoveTokenPriceAlert)
		}

		gasTrackerGroup := defiGroup.Group("/gas-tracker")
		{
			gasTrackerGroup.GET("", h.Defi.GetGasTracker)
			gasTrackerGroup.GET("/:chain", h.Defi.GetChainGasTracker)
		}
	}

	verifyGroup := v1.Group("/verify")
	{
		verifyGroup.POST("/config", h.Verify.NewGuildConfigWalletVerificationMessage)
		verifyGroup.GET("/config/:guild_id", h.Verify.GetGuildConfigWalletVerificationMessage)
		verifyGroup.PUT("/config", h.Verify.UpdateGuildConfigWalletVerificationMessage)
		verifyGroup.DELETE("/config", h.Verify.DeleteGuildConfigWalletVerificationMessage)
	}

	// api/v1/nfts
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
		// api/v1/nfts/collections
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

		defaultNftTickerGroup := nftsGroup.Group("/default-nft-ticker")
		{
			defaultNftTickerGroup.GET("", h.Nft.GetGuildDefaultNftTicker)
			defaultNftTickerGroup.POST("", h.Nft.SetGuildDefaultNftTicker)
		}
	}

	// TODO
	// api/v1/fiat
	fiatGroup := v1.Group("/fiat")
	{
		fiatGroup.GET("/historical-exchange-rates", h.Defi.GetFiatHistoricalExchangeRates)
	}

	webhook := v1.Group("/webhook")
	{
		webhook.POST("/discord", h.Webhook.HandleDiscordWebhook)
	}

	// TODO
	dataWebhookGroup := v1.Group("/data-webhook")
	{
		dataWebhookGroup.POST("/notify-nft-integration", h.Webhook.NotifyNftCollectionIntegration)
		dataWebhookGroup.POST("/notify-nft-add", h.Webhook.NotifyNftCollectionAdd)
		dataWebhookGroup.POST("/notify-nft-sync", h.Webhook.NotifyNftCollectionSync)
		dataWebhookGroup.POST("/notify-sale-marketplace", h.Webhook.NotifySaleMarketplace)
	}

	// api/v1/vault
	vaultGroup := v1.Group("/vault")
	{
		vaultGroup.GET("", h.Vault.GetVaults)
		vaultGroup.POST("", h.Vault.CreateVault)
		vaultGroup.POST("/config/channel", h.Vault.CreateConfigChannel)
		vaultGroup.GET("/config/channel", h.Vault.GetVaultConfigChannel)
		vaultGroup.PUT("/config/threshold", h.Vault.CreateConfigThreshold)
		vaultGroup.GET("/:vault_id/transaction", h.Vault.GetVaultTransactions) // this is also used for fortress-api
		treasurerGroup := vaultGroup.Group("/treasurer")
		{
			treasurerGroup.POST("/request", h.Vault.CreateTreasurerRequest)
			treasurerGroup.GET("/request/:request_id", h.Vault.GetTreasurerRequest)
			treasurerGroup.POST("", h.Vault.AddTreasurerToVault)
			treasurerGroup.DELETE("", h.Vault.RemoveTreasurerFromVault)
			treasurerGroup.POST("/submission", h.Vault.CreateTreasurerSubmission)
			treasurerGroup.POST("/result", h.Vault.CreateTreasurerResult)
			treasurerGroup.POST("/transfer", h.Vault.TransferVaultToken)
		}
		vaultGroup.GET("/detail", h.Vault.GetVaultDetail)
	}

	// api/v1/swap
	swapGroup := v1.Group("/swap")
	{
		swapGroup.GET("/route", h.Swap.GetSwapRoutes)
		swapGroup.POST("", h.Swap.ExecuteSwapRoutes)
	}

	apiKeyGroup := v1.Group("/api-key")
	{
		apiKeyGroup.POST("/me", middleware.ProfileAuthGuard(cfg), h.ApiKey.CreateApiKey)
		apiKeyGroup.POST("/binance", h.ApiKey.IntegrateBinanceKey)
		apiKeyGroup.POST("/unlink-binance", h.ApiKey.UnlinkBinance)
	}

	// TODO
	// move binance to this
	// api/v1/users/:id/associate/binance

	// api/v1/pk-pass
	pkpassGroup := v1.Group("/pk-pass")
	{
		pkpassGroup.GET("", h.PkPass.GeneratePkPass)
	}

	// api/v1/product-metadata
	productMetaData := v1.Group("/product-metadata")
	{
		productMetaData.GET("/emoji", h.Emojis.ListEmojis)
		productMetaData.GET("/copy/:type", h.Content.GetTypeContent)
	}

	// api/v1/earns
	earnGroup := v1.Group("/earns")
	{
		airdropCampaignGroup := earnGroup.Group("/airdrop-campaigns")
		{
			airdropCampaignGroup.GET("", h.AirdropCampaign.GetAirdropCampaigns)
			airdropCampaignGroup.GET("/:id", h.AirdropCampaign.GetAirdropCampaign)
			airdropCampaignGroup.POST("", h.AirdropCampaign.CreateAirdropCampaign)
			airdropCampaignGroup.GET("/stats", h.AirdropCampaign.GetAirdropCampaignStats)
		}
	}
}

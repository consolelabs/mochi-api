package routes

import (
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/handler"
	"github.com/defipod/mochi/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// NewRoutes ...
func NewRoutes(r *gin.Engine, h *handler.Handler, cfg config.Config) {

	v1 := r.Group("/api/v1")
	v1.Use(middleware.WithAuthContext(cfg))

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
		guildGroup.GET("/user-managed", middleware.AuthGuard(cfg), h.ListMyGuilds)

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
		userGroup.GET("/gmstreak", h.GetUserCurrentGMStreak)
		userGroup.GET("/top", h.GetTopUsers)
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
		profleGroup.GET("")
	}

	configGroup := v1.Group("/configs")
	{
		configGroup.GET("")
		configGroup.GET("/gm", h.GetGmConfig)
		configGroup.POST("/gm", h.UpsertGmConfig)
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
		tokenGroup := configGroup.Group("/tokens")
		{
			tokenGroup.GET("", h.GetGuildTokens)
			tokenGroup.POST("", h.UpsertGuildTokenConfig)
		}
		levelRoleGroup := configGroup.Group("/level-roles")
		{
			levelRoleGroup.POST("", h.ConfigLevelRole)
			levelRoleGroup.GET("/:guild_id", h.GetLevelRoleConfigs)
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
	}

	webhook := v1.Group("/webhook")
	{
		webhook.POST("/discord", h.HandleDiscordWebhook)
	}

	verifyGroup := v1.Group("/verify")
	{
		verifyGroup.POST("/config", h.NewGuildConfigWalletVerificationMessage)
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
		nftsGroup.GET("/:symbol/:id", h.GetNFTDetail)
		nftsGroup.GET("/supported-chains", h.GetSupportedChains)
		nftsGroup.POST("/collection", h.CreateNFTCollection)
	}
}

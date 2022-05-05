package routes

import (
	"github.com/defipod/mochi/pkg/handler"
	"github.com/gin-gonic/gin"
)

// NewRoutes ...
func NewRoutes(r *gin.Engine, h *handler.Handler) {

	v1 := r.Group("/api/v1")

	guildGroup := v1.Group("/guilds")
	{
		guildGroup.POST("", h.CreateGuild)
		guildGroup.GET("", h.GetGuilds)
		guildGroup.GET("/:guild_id", h.GetGuild)

		customCommandGroup := guildGroup.Group("/:guild_id/custom-commands")
		{
			customCommandGroup.POST("", h.CreateCustomCommand)
			customCommandGroup.GET("", h.ListCustomCommands)
			customCommandGroup.GET("/:command_id", h.GetCustomCommand)
			customCommandGroup.PUT("/:command_id", h.UpdateCustomCommand)
			customCommandGroup.DELETE("/:command_id", h.DeleteCustomCommand)
		}
	}

	userGroup := v1.Group("/users")
	{
		userGroup.POST("", h.IndexUsers)
		userGroup.GET("/:id", h.GetUser)
		userGroup.GET("/gmstreak", h.GetUserCurrentGMStreak)
	}

	communityGroup := v1.Group("/community")
	{
		invitesGroup := communityGroup.Group("/invites")
		{
			invitesGroup.GET("/", h.GetInvites)
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
		configGroup.POST("/gm", h.CreateGmConfig)
		roleReactionGroup := configGroup.Group("/reaction-roles")
		{
			roleReactionGroup.POST("", h.ProcessReactionEventByMessageID)
			roleReactionGroup.POST("/update-config", h.UpdateConfig)
		}
		defaultRoleGroup := configGroup.Group("/default-roles")
		{
			defaultRoleGroup.GET("", h.GetDefaultRolesByGuildID)
			defaultRoleGroup.POST("", h.CreateDefaultRole)
			defaultRoleGroup.DELETE("", h.DeleteDefaultRoleByGuildID)
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
	}

	webhook := v1.Group("/webhook")
	{
		webhook.POST("/discord", h.HandleDiscordWebhook)
	}

	verifyGroup := v1.Group("/verify")
	{
		verifyGroup.POST("/config", h.NewGuildConfigVerification)
		verifyGroup.POST("/generate", h.GenerateVerification)
		verifyGroup.POST("", h.VerifyWalletAddress)
	}
}

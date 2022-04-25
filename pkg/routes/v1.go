package routes

import (
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/handler"
	"github.com/gin-gonic/gin"
)

// NewRoutes ...
func NewRoutes(r *gin.Engine, h *handler.Handler, cfg config.Config) {

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
	}

	communityGroup := v1.Group("/community")
	{
		inviteHistoriesGroup := communityGroup.Group("/invite-histories")
		{
			inviteHistoriesGroup.POST("", h.IndexInviteHistory)
			inviteHistoriesGroup.GET("/count", h.CountByGuildUser)
			inviteHistoriesGroup.GET("/leaderboard/:id", h.GetInvitesLeaderboard)
		}

		invitesGroup := communityGroup.Group("/invites")
		{
			invitesGroup.GET("/", h.GetInvites)
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
		roleReactionGroup := configGroup.Group("/reaction_roles")
		{
			roleReactionGroup.POST("", h.GetAllReactionRolesByGuildID)
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
}

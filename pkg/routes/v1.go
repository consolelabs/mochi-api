package routes

import (
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/handler"
	"github.com/defipod/mochi/pkg/repo"
	"github.com/gin-gonic/gin"
)

// NewRoutes ...
func NewRoutes(r *gin.Engine, h *handler.Handler, cfg config.Config, s repo.Store) {

	v1 := r.Group("/api/v1")

	guildGroup := v1.Group("/guilds")
	{
		guildGroup.POST("", h.CreateGuild)
		guildGroup.GET("", h.GetGuilds)
		guildGroup.GET("/:id", h.GetGuild)
	}

	userGroup := v1.Group("/users")
	{
		userGroup.POST("", h.IndexUsers)
		userGroup.GET("/:id", h.GetUser)
	}

	inviteHistoriesGroup := v1.Group("/invite-histories")
	{
		inviteHistoriesGroup.POST("", h.IndexInviteHistory)
		inviteHistoriesGroup.GET("/count", h.CountByGuildUser)
	}

	profleGroup := v1.Group("/profiles")
	{
		profleGroup.GET("")
	}

	configGroup := v1.Group("/configs")
	{
		configGroup.GET("")
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

	communitiesGroup := v1.Group("/communities")
	{
		communitiesGroup.GET("")
	}

}

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
	}

	communitiesGroup := v1.Group("/communities")
	{
		communitiesGroup.GET("")
	}

}

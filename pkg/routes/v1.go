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
		guildGroup.GET("")
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

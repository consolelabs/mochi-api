package pg

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/repo"
	discordguilds "github.com/defipod/mochi/pkg/repo/discord_guilds"
)

// NewRepo new pg repo implimentation
func NewRepo(db *gorm.DB) *repo.Repo {
	return &repo.Repo{
		DiscordGuilds: discordguilds.NewPG(db),
	}
}

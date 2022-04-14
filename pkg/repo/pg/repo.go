package pg

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/repo"
	discordguilds "github.com/defipod/mochi/pkg/repo/discord_guilds"
	guildusers "github.com/defipod/mochi/pkg/repo/guild_users"
	invitehistories "github.com/defipod/mochi/pkg/repo/invite_histories"
	"github.com/defipod/mochi/pkg/repo/users"
)

// NewRepo new pg repo implimentation
func NewRepo(db *gorm.DB) *repo.Repo {
	return &repo.Repo{
		DiscordGuilds:   discordguilds.NewPG(db),
		InviteHistories: invitehistories.NewPG(db),
		Users:           users.NewPG(db),
		GuildUsers:      guildusers.NewPG(db),
	}
}

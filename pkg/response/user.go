package response

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
)

type GetUserResponse struct {
	ID                     string                  `json:"id"`
	Username               string                  `json:"username"`
	InDiscordWalletAddress *string                 `json:"in_discord_wallet_address"`
	InDiscordWalletNumber  *int64                  `json:"in_discord_wallet_number"`
	GuildUsers             []*GetGuildUserResponse `json:"guild_users"`
}

type GetGuildUserResponse struct {
	GuildID   string `json:"guild_id"`
	UserID    string `json:"user_id"`
	Nickname  string `json:"nickname"`
	InvitedBy string `json:"invited_by"`
}

type HandleUserActivityResponse struct {
	GuildID      string    `json:"guild_id"`
	UserID       string    `json:"user_id"`
	Action       string    `json:"action"`
	AddedXP      int       `json:"added_xp"`
	CurrentXP    int       `json:"current_xp"`
	CurrentLevel int       `json:"current_level"`
	Timestamp    time.Time `json:"timestamp"`
	LevelUp      bool      `json:"level_up"`
}

type GetTopUsersResponse struct {
	Author      *model.GuildUserXP  `json:"author"`
	Leaderboard []model.GuildUserXP `json:"leaderboard"`
}

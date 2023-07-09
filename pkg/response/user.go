package response

import (
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/defipod/mochi/pkg/model"
)

type User struct {
	ID            string                  `json:"id"`
	Username      string                  `json:"username"`
	GuildUsers    []*GetGuildUserResponse `json:"guild_users"`
	NrOfJoin      int64                   `json:"nr_of_join"`
	Discriminator string                  `json:"discriminator"`
}

type GetGuildUserResponse struct {
	GuildID   string `json:"guild_id"`
	UserID    string `json:"user_id"`
	Nickname  string `json:"nickname"`
	InvitedBy string `json:"invited_by"`
}

type HandleUserActivityResponse struct {
	ChannelID    string               `json:"channel_id"`
	GuildID      string               `json:"guild_id"`
	UserID       string               `json:"user_id"`
	Action       string               `json:"action"`
	AddedXP      int                  `json:"added_xp"`
	CurrentXP    int                  `json:"current_xp"`
	CurrentLevel int                  `json:"current_level"`
	NextLevel    *model.ConfigXpLevel `json:"next_level"`
	Timestamp    time.Time            `json:"timestamp"`
	LevelUp      bool                 `json:"level_up"`
}

type TopUser struct {
	Metadata    PaginationResponse  `json:"metadata"`
	Author      *model.GuildUserXP  `json:"author"`
	Leaderboard []model.GuildUserXP `json:"leaderboard"`
}

type GetUserProfileResponse struct {
	ID             string                       `json:"id"`
	AboutMe        string                       `json:"about_me"`
	CurrentLevel   *model.ConfigXpLevel         `json:"current_level"`
	NextLevel      *model.ConfigXpLevel         `json:"next_level"`
	GuildXP        int                          `json:"guild_xp"`
	NrOfActions    int                          `json:"nr_of_actions"`
	Progress       float64                      `json:"progress"`
	Guild          *model.DiscordGuild          `json:"guild"`
	GuildRank      int                          `json:"guild_rank"`
	UserFactionXps *model.UserFactionXpsMapping `json:"user_faction_xps"`
}

// For swagger
type GetDataUserProfileResponse struct {
	Data *GetUserProfileResponse `json:"data"`
}
type GetMyInfoResponse struct {
	Data *discordgo.User `json:"data"`
}

type GetUserResponse struct {
	Data User `json:"data"`
}

type GetTopUsersResponse struct {
	Data TopUser `json:"data"`
}
type GetUserCurrentGMStreakResponse struct {
	Data *model.DiscordUserGMStreak `json:"data"`
}

type CreateEnvelop struct {
	Data *model.Envelop `json:"data"`
}

type GetUserEnvelopStreak struct {
	Data *model.UserEnvelopStreak `json:"data"`
}

type UserBalanceResponse struct {
	Summarize []WalletAssetData  `json:"summarize"`
	Onchain   UserBalanceOnchain `json:"onchain"`
	Offchain  []WalletAssetData  `json:"offchain"`
}

type UserBalanceOnchain struct {
	Evm []WalletAssetData `json:"evm"`
	Sol []WalletAssetData `json:"sol"`
	Sui []WalletAssetData `json:"sui"`
	Ron []WalletAssetData `json:"ron"`
}

package response

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
)

type GuildNFTRolesResponse struct {
	model.GuildConfigNFTRole
	RoleName      string              `json:"role_name"`
	Color         int                 `json:"color"`
	NFTCollection model.NFTCollection `json:"nft_collection"`
}

type TwitterHashtag struct {
	UserID          string    `json:"user_id"`
	GuildID         string    `json:"guild_id"`
	ChannelID       string    `json:"channel_id"`
	RuleID          string    `json:"rule_id"`
	TwitterUsername []string  `json:"twitter_username"`
	Hashtag         []string  `json:"hashtag"`
	FromTwitter     []string  `json:"from_twitter"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type GetGmConfigResponse struct {
	Message string                 `json:"message"`
	Data    *model.GuildConfigGmGn `json:"data"`
}

type GetSalesTrackerConfigResponse struct {
	Message string                         `json:"message"`
	Data    *model.GuildConfigSalesTracker `json:"data"`
}

type GetGuildTokensResponse struct {
	Data []model.Token `json:"data"`
}

type GetLevelRoleConfigsResponse struct {
	Data []model.GuildConfigLevelRole `json:"data"`
}

type ListGuildNFTRolesResponse struct {
	Data []GuildNFTRolesResponse `json:"data"`
}

type NewGuildNFTRoleResponse struct {
	Message string                    `json:"message"`
	Data    *model.GuildConfigNFTRole `json:"data"`
}

type EditGuildNFTRoleResponse struct {
	Message string                    `json:"message"`
	Data    *model.GuildConfigNFTRole `json:"data"`
}

type GetRepostReactionConfigsResponse struct {
	Data []model.GuildConfigRepostReaction `json:"data"`
}

type ToggleActivityConfigResponse struct {
	Message string                     `json:"message"`
	Data    *model.GuildConfigActivity `json:"data"`
}

type GetAllTwitterConfigResponse struct {
	Message string                         `json:"message"`
	Data    []model.GuildConfigTwitterFeed `json:"data"`
}

type GetTwitterHashtagConfigResponse struct {
	Data *TwitterHashtag `json:"data"`
}

type GetAllTwitterHashtagConfigResponse struct {
	Data []TwitterHashtag `json:"data"`
}

type GetDefaultTokenResponse struct {
	Data *model.Token `json:"data"`
}

package request

import (
	"fmt"

	"github.com/defipod/mochi/pkg/model"
)

type UpsertGmConfigRequest struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
}
type UpsertSalesTrackerConfigRequest struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
}
type UpsertGuildTokenConfigRequest struct {
	GuildID string `json:"guild_id"`
	Symbol  string `json:"symbol"`
	Active  bool   `json:"active"`
}

type ConfigLevelRoleRequest struct {
	GuildID string `json:"guild_id"`
	RoleID  string `json:"role_id"`
	Level   int    `json:"level"`
}

type ConfigNFTRoleRequest struct {
	model.GuildConfigNFTRole
}

func (cfg ConfigNFTRoleRequest) Validate() error {
	if cfg.GuildID == "" {
		return fmt.Errorf("guild_id is required")
	}
	if cfg.RoleID == "" {
		return fmt.Errorf("role_id is required")
	}
	if !cfg.NFTCollectionID.Valid {
		return fmt.Errorf("nft_collection_id is invalid")
	}
	return nil
}

type ConfigRepostRequest struct {
	GuildID         string `json:"guild_id"`
	Emoji           string `json:"emoji"`
	Quantity        int    `json:"quantity"`
	RepostChannelID string `json:"repost_channel_id"`
}

type TwitterHashtag struct {
	UserID          string   `json:"user_id"`
	TwitterUsername []string `json:"twitter_username"`
	GuildID         string   `json:"guild_id"`
	RuleID          string   `json:"rule_id"`
	ChannelID       string   `json:"channel_id"`
	Hashtag         []string `json:"hashtag"`
	FromTwitter     []string `json:"from_twitter"`
}

type ConfigDefaultTokenRequest struct {
	GuildID string `json:"guild_id"`
	Symbol  string `json:"symbol"`
}

type ConfigDefaultCollection struct {
	GuildID string `json:"guild_id"`
	Symbol  string `json:"symbol"`
	Address string `json:"address"`
	ChainID string `json:"chain"`
}

type GetGuildDefaultTickerRequest struct {
	GuildID string `json:"guild_id" form:"guild_id" binding:"required"`
	Query   string `json:"query" form:"query" binding:"required"`
}

type GuildConfigDefaultTickerRequest struct {
	GuildID       string `json:"guild_id"`
	Query         string `json:"query"`
	DefaultTicker string `json:"default_ticker"`
}

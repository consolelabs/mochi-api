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
	GuildID   string   `json:"guild_id"`
	ChannelID string   `json:"channel_id"`
	Hashtag   []string `json:"hashtag"`
}

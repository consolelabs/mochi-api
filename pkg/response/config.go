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
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

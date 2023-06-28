package response

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
)

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

type ListGuildNFTRoleConfigsResponse struct {
	Id                   string                `json:"id"`
	GuildId              string                `json:"guild_id"`
	GroupName            string                `json:"group_name"`
	RoleId               string                `json:"role_id"`
	NumberOfTokens       int                   `json:"number_of_tokens"`
	RoleName             string                `json:"role_name"`
	Color                int                   `json:"color"`
	NFTCollectionConfigs []NFTCollectionConfig `json:"nft_collection_configs"`
}

type GetGmConfigResponse struct {
	Message string                 `json:"message"`
	Data    *model.GuildConfigGmGn `json:"data"`
}

type GetWelcomeChannelConfigResponse struct {
	Message string                           `json:"message"`
	Data    *model.GuildConfigWelcomeChannel `json:"data"`
}

type GetVoteChannelConfigResponse struct {
	Message string                        `json:"message"`
	Data    *model.GuildConfigVoteChannel `json:"data"`
}

type GetSalesTrackerConfigResponse struct {
	Data []model.GuildConfigSalesTracker `json:"data"`
}

type GetGuildTokensResponse struct {
	Data []model.Token `json:"data"`
}

type GetLevelRoleConfigsResponse struct {
	*PaginationResponse `json:",omitempty"`
	Data                []model.GuildConfigLevelRole `json:"data"`
}

type ListGuildGroupNFTRolesResponse struct {
	Data []ListGuildNFTRoleConfigsResponse `json:"data"`
}

type NewGuildGroupNFTRoleResponse struct {
	Message string                      `json:"message"`
	Data    *ConfigGroupNFTRoleResponse `json:"data"`
}

type GetUpvoteTiersConfig struct {
	Message string                   `json:"message"`
	Data    []model.UpvoteStreakTier `json:"data"`
}

type ConfigGroupNFTRoleResponse struct {
	GuildID              string                `json:"guild_id"`
	RoleID               string                `json:"role_id"`
	GroupName            string                `json:"group_name"`
	NFTCollectionConfigs []NFTCollectionConfig `json:"nft_collection_configs"`
	NumberOfTokens       int                   `json:"number_of_tokens"`
}

type NFTCollectionConfig struct {
	ID           string    `json:"id"`
	CollectionID string    `json:"collection_id"`
	Address      string    `json:"address"`
	ExplorerUrl  string    `json:"explorer_url"`
	ChainName    string    `json:"chain_name"`
	Name         string    `json:"name"`
	Symbol       string    `json:"symbol"`
	ChainID      string    `json:"chain_id"`
	ERCFormat    string    `json:"erc_format"`
	IsVerified   bool      `json:"is_verified"`
	CreatedAt    time.Time `json:"created_at"`
	Image        string    `json:"image"`
	Author       string    `json:"author"`
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

type GetTwitterBlackListResponse struct {
	Data []model.GuildConfigTwitterBlacklist `json:"data"`
}

type GetDefaultTokenResponse struct {
	Data *model.Token `json:"data"`
}

type GetGuildPruneExcludeResponse struct {
	Message string                 `json:"message"`
	Data    *GuildPruneExcludeList `json:"data"`
}

type GuildPruneExcludeList struct {
	GuildID string   `json:"guild_id"`
	Roles   []string `json:"roles"`
}

type GuildConfigDefaultCurrencyResponse struct {
	GuildID     string                     `json:"guild_id"`
	TipBotToken *model.OffchainTipBotToken `json:"tip_bot_token"`
	UpdatedAt   time.Time                  `json:"updated_at"`
	CreatedAt   time.Time                  `json:"created_at"`
}

type CreateProposalChannelConfigResponse struct {
	Data model.GuildConfigDaoProposal `json:"data"`
}

type CreateGuildTokenRole struct {
	Data *model.GuildConfigTokenRole `json:"data"`
}

type ListGuildTokenRoles struct {
	Data []model.GuildConfigTokenRole `json:"data"`
	Meta map[string]interface{}       `json:"meta"`
}

type UpdateGuildTokenRole struct {
	Data *model.GuildConfigTokenRole `json:"data"`
}

type GetGuildLevelUpMessage struct {
	Data *model.GuildConfigLevelupMessage `json:"data"`
}

type CreateGuildXPRole struct {
	Data *model.GuildConfigXPRole `json:"data"`
}

type ListGuildXPRoles struct {
	Data []model.GuildConfigXPRole `json:"data"`
}

type GuildConfigDaoTrackerResponse struct {
	Data *[]model.GuildConfigDaoTracker `json:"data"`
}

type CreateGuildMixRole struct {
	Data *model.GuildConfigMixRole `json:"data"`
}

type ListGuildMixRoles struct {
	Data []model.GuildConfigMixRole `json:"data"`
}

type GuildConfigTipRangeResponse struct {
	GuildID   string    `json:"guild_id"`
	Min       *float64  `json:"min,omitempty"`
	Max       *float64  `json:"max,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type ListGuildAdminRoles struct {
	Data []model.GuildConfigAdminRole `json:"data"`
}

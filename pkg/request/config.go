package request

import (
	"fmt"
)

type UpsertGmConfigRequest struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
	Msg       string `json:"msg"`
	Emoji     string `json:"emoji"`
	Sticker   string `json:"sticker"`
}

type UpsertWelcomeConfigRequest struct {
	GuildID    string `json:"guild_id"`
	ChannelID  string `json:"channel_id"`
	WelcomeMsg string `json:"welcome_message"`
}
type DeleteWelcomeConfigRequest struct {
	GuildID string `json:"guild_id"`
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
	GuildID string `uri:"guild_id"`
	RoleID  string `json:"role_id"`
	Level   int    `json:"level"`
}

type ConfigGroupNFTRoleRequest struct {
	GuildID           string   `uri:"guild_id"`
	RoleID            string   `json:"role_id"`
	GroupName         string   `json:"group_name"`
	CollectionAddress []string `json:"collection_address"`
	NumberOfTokens    int      `json:"number_of_tokens"`
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

type GetGuildDefaultNftTickerRequest struct {
	GuildID string `json:"guild_id" form:"guild_id" binding:"required"`
	Query   string `json:"query" form:"query" binding:"required"`
}

type GuildConfigDefaultTickerRequest struct {
	GuildID       string `json:"guild_id"`
	Query         string `json:"query"`
	DefaultTicker string `json:"default_ticker"`
}

type GuildConfigDefaultNftTickerRequest struct {
	GuildID           string `json:"guild_id"`
	Query             string `json:"query"`
	CollectionAddress string `json:"collection_address"`
	ChainId           int64  `json:"chain_id"`
	Symbol            string `json:"symbol"`
}

type UpsertGuildPruneExcludeRequest struct {
	GuildID string `json:"guild_id"`
	RoleID  string `json:"role_id"`
}

type EditMessageRepostRequest struct {
	GuildID         string `json:"guild_id"`
	OriginMessageID string `json:"origin_message_id"`
	OriginChannelID string `json:"origin_channel_id"`
	RepostChannelID string `json:"repost_channel_id"`
	RepostMessageID string `json:"repost_message_id"`
}

type ConfigRepostReactionStartStop struct {
	GuildID         string `json:"guild_id"`
	EmojiStart      string `json:"emoji_start"`
	EmojiStop       string `json:"emoji_stop"`
	RepostChannelID string `json:"repost_channel_id"`
}

type BalcklistChannelRepostConfigRequest struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
}

type UpsertGuildDefaultCurrencyRequest struct {
	GuildID string `json:"guild_id"`
	Symbol  string `json:"symbol"`
}

type CreateGuildTokenRole struct {
	GuildID string  `uri:"guild_id"`
	RoleID  string  `json:"role_id"`
	Chain   string  `json:"chain"`
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
}

type UpdateGuildTokenRole struct {
	RoleID  *string  `json:"role_id"`
	Chain   *string  `json:"chain" binding:"required_with=Address"`
	Address *string  `json:"address" binding:"required_with=Chain"`
	Amount  *float64 `json:"amount"`
}

type UpsertGuildLevelUpMessageRequest struct {
	GuildID   string `json:"guild_id"`
	Message   string `json:"message"`
	ImageURL  string `json:"image_url"`
	ChannelID string `json:"channel_id"`
}

type CreateGuildXPRole struct {
	GuildID string `json:"guild_id" binding:"required"`
	RoleID  string `json:"role_id" binding:"required"`
	XP      int    `json:"xp"`
}

type UpsertGuildConfigTipRangeRequest struct {
	GuildID string   `json:"guild_id"`
	Min     *float64 `json:"min"`
	Max     *float64 `json:"max"`
}

func (r *UpsertGuildConfigTipRangeRequest) Validate() error {
	if r.Max != nil && *r.Max <= 0 {
		return fmt.Errorf("max value is invalid")
	}
	if r.Min != nil && *r.Min < 0 {
		return fmt.Errorf("min value is invalid")
	}

	return nil
}

type ConfigLogChannelRequest struct {
}

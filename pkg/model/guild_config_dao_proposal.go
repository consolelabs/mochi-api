package model

import "time"

type GuildConfigDaoProposal struct {
	Id                 int64     `json:"id"`
	GuildId            string    `json:"guild_id"`
	ProposalChannelId  string    `json:"proposal_channel_id"`
	GuidelineChannelId string    `json:"guideline_channel_id"`
	Type               string    `json:"type"`
	RequiredAmount     int64     `json:"required_amount"`
	ChainID            int64     `json:"chain_id"`
	Symbol             string    `json:"symbol"`
	Address            string    `json:"address"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (GuildConfigDaoProposal) TableName() string {
	return "guild_config_dao_proposal"
}

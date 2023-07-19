package model

import (
	"time"
)

type ProposalVotingType string

const (
	NFT         ProposalVotingType = "nft_collection"
	CryptoToken ProposalVotingType = "crypto_token"
)

type ProposalAuthorityType string

const (
	Admin       ProposalAuthorityType = "admin"
	TokenHolder ProposalAuthorityType = "token_holder"
)

type GuildConfigDaoProposal struct {
	Id                 int64                 `json:"id"`
	GuildId            string                `json:"guild_id"`
	ProposalChannelId  string                `json:"proposal_channel_id"`
	GuidelineChannelId string                `json:"guideline_channel_id"`
	Authority          ProposalAuthorityType `json:"authority"`
	Type               *ProposalVotingType   `json:"type"`
	RequiredAmount     string                `json:"required_amount" gorm:"type:numeric"`
	ChainID            int64                 `json:"chain_id"`
	Symbol             string                `json:"symbol"`
	Address            string                `json:"address"`
	CreatedAt          time.Time             `json:"created_at"`
	UpdatedAt          time.Time             `json:"updated_at"`
}

func (GuildConfigDaoProposal) TableName() string {
	return "guild_config_dao_proposal"
}

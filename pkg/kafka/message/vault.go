package message

type Message struct {
	Topic string `json:"topic"`
}

type VaultVoteTreasurer struct {
	Type              string            `json:"type"`
	VaultVoteMetadata VaultVoteMetadata `json:"vault_vote_metadata"`
}

type VaultVoteMetadata struct {
	TreasurerProfileId       string            `json:"treasurer_profile_id"`
	TreasurerVote            string            `json:"treasurer_vote"`
	RequestId                int64             `json:"request_id"`
	DaoThresholdInPercentage float64           `json:"dao_threshold_in_percentage"`
	DaoThresholdInNumber     int64             `json:"dao_threshold_in_number"`
	CurrentApproval          int64             `json:"current_approval"`
	CurrentRejection         int64             `json:"current_rejection"`
	CurrentWaiting           int64             `json:"current_waiting"`
	DaoGuild                 string            `json:"dao_guild"`
	DaoVault                 string            `json:"dao_vault"`
	Message                  string            `json:"message"`
	MessageUrl               string            `json:"message_url"`
	DaoVaultTotalTreasurer   map[string]string `json:"dao_vault_total_treasurer"`
	Action                   VaultAction       `json:"action"`
}

type VaultAction struct {
	Type                               string                             `json:"type"`
	VaultChangeTreasurerActionMetadata VaultChangeTreasurerActionMetadata `json:"vault_change_treasurer_action_metadata"`
	VaultTransferActionMetadata        VaultTransferActionMetadata        `json:"vault_transfer_action_metadata"`
}

type VaultChangeTreasurerActionMetadata struct {
	TreasurerProfileId string `json:"treasurer_profile_id"`
	TreasurerAction    string `json:"treasurer_action"`
}

type VaultTransferActionMetadata struct {
	TokenAmount        string `json:"token_amount"`
	TokenDecimal       int64  `json:"token_decimal"`
	TokenAmountInUsd   string `json:"token_amount_in_usd"`
	Token              string `json:"token"`
	RecipientProfileId string `json:"recipient_profile_id"`
}

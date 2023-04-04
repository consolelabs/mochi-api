package model

import "time"

type TreasurerSubmission struct {
	Id               int64            `json:"id"`
	VaultId          int64            `json:"vault_id"`
	GuildId          string           `json:"guild_id"`
	RequestId        int64            `json:"request_id"`
	Submitter        string           `json:"submitter"`
	Status           string           `json:"status"`
	TreasurerRequest TreasurerRequest `json:"treasurer_request" gorm:"foreignKey:RequestId;references:Id"`
	Vault            Vault            `json:"vault" gorm:"foreignKey:VaultId;references:Id"`
	CreatedAt        time.Time        `json:"created_at"`
	UpdatedAt        time.Time        `json:"updated_at"`
}

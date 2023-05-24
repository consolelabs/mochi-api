package model

import "time"

type Vault struct {
	Id                  int64       `json:"id"`
	GuildId             string      `json:"guild_id"`
	Name                string      `json:"name"`
	Threshold           string      `json:"threshold"`
	WalletAddress       string      `json:"wallet_address"`
	SolanaWalletAddress string      `json:"solana_wallet_address"`
	WalletNumber        int64       `json:"wallet_number"`
	CreatedAt           time.Time   `json:"created_at"`
	UpdatedAt           time.Time   `json:"updated_at"`
	Treasurers          []Treasurer `json:"treasurers" gorm:"foreignkey:VaultId"`
	TotalAmountEVM      string      `json:"total_amount_evm" gorm:"-"`
	TotalAmountSolana   string      `json:"total_amount_solana" gorm:"-"`
}

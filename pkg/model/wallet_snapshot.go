package model

import "time"

type WalletSnapshot struct {
	Id              int64     `json:"id" gorm:"primary_key"`
	WalletAddress   string    `json:"wallet_address"`
	IsEvm           bool      `json:"is_evm"`
	TotalUsdBalance string    `json:"total_usd_balance"`
	SnapshotTime    time.Time `json:"snapshot_time"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type EvmAddress struct {
	Address   string         `json:"address"`
	ChainType JSONNullString `json:"chain_type"`
	ProfileId string         `json:"profile_id"`
}

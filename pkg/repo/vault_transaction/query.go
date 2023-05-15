package vaulttransaction

type VaultTransactionQuery struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	VaultId   int64  `json:"vault_id"`
}

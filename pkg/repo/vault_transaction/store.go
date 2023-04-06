package vaulttransaction

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(vaultTx *model.VaultTransaction) (*model.VaultTransaction, error)
	GetRecentTx(vaultId int64, guildId string) (vaultTxs []model.VaultTransaction, err error)
}

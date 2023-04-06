package vaulttransaction

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) Create(vaultTx *model.VaultTransaction) (*model.VaultTransaction, error) {
	return vaultTx, pg.db.Create(vaultTx).Error
}

func (pg *pg) GetRecentTx(vaultId int64, guildId string) (vaultTxs []model.VaultTransaction, err error) {
	return vaultTxs, pg.db.Where("vault_id = ? AND guild_id = ?", vaultId, guildId).Order("created_at desc").Limit(10).Find(&vaultTxs).Error
}

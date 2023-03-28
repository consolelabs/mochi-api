package vault

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

func (pg *pg) Create(vault *model.Vault) error {
	return pg.db.Create(vault).Error
}

func (pg *pg) GetByGuildId(guildId string) (vaults []model.Vault, err error) {
	return vaults, pg.db.Where("guild_id = ?", guildId).Find(&vaults).Error
}

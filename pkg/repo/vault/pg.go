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

func (pg *pg) Create(vault *model.Vault) (*model.Vault, error) {
	return vault, pg.db.Create(vault).Error
}

func (pg *pg) GetByGuildId(guildId string) (vaults []model.Vault, err error) {
	return vaults, pg.db.Where("guild_id = ?", guildId).Find(&vaults).Error
}

func (pg *pg) UpdateThreshold(vault *model.Vault) (*model.Vault, error) {
	return vault, pg.db.Model(&vault).Where("guild_id = ? and name = ?", vault.GuildId, vault.Name).Update("threshold", vault.Threshold).Error
}

func (pg *pg) GetById(id int64) (vault *model.Vault, err error) {
	return vault, pg.db.First(&vault, id).Error
}

func (pg *pg) GetByNameAndGuildId(name string, guildId string) (vault *model.Vault, err error) {
	return vault, pg.db.Where("name = ? and guild_id = ?", name, guildId).First(&vault).Error
}

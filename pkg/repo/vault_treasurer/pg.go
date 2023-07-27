package treasurer

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

func (pg *pg) Create(treasurer *model.VaultTreasurer) (*model.VaultTreasurer, error) {
	return treasurer, pg.db.Create(treasurer).Error
}

func (pg *pg) GetByVaultId(vaultId int64) (treasurers []model.VaultTreasurer, err error) {
	return treasurers, pg.db.Where("vault_id = ?", vaultId).Find(&treasurers).Error
}

func (pg *pg) Delete(treasurer *model.VaultTreasurer) (*model.VaultTreasurer, error) {
	return treasurer, pg.db.Where("guild_id = ? and vault_id = ? and user_profile_id = ?", treasurer.GuildId, treasurer.VaultId, treasurer.UserProfileId).Delete(treasurer).Error
}

func (pg *pg) GetByGuildIdAndVaultId(guildId string, vaultId int64) (treasurer []model.VaultTreasurer, err error) {
	return treasurer, pg.db.Where("guild_id = ? and vault_id = ?", guildId, vaultId).Find(&treasurer).Error
}

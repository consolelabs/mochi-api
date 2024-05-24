package vault

import (
	"database/sql"

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
	return vault, pg.db.Preload("DiscordGuild").First(&vault, id).Error
}

func (pg *pg) GetByNameAndGuildId(name string, guildId string) (vault *model.Vault, err error) {
	return vault, pg.db.Where("name = ? and guild_id = ?", name, guildId).First(&vault).Error
}

func (pg *pg) GetLatestWalletNumber() (walletNumber sql.NullInt64, err error) {
	row := pg.db.Table("vaults").Select("max(wallet_number)").Row()
	err = row.Scan(&walletNumber)
	return walletNumber, err
}

func (pg *pg) List(q ListQuery) (vaults []model.Vault, err error) {
	db := pg.db
	if !q.GetArchived {
		db = db.Where("vaults.deleted_at IS NULL")
	}
	if q.GuildID != "" {
		db = db.Where("vaults.guild_id = ?", q.GuildID)
	}
	if q.UserProfileID != "" {
		db = db.Joins("join vault_treasurers on vaults.id = vault_treasurers.vault_id").Where("vault_treasurers.user_profile_id = ?", q.UserProfileID)
	}
	if q.EvmWallet != "" {
		db = db.Where("vaults.wallet_address = ?", q.EvmWallet)
	}
	if q.SolanaWallet != "" {
		db = db.Where("vaults.solana_wallet_address = ?", q.SolanaWallet)
	}
	if q.Threshold != "" {
		db = db.Where("vaults.threshold = ?", q.Threshold)
	}
	if len(q.VaultIDs) > 0 {
		db = db.Where("vaults.id IN ?", q.VaultIDs)
	}
	return vaults, db.Preload("VaultTreasurers").Preload("DiscordGuild").Find(&vaults).Error
}

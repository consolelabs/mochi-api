package discordwalletverification

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetOne(dicordID, guildID string) (*model.DiscordWalletVerification, error) {
	var dv model.DiscordWalletVerification
	return &dv, pg.db.Where("user_discord_id = ? and guild_id = ?", dicordID, guildID).First(&dv).Error
}

func (pg *pg) UpsertOne(v model.DiscordWalletVerification) error {
	tx := pg.db.Begin()
	err := tx.Table("discord_wallet_verifications").Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "user_discord_id"},
			{Name: "guild_id"},
		},
		DoUpdates: clause.AssignmentColumns([]string{"code", "created_at"}),
	}).Create(&v).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (pg *pg) GetByValidCode(code string) (*model.DiscordWalletVerification, error) {
	var dv model.DiscordWalletVerification
	return &dv, pg.db.Where("code = ? and created_at >= (now() - '10 minutes'::interval)", code).First(&dv).Error
}

func (pg *pg) DeleteByCode(code string) error {
	return pg.db.Where("code = ?", code).Delete(&model.DiscordWalletVerification{}).Error
}

func (pg *pg) TotalVerifiedWallets() (count int64, err error) {
	return count, pg.db.Table("discord_wallet_verifications").Count(&count).Error
}

func (pg *pg) TotalVerifiedWalletsByGuildID(guildId string) (count int64, err error) {
	return count, pg.db.Table("discord_wallet_verifications").Where("guild_id = ?", guildId).Count(&count).Error
}

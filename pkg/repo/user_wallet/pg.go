package userwallet

import (
	"database/sql"
	"strings"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (pg *pg) GetOneByDiscordIDAndGuildID(discordID, guildID string) (*model.UserWallet, error) {
	var uw model.UserWallet
	return &uw, pg.db.Where("user_discord_id = ? and guild_id = ?", discordID, guildID).First(&uw).Error
}

func (pg *pg) GetOneByGuildIDAndAddress(guildID, address string) (*model.UserWallet, error) {
	var uw model.UserWallet
	return &uw, pg.db.Where("guild_id = ? and address = ?", guildID, address).First(&uw).Error
}

func (pg *pg) UpsertOne(uw model.UserWallet) error {
	// make sure address in lowercase
	uw.Address = strings.ToLower(uw.Address)
	uw.ChainType = model.JSONNullString{NullString: sql.NullString{String: string(util.GetChainTypeFromAddress(uw.Address)), Valid: true}}

	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "user_discord_id"},
			{Name: "guild_id"},
		},
		DoUpdates: clause.AssignmentColumns([]string{"address", "chain_type", "created_at"}),
	}).Create(&uw).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

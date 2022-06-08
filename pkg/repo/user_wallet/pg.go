package userwallet

import (
	"database/sql"
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/util"
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

func (pg *pg) ListWalletAddresses(chainType ...string) ([]model.WalletAddress, error) {
	addresses := make([]model.WalletAddress, 0)

	q := pg.db.Table("user_wallets").Select("address, chain_type")

	if len(chainType) > 0 {
		q = q.Where("chain_type in (?)", chainType)
	}

	rows, err := q.Group("address, chain_type").Rows()
	if err != nil {
		return nil, fmt.Errorf("failed to query wallet addresses: %w", err)
	}

	for rows.Next() {
		var address model.WalletAddress
		if err := rows.Scan(&address.Address, &address.ChainType); err != nil {
			return nil, fmt.Errorf("failed to scan wallet address: %w", err)
		}

		addresses = append(addresses, address)
	}

	return addresses, nil
}

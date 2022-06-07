package token

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetBySymbol(symbol string, botSupported bool) (model.Token, error) {
	var token model.Token
	return token, pg.db.First(&token, "lower(symbol) = lower(?) AND discord_bot_supported = ?", symbol, botSupported).Error
}

func (pg *pg) GetAllSupported() ([]model.Token, error) {
	var tokens []model.Token
	return tokens, pg.db.Where("discord_bot_supported = TRUE").Find(&tokens).Error
}

func (pg *pg) GetByAddress(address string, chainID int) (*model.Token, error) {
	var token model.Token
	return &token, pg.db.First(&token, "lower(address) = lower(?) and chain_id = ?", address, chainID).Error
}

func (pg *pg) GetDefaultTokens() ([]model.Token, error) {
	var tokens []model.Token
	return tokens, pg.db.Where("guild_default = TRUE").Find(&tokens).Error
}

func (pg *pg) UpsertOne(token model.Token) error {

	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "id"},
		},
		UpdateAll: true,
	}).Create(&token).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

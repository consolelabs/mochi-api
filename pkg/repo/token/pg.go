package token

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

// TODO: remove all Get funcs, replace with List()
func (pg *pg) GetBySymbol(symbol string, botSupported bool) (model.Token, error) {
	var token model.Token
	return token, pg.db.Preload("Chain").First(&token, "lower(symbol) = lower(?) AND discord_bot_supported = ?", symbol, botSupported).Error
}

func (pg *pg) GetAllSupported(q ListQuery) ([]model.Token, int64, error) {
	var tokens []model.Token
	var total int64
	db := pg.db.Model(&model.Token{}).Preload("Chain").Where("discord_bot_supported = TRUE").Order("name ASC")
	db = db.Count(&total).Offset(q.Offset)
	if q.Limit != 0 {
		db = db.Limit(q.Limit)
	}
	return tokens, total, db.Find(&tokens).Error
}

func (pg *pg) GetByAddress(address string, chainID int) (*model.Token, error) {
	var token model.Token
	return &token, pg.db.Preload("Chain").First(&token, "lower(address) = lower(?) and chain_id = ?", address, chainID).Error
}

func (pg *pg) GetDefaultTokens() ([]model.Token, error) {
	var tokens []model.Token
	return tokens, pg.db.Preload("Chain").Where("guild_default = TRUE").Find(&tokens).Error
}

func (pg *pg) CreateOne(record *model.Token) error {
	return pg.db.Create(record).Error
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

func (pg *pg) GetAll() ([]model.Token, error) {
	var tokens []model.Token
	return tokens, pg.db.Find(&tokens).Error
}

func (pg *pg) Get(id int) (model *model.Token, err error) {
	return model, pg.db.First(&model, id).Error
}

func (pg *pg) GetSupportedTokenByGuildId(guildID string) ([]model.Token, error) {
	var tokens []model.Token
	return tokens, pg.db.
		Table("tokens").
		Joins("JOIN guild_config_tokens ON guild_config_tokens.token_id = tokens.id").
		Where("guild_config_tokens.guild_id = ?", guildID).
		Find(&tokens).Error
}

func (pg *pg) GetOneBySymbol(symbol string) (*model.Token, error) {
	token := &model.Token{}
	if err := pg.db.Preload("Chain").First(token, "lower(symbol) = lower(?)", symbol).Error; err != nil {
		return nil, err
	}
	return token, nil
}

func (pg *pg) GetDefaultTokenByGuildID(guildID string) (model.Token, error) {
	var token model.Token
	return token, pg.db.
		Preload("Chain").
		Table("tokens").
		Joins("JOIN guild_config_tokens ON guild_config_tokens.token_id = tokens.id").
		Where("guild_config_tokens.guild_id = ? AND is_default = TRUE", guildID).
		First(&token).Error
}

func (pg *pg) GetByChainID(chainID int) ([]model.Token, error) {
	var tokens []model.Token
	return tokens, pg.db.Preload("Chain").Where("chain_id = ? AND discord_bot_supported = TRUE", chainID).Find(&tokens).Error
}

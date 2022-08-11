package guild_config_token

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

func (pg *pg) GetByGuildID(guildID string) ([]model.GuildConfigToken, error) {
	var configs []model.GuildConfigToken
	return configs, pg.db.Where("guild_id = ? AND active = TRUE", guildID).Preload("Token").Find(&configs).Error
}

func (pg *pg) UpsertMany(configs []model.GuildConfigToken) error {
	tx := pg.db.Begin()

	for _, config := range configs {
		err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "guild_id"}, {Name: "token_id"}},
			UpdateAll: true,
		}).Create(&config).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (pg *pg) CreateOne(record model.GuildConfigToken) error {
	return pg.db.Create(&record).Error
}

func (pg *pg) UpsertOne(configs model.GuildConfigToken) error {

	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "guild_id"}, {Name: "token_id"}},

		UpdateAll: true,
	}).Create(&configs).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) GetAll() ([]model.GuildConfigToken, error) {
	var guildConfigToken []model.GuildConfigToken
	return guildConfigToken, pg.db.Find(&guildConfigToken).Error
}

func (pg *pg) GetByGuildIDAndTokenID(guildID string, tokenID int) (*model.GuildConfigToken, error) {
	gct := &model.GuildConfigToken{}
	if err := pg.db.First(gct, "guild_id = ? AND token_id = ?", guildID, tokenID).Error; err != nil {
		return nil, err
	}
	return gct, nil
}

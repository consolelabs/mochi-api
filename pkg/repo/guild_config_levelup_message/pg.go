package guild_config_levelup_message

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

func (pg *pg) GetByGuildId(guildId string) (model *model.GuildConfigLevelupMessage, err error) {
	return model, pg.db.Where("guild_id = ?", guildId).First(&model).Error
}
func (pg *pg) DeleteByGuildId(guildId string) error {
	return pg.db.Where("guild_id = ?", guildId).Delete(model.GuildConfigLevelupMessage{}).Error
}
func (pg *pg) UpsertOne(config model.GuildConfigLevelupMessage) (*model.GuildConfigLevelupMessage, error) {
	tx := pg.db.Begin()

	// update on conflict
	err := tx.Table("guild_config_levelup_messages").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "guild_id"}},
		UpdateAll: true,
	}).Create(&config).Error
	if err != nil {
		tx.Rollback()
		return &config, err
	}

	return &config, tx.Commit().Error
}

package auto_trigger

import (
	"os"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) Update(triggerId string, field string, value interface{}) error {
	return pg.db.Model(&model.AutoCondition{}).Where("id = ?", triggerId).Update(field, value).Error
}

func (pq *pg) CountByGuildUser(guildId, userId string) (int64, error) {
	var count int64
	err := pq.db.Model(&model.AutoTrigger{}).Where("guild_id = ? AND user_id = ?", guildId, userId).Count(&count).Error
	return count, err
}

func (pg *pg) FirstOrCreate(autoTrigger *model.AutoTrigger) error {
	return pg.db.Where("guild_id = ? AND user_id = ?", autoTrigger.GuildId, autoTrigger.UserId).FirstOrCreate(autoTrigger).Error
}

func (pg *pg) GetAutoTriggers(guildId string) ([]model.AutoTrigger, error) {
	var result []model.AutoTrigger
	db := pg.db.Preload("Conditions", func(db *gorm.DB) *gorm.DB { return db.Order("auto_conditions.index") }).Preload("Conditions.ConditionValues", func(db *gorm.DB) *gorm.DB { return db.Order("auto_condition_values.index") })
	db = db.Preload("Conditions.ConditionValues.Type").Preload("Conditions.ChildConditions.Type").Preload("Conditions.Type").Preload("Actions", func(db *gorm.DB) *gorm.DB { return db.Order("auto_actions.index asc") })
	db = db.Preload("Actions.Type").Preload("Actions.Embed").Preload("Actions.Embed.Image").Preload("Actions.Embed.Video").Preload("Actions.Embed.Footer")
	db = db.Where("guild_id = ?", guildId)
	return result, db.Find(&result).Error
}

func (pg *pg) Create(autoTrigger *model.AutoTrigger) error {
	return pg.db.Create(autoTrigger).Error
}

func (pg *pg) UpsertMany(autoTriggers []model.AutoTrigger) error {
	log := logger.NewLogrusLogger()
	tx := pg.db.Begin()
	for _, at := range autoTriggers {
		err := tx.Clauses(
			clause.OnConflict{
				Columns: []clause.Column{{Name: "guild_id"}, {Name: "user_id"}},
				DoUpdates: []clause.Assignment{
					{Column: clause.Column{Name: "name"}, Value: at.Name},
					{Column: clause.Column{Name: "status"}, Value: at.Status},
				},
			}).Create(&at).Error
		if err != nil {
			log.Error(err, "[auto_triggers.UpsertMany] failed")
			os.Exit(0)
		}
	}
	return tx.Commit().Error
}

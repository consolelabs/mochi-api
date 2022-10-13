package conversation_repost_histories

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) *pg {
	return &pg{
		db: db,
	}
}

func (pg *pg) Upsert(model model.ConversationRepostHistories) error {
	tx := pg.db.Begin()

	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "guild_id"}, {Name: "origin_channel_id"}, {Name: "origin_start_message_id"}},
		UpdateAll: true,
	}).Create(&model).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) Update(model *model.ConversationRepostHistories) error {
	return pg.db.Model(&model).Where("id = ?", model.ID).Save(model).Error
}

func (pg *pg) GetByGuildAndChannel(guildID, channelID string) (*model.ConversationRepostHistories, error) {
	var model model.ConversationRepostHistories
	err := pg.db.Where("guild_id = ? AND origin_channel_id = ? and origin_start_message_id != '' and origin_stop_message_id = ''", guildID, channelID).First(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

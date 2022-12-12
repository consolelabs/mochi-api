package message_repost_history

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetByMessageID(guildID, messageID string) (*model.MessageRepostHistory, error) {
	var config model.MessageRepostHistory
	return &config, pg.db.Model(model.MessageRepostHistory{}).Where("guild_id = ? AND origin_message_id = ?", guildID, messageID).First(&config).Error
}

func (pg *pg) Upsert(record model.MessageRepostHistory) error {
	tx := pg.db.Begin()

	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "origin_message_id"}},
		UpdateAll: true,
	}).Create(&record).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) EditMessageRepost(req *request.EditMessageRepostRequest) error {
	return pg.db.Table("message_repost_histories").Where("guild_id=? AND repost_channel_id=? AND origin_channel_id=? AND origin_message_id=?", req.GuildID, req.RepostChannelID, req.OriginChannelID, req.OriginMessageID).Update("repost_message_id", req.RepostMessageID).Error
}

package guild_users

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) Update(guildId, userId int64, field string, value interface{}) error {
	return pg.db.Model(&model.GuildUser{}).Where("guild_id = ? AND user_id = ?", guildId, userId).Update(field, value).Error
}

func (pq *pg) CountByGuildUser(guildId, userId int64) (int64, error) {
	var count int64
	err := pq.db.Model(&model.GuildUser{}).Where("guild_id = ? AND invited_by = ?", guildId, userId).Count(&count).Error
	return count, err
}

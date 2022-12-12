package guild_users

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) Update(guildId, userId string, field string, value interface{}) error {
	return pg.db.Model(&model.GuildUser{}).Where("guild_id = ? AND user_id = ?", guildId, userId).Update(field, value).Error
}

func (pq *pg) CountByGuildUser(guildId, userId string) (int64, error) {
	var count int64
	err := pq.db.Model(&model.GuildUser{}).Where("guild_id = ? AND invited_by = ?", guildId, userId).Count(&count).Error
	return count, err
}

func (pg *pg) FirstOrCreate(guildUser *model.GuildUser) error {
	return pg.db.Where("guild_id = ? AND user_id = ?", guildUser.GuildID, guildUser.UserID).FirstOrCreate(guildUser).Error
}

func (pg *pg) GetGuildUsers(guildID string) ([]model.GuildUser, error) {
	var result []model.GuildUser
	return result, pg.db.Where("guild_id = ?", guildID).Find(&result).Error
}

func (pg *pg) Create(guildUser *model.GuildUser) error {
	return pg.db.Create(guildUser).Error
}

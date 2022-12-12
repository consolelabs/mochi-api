package guild_user_xp

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

func (pg *pg) GetOne(guildID, userID string) (*model.GuildUserXP, error) {
	userXP := &model.GuildUserXP{}
	return userXP, pg.db.Where("guild_id = ? AND user_id = ?", guildID, userID).Preload("Guild").First(userXP).Error
}

func (pg *pg) GetByGuildID(guildID string) ([]model.GuildUserXP, error) {
	var result []model.GuildUserXP
	return result, pg.db.Where("guild_id = ?", guildID).Find(&result).Error
}

func (pg *pg) GetTopUsers(guildID string, limit, offset int) ([]model.GuildUserXP, error) {
	var userXPs []model.GuildUserXP
	return userXPs, pg.db.Where("guild_id = ?", guildID).Preload("User").Offset(offset).Limit(limit).Order("guild_rank").Find(&userXPs).Error
}

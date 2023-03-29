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

func (pg *pg) GetTopUsers(guildID, query, sort string, limit, offset int) ([]model.GuildUserXP, error) {
	var userXPs []model.GuildUserXP
	q := pg.db.Where("guild_id = ?", guildID).Preload("User").Preload("User.GuildUsers", "guild_id=?", guildID).Offset(offset).Limit(limit)

	if query != "" {
		q = q.Where("username LIKE ? OR nickname LIKE ?", "%"+query+"%", "%"+query+"%")
	}

	if sort != "" {
		q = q.Order("total_xp " + sort)
	} else {
		q = q.Order("guild_rank")
	}

	return userXPs, q.Find(&userXPs).Error
}

func (pg *pg) GetTotalTopUsersCount(guildID, query string) (int64, error) {
	var count int64
	q := pg.db.Model(&model.GuildUserXP{}).Where("guild_id = ?", guildID)

	if query != "" {
		q = q.Where("username LIKE ? OR nickname LIKE ?", "%"+query+"%", "%"+query+"%")
	}

	return count, q.Count(&count).Error
}

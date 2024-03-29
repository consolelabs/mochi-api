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

func (pg *pg) GetOne(q GetOneQuery) (*model.GuildUserXP, error) {
	userXP := &model.GuildUserXP{}
	db := pg.db.Where("guild_id = ?", q.GuildID)
	if q.ProfileID != "" {
		db = db.Where("profile_id = ?", q.ProfileID)
	}
	if q.UserID != "" {
		db = db.Where("user_id = ?", q.UserID)
	}
	return userXP, db.Preload("Guild").First(userXP).Error
}

func (pg *pg) GetByGuildID(guildID string) ([]model.GuildUserXP, error) {
	var result []model.GuildUserXP
	return result, pg.db.Where("guild_id = ?", guildID).Find(&result).Error
}

func (pg *pg) GetTopUsers(guildID, query, sort string, limit, offset int) ([]model.GuildUserXP, error) {
	var userXPs []model.GuildUserXP
	q := pg.db.Where("guild_id = ?", guildID).Preload("User").Preload("User.GuildUsers", "guild_id=?", guildID).Offset(offset).Limit(limit)

	if query != "" {
		q = q.Where("username ILIKE ? OR nickname ILIKE ?", "%"+query+"%", "%"+query+"%")
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
		q = q.Where("username ILIKE ? OR nickname ILIKE ?", "%"+query+"%", "%"+query+"%")
	}

	return count, q.Count(&count).Error
}

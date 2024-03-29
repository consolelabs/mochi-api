package guild_users

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

func (pg *pg) UpsertMany(guildUsers []model.GuildUser) error {
	log := logger.NewLogrusLogger()
	tx := pg.db.Begin()
	for _, gu := range guildUsers {
		err := tx.Clauses(
			clause.OnConflict{
				Columns: []clause.Column{{Name: "guild_id"}, {Name: "user_id"}},
				DoUpdates: []clause.Assignment{
					{Column: clause.Column{Name: "invited_by"}, Value: gu.InvitedBy},
					{Column: clause.Column{Name: "nickname"}, Value: gu.Nickname},
					{Column: clause.Column{Name: "avatar"}, Value: gu.Avatar},
					{Column: clause.Column{Name: "joined_at"}, Value: gu.JoinedAt},
					{Column: clause.Column{Name: "roles"}, Value: gu.Roles},
				},
			}).Create(&gu).Error
		if err != nil {
			log.Error(err, "[guild_users.UpsertMany] failed")
			os.Exit(0)
		}
	}
	return tx.Commit().Error
}

func (pg *pg) GetUsersOfGuild(ids []string, guildId string) (res []model.GuildUser, err error) {
	return res, pg.db.Where("guild_id = ? AND user_id IN ?", guildId, ids).Find(&res).Error
}

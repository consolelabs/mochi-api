package users

import (
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) Upsert(user *model.User) error {
	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"username"}),
	}).Create(user).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) UpsertMany(users []model.User) error {
	log := logger.NewLogrusLogger()
	tx := pg.db.Begin()
	for _, user := range users {
		err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"username", "discriminator"}),
		}).Create(&user).Error
		if err != nil {
			log.Error(err, "[users.UpsertMany] failed")
			continue
		}
	}
	return tx.Commit().Error
}

func (pg *pg) GetOne(discordID string) (*model.User, error) {
	u := &model.User{}
	return u, pg.db.Where("id = ?", discordID).Preload("GuildUsers").First(u).Error
}

func (pg *pg) GetByDiscordIDs(discordIDs []string) ([]model.User, error) {
	users := []model.User{}
	return users, pg.db.Where("id IN (?)", discordIDs).Find(&users).Error
}

func (pg *pg) UpdateNrOfJoin(discordId string, nrOfJoin int64) error {
	u := &model.User{}
	return pg.db.Model(u).Where("id = ?", discordId).Update("nr_of_join", nrOfJoin).Error
}

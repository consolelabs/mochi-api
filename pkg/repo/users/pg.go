package users

import (
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

func (pg *pg) Create(user *model.User) error {
	return pg.db.Create(user).Error
}

func (pg *pg) GetLatestWalletNumber() int {
	var result int
	row := pg.db.Table("users").Select("max(in_discord_wallet_number)").Row()
	row.Scan(&result)
	return result
}

func (pg *pg) UpsertOne(user model.User) error {
	tx := pg.db.Begin()
	cols := []clause.Column{{Name: "id"}}

	err := tx.Omit(clause.Associations).Clauses(clause.OnConflict{
		Columns:   cols,
		UpdateAll: true,
	}).Create(&user).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) GetOne(discordID string) (*model.User, error) {
	u := &model.User{}
	return u, pg.db.Where("id = ?", discordID).First(u).Error
}

func (pg *pg) GetByDiscordIDs(discordIDs []string) ([]model.User, error) {
	users := []model.User{}
	return users, pg.db.Where("id IN (?)", discordIDs).Find(&users).Error
}

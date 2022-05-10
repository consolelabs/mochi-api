package users

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

func (pg *pg) Create(user *model.User) error {
	return pg.db.Create(user).Error
}

func (pg *pg) FirstOrCreate(user *model.User) error {
	return pg.db.Where("id = ?", user.ID).FirstOrCreate(user).Error
}

func (pg *pg) GetLatestWalletNumber() int {
	var result int
	row := pg.db.Table("users").Select("max(in_discord_wallet_number)").Row()
	row.Scan(&result)
	return result
}

func (pg *pg) GetOne(discordID string) (*model.User, error) {
	u := &model.User{}
	return u, pg.db.Where("id = ?", discordID).Preload("GuildUsers").First(u).Error
}

func (pg *pg) GetByDiscordIDs(discordIDs []string) ([]model.User, error) {
	users := []model.User{}
	return users, pg.db.Where("id IN (?)", discordIDs).Find(&users).Error
}

func (pg *pg) Update(u *model.User) error {
	return pg.db.Save(u).Error
}

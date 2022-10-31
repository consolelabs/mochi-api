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

func (pg *pg) Upsert(user *model.User) error {
	tx := pg.db.Begin()
	onConflict := clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoNothing: true,
	}
	if user.InDiscordWalletAddress.String != "" {
		onConflict.DoNothing = false
		onConflict.DoUpdates = clause.AssignmentColumns([]string{"in_discord_wallet_address", "in_discord_wallet_number"})
	}
	if user.Username != "" {
		onConflict.DoNothing = false
		onConflict.DoUpdates = append(onConflict.DoUpdates, clause.AssignmentColumns([]string{"username"})...)
	}
	err := tx.Clauses(onConflict).Create(user).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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

func (pg *pg) UpdateNrOfJoin(discordId string, nrOfJoin int64) error {
	u := &model.User{}
	return pg.db.Model(u).Where("id = ?", discordId).Update("nr_of_join", nrOfJoin).Error
}

func (pg *pg) UpdateUserIsMigrateBals(userID string) error {
	u := &model.User{}
	return pg.db.Model(u).Where("id = ?", userID).Update("is_migrate_bal", true).Error

}

func (pg *pg) Get50Users() ([]model.User, error) {
	users := []model.User{}
	return users, pg.db.Limit(50).Where("in_discord_wallet_address IS NOT NULL AND in_discord_wallet_address != '' and username = 'digdug'").Find(&users).Error
}

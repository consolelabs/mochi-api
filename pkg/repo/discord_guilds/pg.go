package discord_guilds

import (
	"fmt"

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

func (pg *pg) Gets() ([]model.DiscordGuild, error) {
	guilds := []model.DiscordGuild{}
	err := pg.db.Preload("GuildConfigInviteTracker").Find(&guilds).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get guilds: %w", err)
	}

	return guilds, nil
}

func (pg *pg) CreateIfNotExists(guild model.DiscordGuild) error {
	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "id"},
		},
		DoNothing: true,
	}).Create(&guild).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (pg *pg) GetByID(id string) (*model.DiscordGuild, error) {
	var guild model.DiscordGuild
	return &guild, pg.db.Preload("GuildConfigInviteTracker").First(&guild, "id = ?", id).Error
}

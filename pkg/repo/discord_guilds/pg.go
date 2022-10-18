package discord_guilds

import (
	"fmt"
	"time"

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
	err := pg.db.Where("active = TRUE").Preload("GuildConfigInviteTracker").Find(&guilds).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get guilds: %w", err)
	}

	return guilds, nil
}

func (pg *pg) CreateOrReactivate(guild model.DiscordGuild) error {
	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "id"},
		},
		DoUpdates: clause.Assignments(map[string]interface{}{"active": true, "joined_at": time.Now(), "left_at": nil}),
	}).Create(&guild).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (pg *pg) GetByID(id string) (*model.DiscordGuild, error) {
	var guild model.DiscordGuild
	return &guild, pg.db.Preload("GuildConfigInviteTracker").Where("active = TRUE").First(&guild, "id = ?", id).Error
}

func (pg *pg) ToggleGlobalXP(guildID string, globalXP bool) error {
	return pg.db.Model(&model.DiscordGuild{}).Where("id = ?", guildID).Update("global_xp", globalXP).Error
}

// func (pg *pg) Update(omit string, guild model.DiscordGuild) error {
// 	return pg.db.Model(&guild).Where("id = ?", guild.ID).Omit(omit).Updates(map[string]interface{}{"global_xp": guild.GlobalXP, "log_channel": guild.LogChannel}).Error
// }

func (pg *pg) Update(guild *model.DiscordGuild) error {
	return pg.db.Model(&guild).Where("id = ?", guild.ID).Save(guild).Error
}

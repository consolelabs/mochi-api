package guildcustomcommand

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

func (pg *pg) GetAll(q GetAllQuery) ([]model.GuildCustomCommand, error) {
	var commands []model.GuildCustomCommand

	tx := pg.db

	if q.GuildID != "" {
		tx = tx.Where("guild_id = ?", q.GuildID)
	}

	if q.Enabled != nil {
		tx = tx.Where("enabled = ?", *q.Enabled)
	}

	return commands, tx.Find(&commands).Error
}

func (pg *pg) GetByIDAndGuildID(ID, guildID string) (*model.GuildCustomCommand, error) {
	var command model.GuildCustomCommand
	return &command, pg.db.Where("id = ? and guild_id = ?", ID, guildID).First(&command).Error
}

func (pg *pg) UpsertOne(command model.GuildCustomCommand) error {
	tx := pg.db.Begin()

	err := tx.Omit(clause.Associations).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}, {Name: "guild_id"}},
		UpdateAll: true,
	}).Create(&command).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) Update(ID, guilID string, command model.GuildCustomCommand) error {
	return pg.db.Model(&model.GuildCustomCommand{}).Where("id = ? and guild_id = ?", ID, guilID).Updates(command).Error
}

func (pg *pg) Delete(command model.GuildCustomCommand) error {
	return pg.db.Delete(&command).Error
}

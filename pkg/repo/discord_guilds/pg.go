package discord_guilds

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

func (pg *pg) Gets() ([]*model.DiscordGuild, error) {
	var guilds []*model.DiscordGuild
	return guilds, pg.db.Preload("LogChannel").Find(&guilds).Error
}

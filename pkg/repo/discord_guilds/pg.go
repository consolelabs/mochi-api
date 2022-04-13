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
	err := pg.db.Find(&guilds).Error
	if err != nil {
		return nil, err
	}
	return guilds, nil
}

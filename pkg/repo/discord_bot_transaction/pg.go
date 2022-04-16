package discordbottransaction

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

func (pg *pg) Create(param model.DiscordBotTransaction) (*model.DiscordBotTransaction, error) {
	return &param, pg.db.Table("discord_bot_transactions").Create(&param).Error
}

func (pg *pg) Delete(id string) error {
	return pg.db.Delete(&model.DiscordBotTransaction{}, id).Error
}

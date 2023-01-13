package guild_config_token_role

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

func (pg *pg) Create(config *model.GuildConfigTokenRole) error {
	return pg.db.Create(config).Error
}

func (pg *pg) Get(id int) (model *model.GuildConfigTokenRole, err error) {
	return model, pg.db.First(&model, id).Error
}

func (pg *pg) ListByGuildID(guildID string) ([]model.GuildConfigTokenRole, error) {
	var configs []model.GuildConfigTokenRole
	return configs, pg.db.Preload("Token").Where("guild_id = ?", guildID).Order("token_id, required_amount asc").Find(&configs).Error
}

func (pg *pg) Update(config *model.GuildConfigTokenRole) error {
	return pg.db.Save(config).Error
}

func (pg *pg) Delete(id int) error {
	return pg.db.Delete(&model.GuildConfigTokenRole{}, "id = ?", id).Error
}

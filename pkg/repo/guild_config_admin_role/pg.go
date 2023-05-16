package guild_config_admin_role

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) *pg {
	return &pg{
		db: db,
	}
}

func (pg *pg) Create(config []model.GuildConfigAdminRole) error {
	return pg.db.Create(&config).Error
}

func (pg *pg) ListByGuildID(guildID string) (config []model.GuildConfigAdminRole, err error) {
	return config, pg.db.Where("guild_id = ?", guildID).Order("created_at desc").Find(&config).Error
}

func (pg *pg) Delete(id int) error {
	return pg.db.Delete(&model.GuildConfigAdminRole{}, "id = ?", id).Error
}

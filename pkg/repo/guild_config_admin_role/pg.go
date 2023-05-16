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
	tx := pg.db.Begin()
	for _, item := range config {
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (pg *pg) ListByGuildID(guildID string) ([]model.GuildConfigAdminRole, error) {
	var config []model.GuildConfigAdminRole
	return config, pg.db.Where("guild_id = ?", guildID).Order("created_at desc").Find(&config).Error
}

func (pg *pg) Delete(id int) error {
	return pg.db.Delete(&model.GuildConfigAdminRole{}, "id = ?", id).Error
}

package commandpermission

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) List(q ListQuery) (cmds []model.CommandPermission, err error) {
	db := pg.db
	if q.Code != "" {
		db = db.Where("lower(code) = lower(?)", q.Code)
	}

	return cmds, db.Find(&cmds).Error
}

func (pg *pg) ListUniqueDiscordPermission() (pers []string, err error) {
	db := pg.db.
		Table("command_permissions").
		Select("DISTINCT ON (discord_permission_flag) discord_permission_flag")
	return pers, db.Find(&pers).Error
}

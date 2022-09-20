package serversusagestats

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

func (pg *pg) CreateOne(info *model.UsageStat) error {
	return pg.db.Table("servers_usage_stats").Create(&info).Error
}

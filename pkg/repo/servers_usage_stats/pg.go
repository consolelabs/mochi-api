package serversusagestats

import (
	"github.com/defipod/mochi/pkg/request"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) CreateOne(info *request.UsageInformation) error {
	return pg.db.Table("servers_usage_stats").Create(&info).Error
}

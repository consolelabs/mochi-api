package product_changelogs_view

import (
	"gorm.io/gorm"
	"strings"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) List(q ListQuery) (changeLogsView []model.ProductChangelogView, err error) {
	db := pg.db
	if q.Key != "" {
		db = db.Where("lower(key) = ?", strings.ToLower(q.Key))
	}
	if q.ChangelogName != "" {
		db = db.Where("lower(changelog_name) = ?", strings.ToLower(q.ChangelogName))
	}

	return changeLogsView, db.Order("created_at DESC").Find(&changeLogsView).Error
}

func (pg *pg) Create(changeLogsView *model.ProductChangelogView) error {
	db := pg.db
	return db.Where("key = ? and changelog_name = ?", changeLogsView.Key, changeLogsView.ChangelogName).FirstOrCreate(changeLogsView).Error
}

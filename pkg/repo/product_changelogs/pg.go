package product_changelogs

import (
	"strings"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) List(q ListQuery) (changeLogs []model.ProductChangelogs, total int64, err error) {
	db := pg.db.Model(&model.ProductChangelogs{})
	if q.Product != "" {
		db = db.Where("lower(product) = ?", strings.ToLower(q.Product))
	}
	return changeLogs, total, db.
		Count(&total).
		Offset(q.Page * q.Size).
		Limit(q.Size).
		Order("created_at DESC").Find(&changeLogs).Error
}

func (pg *pg) Create(changelog *model.ProductChangelogs) error {
	db := pg.db
	return db.Create(changelog).Error
}

func (pg *pg) DeleteAll() error {
	db := pg.db
	return db.Where("title != ''").Delete(model.ProductChangelogs{}).Error
}

func (pg *pg) GetNewChangelog() (changeLogs []model.ProductChangelogs, err error) {
	db := pg.db
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return changeLogs, err
	}
	tx.Raw(`Select * from product_changelogs as pc 
    				left join product_changelog_snapshots as pcs 
    				on pc.file_name = pcs.filename 
         			where pcs.filename is null`).Scan(&changeLogs)
	return changeLogs, nil
}

func (pg *pg) InsertBulkProductChangelogSnapshot(changelogSnapshot []model.ProductChangelogSnapshot) error {
	db := pg.db
	return db.Create(changelogSnapshot).Error
}

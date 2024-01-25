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

func (pg pg) GetChangelogByFilename(filename string) (p model.ProductChangelogs, err error) {
	db := pg.db
	return p, db.Where("file_name = ?", filename).First(&p).Error
}

func (pg pg) GetChangelogSnapshotByFilename(filename string) (p model.ProductChangelogSnapshot, err error) {
	db := pg.db
	return p, db.Where("filename = ?", filename).First(&p).Error
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

func (pg *pg) GetChangelogNotConfirmed() (changeLogs []model.ProductChangelogs, err error) {
	db := pg.db
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return changeLogs, err
	}
	tx.Raw(`Select * from product_changelogs as pc 
    				left join product_changelog_snapshots as pcs 
    				on pc.file_name = pcs.filename
         			where pcs.is_public = false`).Scan(&changeLogs)
	return changeLogs, nil
}

func (pg *pg) UpdateProductChangelogSnapshot(q ProductChangelogSnapshotQuery) error {
	p := &model.ProductChangelogSnapshot{}
	return pg.db.Model(p).Where("filename = ?", q.Filename).Update("is_public", q.IsPublic).Error
}

func (pg *pg) InsertBulkProductChangelogSnapshot(changelogSnapshot []model.ProductChangelogSnapshot) error {
	db := pg.db
	return db.Create(changelogSnapshot).Error
}

func (pg *pg) GetByVersion(version string) (*model.ProductChangelogs, error) {
	changelog := &model.ProductChangelogs{}
	db := pg.db.Where("version = ?", version)
	return changelog, db.First(changelog).Error
}

func (pg *pg) GetNextVersion(id int64) (string, error) {
	changelog := &model.ProductChangelogs{}
	db := pg.db.Where("id > ?", id).Order("id DESC").Limit(1)
	err := db.First(&changelog).Error
	if err != nil {
		return "", err
	}
	return changelog.Version, nil
}

func (pg *pg) GetPreviousVersion(id int64) (string, error) {
	changelog := &model.ProductChangelogs{}
	db := pg.db.Where("id < ?", id).Order("id DESC").Limit(1)
	err := db.First(&changelog).Error
	if err != nil {
		return "", err
	}
	return changelog.Version, nil
}

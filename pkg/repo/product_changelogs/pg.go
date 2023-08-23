package product_changelogs

import (
	"gorm.io/gorm"
	"strconv"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) List(q ListQuery) (changeLogs []model.ProductChangelogs, err error) {
	db := pg.db
	if q.Product != "" {
		product, err := strconv.Atoi(q.Product)
		if err != nil {
			product = 0
		}
		db = db.Where("product = ?", product)
	}
	if q.Size != "" {
		size, err := strconv.Atoi(q.Size)
		if err != nil {
			size = 1
		}
		db.Limit(size)
	}

	return changeLogs, db.Order("created_at DESC").Find(&changeLogs).Error
}

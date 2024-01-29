package tokenpricesnapshot

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) UpsertOne(walletSnapshot *model.TokenPriceSnapshot) error {
	tx := pg.db.Table("token_price_snapshot").Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "symbol"},
		},
		DoUpdates: clause.AssignmentColumns([]string{"price", "symbol", "snapshot_time", "updated_at"}),
	}).Create(&walletSnapshot).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

}

func (pg *pg) GetLatestSnapshotWithTime(symbol string, time time.Time) (snapshotPrice float64, err error) {
	return snapshotPrice, pg.db.Table("token_price_snapshot").Model(model.TokenPriceSnapshot{}).Select("price").Where("lower(symbol) = lower(?) and snapshot_time > ?", symbol, time).Order("snapshot_time desc").First(&snapshotPrice).Error
}

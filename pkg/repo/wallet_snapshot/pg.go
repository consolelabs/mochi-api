package walletsnapshot

import (
	"time"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) Create(walletSnapshot *model.WalletSnapshot) (*model.WalletSnapshot, error) {
	return walletSnapshot, pg.db.Table("wallet_snapshot").Create(walletSnapshot).Error
}

func (pg *pg) GetSnapshotInTime(address string, time time.Time) (snapshots []model.WalletSnapshot, err error) {
	return snapshots, pg.db.Table("wallet_snapshot").Select("total_usd_balance").Where("wallet_address = ? and snapshot_time >= ?", address, time).Order("snapshot_time asc").Find(&snapshots).Error
}

func (pg *pg) GetLatestInPast(address string, time time.Time) (snapshots []model.WalletSnapshot, err error) {
	return snapshots, pg.db.Table("wallet_snapshot").Select("total_usd_balance").Where("wallet_address = ? and snapshot_time < ?", address, time).Order("snapshot_time desc").Limit(20).Find(&snapshots).Error
}

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
	return walletSnapshot, pg.db.Create(walletSnapshot).Error
}

func (pg *pg) GetSnapshotInTime(address string, time time.Time) (snapshots []model.WalletSnapshot, err error) {
	return snapshots, pg.db.Where("wallet_address = ? and snapshot_time >= ?", address, time).Order("desc snapshot_time").Find(&snapshots).Error
}

func (pg *pg) GetLatestInPast(address string, time time.Time) (snapshots []model.WalletSnapshot, err error) {
	return snapshots, pg.db.Where("wallet_address = ? and snapshot_time < ?", address, time).Order("desc snapshot_time").Find(&snapshots).Error
}

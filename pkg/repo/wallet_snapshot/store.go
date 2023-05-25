package walletsnapshot

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	Create(walletSnapshot *model.WalletSnapshot) (*model.WalletSnapshot, error)
	GetSnapshotInTime(address string, time time.Time) (snapshots []model.WalletSnapshot, err error)
	GetLatestInPast(address string, time time.Time) (snapshots []model.WalletSnapshot, err error)
}

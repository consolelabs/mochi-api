package tokenpricesnapshot

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	UpsertOne(walletSnapshot *model.TokenPriceSnapshot) error
	GetLatestSnapshotWithTime(id string, time time.Time) (snapshotPrice float64, err error)
}

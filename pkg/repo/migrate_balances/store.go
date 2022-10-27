package migratebalances

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	StoreMigrateBalances(*model.MigrateBalance) error
}

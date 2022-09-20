package serversusagestats

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	CreateOne(info *model.UsageStat) error
}

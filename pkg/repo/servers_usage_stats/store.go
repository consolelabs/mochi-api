package serversusagestats

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	CreateOne(info *model.UsageStat) error
	TotalUsage() (count int64, err error)
	TotalUsageByGuildId(guildId string) (count int64, err error)
}

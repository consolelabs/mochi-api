package profilecommandusage

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetTopProfileUsage(top int) ([]model.CommandUsageCounter, error)
}

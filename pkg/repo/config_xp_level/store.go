package config_xp_level

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetNextLevel(xp int, next bool) (*model.ConfigXpLevel, error)
	GetLevelInfo(xp int) (*model.ConfigXpLevel, *model.ConfigXpLevel, error)
}

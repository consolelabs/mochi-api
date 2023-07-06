package config_xp_level

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetNextLevel(xp int, next bool) (*model.ConfigXpLevel, error) {
	config := &model.ConfigXpLevel{}
	where := "min_xp <= ?"
	order := "level DESC"
	if next {
		where = "min_xp > ?"
		order = "level ASC"
	}
	return config, pg.db.Where(where, xp).Order(order).First(config).Error
}

func (pg *pg) GetLevelInfo(xp int) (*model.ConfigXpLevel, *model.ConfigXpLevel, error) {
	currentLevel := &model.ConfigXpLevel{}
	nextLevel := &model.ConfigXpLevel{}
	err := pg.db.Where("min_xp <= ?", xp).Order("level DESC").First(currentLevel).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, nil, err
		}
		// Handle ErrRecordNotFound as a special case
	}

	err = pg.db.Where("min_xp > ?", xp).Order("level ASC").First(nextLevel).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, nil, err
		}
		// Handle ErrRecordNotFound as a special case
	}

	return currentLevel, nextLevel, nil
}

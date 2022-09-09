package upvotestreaktier

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (pg *pg) GetByUpvoteCount(upvote int) (*model.UpvoteStreakTier, error) {
	tier := model.UpvoteStreakTier{}
	return &tier, pg.db.Where("streak_required <= ?", upvote).Last(&tier).Error
}

func (pg *pg) GetByID(tierID int) (*model.UpvoteStreakTier, error) {
	tier := model.UpvoteStreakTier{}
	return &tier, pg.db.Where("id = ?", tier).First(&tier).Error
}

func (pg *pg) GetAll() ([]model.UpvoteStreakTier, error) {
	tiers := []model.UpvoteStreakTier{}
	return tiers, pg.db.Find(&tiers).Error
}

package upvotestreaktier

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	GetByUpvoteCount(upvote int) (*model.UpvoteStreakTier, error)
	GetByID(tierID int) (*model.UpvoteStreakTier, error)
	GetAll() ([]model.UpvoteStreakTier, error)
}

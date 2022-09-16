package job

import (
	"time"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
)

type updateUpvoteStreak struct {
	entity *entities.Entity
	log    logger.Logger
}

func NewUpdateUpvoteStreakJob(e *entities.Entity, l logger.Logger) Job {
	return &updateUpvoteStreak{
		entity: e,
		log:    l,
	}
}

func (c *updateUpvoteStreak) Run() error {
	c.log.Info("start updating upvote streak")
	streaks, err := c.entity.GetAllUpvoteStreak()
	if err != nil {
		c.log.Error(err, "entity.GetAllUpvoteStreak failed")
		return err
	}
	expiredStreak := []model.DiscordUserUpvoteStreak{}
	for _, streak := range streaks {
		expireTime := streak.LastStreakDate.Add(time.Hour * 24)
		if time.Now().After(expireTime) {
			streak.StreakCount = 0
			expiredStreak = append(expiredStreak, streak)
		}
	}
	err = c.entity.UpsertBatchUpvoteStreak(expiredStreak)
	if err != nil {
		c.log.Error(err, "entity.UpsertBatchUpvoteStreak failed")
	}
	return nil
}

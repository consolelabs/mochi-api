package job

import (
	"time"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
)

type updateGMStreak struct {
	entity *entities.Entity
	log    logger.Logger
}

func NewUpdateGMStreakJob(e *entities.Entity, l logger.Logger) Job {
	return &updateGMStreak{
		entity: e,
		log:    l,
	}
}

func (c *updateGMStreak) Run() error {
	c.log.Info("start updating gm streak")
	streaks, err := c.entity.GetAllGMStreak()
	if err != nil {
		c.log.Error(err, "entity.GetAllGMStreak failed")
		return err
	}
	expiredStreak := []model.DiscordUserGMStreak{}
	for _, streak := range streaks {
		expireTime := streak.LastStreakDate.Add(time.Hour * 24)
		if time.Now().After(expireTime) {
			streak.StreakCount = 0
			expiredStreak = append(expiredStreak, streak)
		}
	}
	err = c.entity.UpsertBatchGMStreak(expiredStreak)
	if err != nil {
		c.log.Error(err, "entity.UpsertBatchGMStreak failed")
	}
	return nil
}

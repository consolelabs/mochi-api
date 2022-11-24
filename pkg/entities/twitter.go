package entities

import (
	"time"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	twitterpoststreak "github.com/defipod/mochi/pkg/repo/twitter_post_streak"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) GetUnnotifiedSalesMessage() ([]model.TwitterSalesMessage, error) {
	messages, err := e.repo.MochiNFTSales.GetAllUnnotified()
	if err != nil {
		e.log.Errorf(err, "[e.HandleMochiSalesMessage] failed to get mochi nft sales: %s", err)
		return nil, err
	}
	return messages, nil
}

func (e *Entity) DeleteSalesMessages(message *model.TwitterSalesMessage) error {
	err := e.repo.MochiNFTSales.DeleteOne(message)
	if err != nil {
		e.log.Errorf(err, "[e.HandleMochiSalesMessage] failed to update mochi nft sales: %s", err)
		return err
	}
	return nil
}

func (e *Entity) CreateTwitterPost(req *request.TwitterPost) error {
	err := e.repo.TwitterPost.CreateOne(&model.TwitterPost{
		TwitterID:     req.TwitterID,
		TwitterHandle: req.TwitterHandle,
		TweetID:       req.TweetID,
		GuildID:       req.GuildID,
		Content:       req.Content,
	})
	if err != nil {
		e.log.Error(err, "[entity.CreateTwitterPost] repo.TwitterPost.CreateOne() failed")
		return err
	}

	err = e.updateTwitterPostStreak(req)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.CreateTwitterPost] entity.updateTwitterPostStreak() failed")
	}
	return nil
}

func (e *Entity) updateTwitterPostStreak(req *request.TwitterPost) error {
	today := util.StartOfDay(time.Now().UTC())
	q := twitterpoststreak.ListQuery{GuildID: req.GuildID, TwitterID: req.TwitterID}
	streaks, _, err := e.repo.TwitterPostStreak.List(q)
	if err != nil {
		e.log.Fields(logger.Fields{"query": q}).Error(err, "[entity.updateTwitterPostStreak] repo.TwitterPostStreak.List() failed")
		return err
	}
	// create if not exists
	if len(streaks) == 0 {
		newStreak := model.TwitterPostStreak{
			GuildID:        req.GuildID,
			TwitterID:      req.TwitterID,
			TwitterHandle:  req.TwitterHandle,
			StreakCount:    1,
			TotalCount:     1,
			LastStreakDate: today,
		}
		err = e.repo.TwitterPostStreak.UpsertOne(&newStreak)
		if err != nil {
			e.log.Fields(logger.Fields{"newStreak": newStreak}).Error(err, "[entity.updateTwitterPostStreak] repo.TwitterPostStreak.UpsertOne() failed")
		}
		return err
	}

	currentStreak := streaks[0]
	nextStreakDate := util.StartOfDay(currentStreak.LastStreakDate.Add(time.Hour * 24))
	switch {
	// today == last streak date => already update streak today => skip
	case currentStreak.LastStreakDate.Equal(today):
		return nil
		// today - last streak date = 1 day => streak maintains
	case nextStreakDate.Equal(today):
		currentStreak.StreakCount++
		// today - last streak date > 1 day => broken streak => reset
	case nextStreakDate.Before(today):
		currentStreak.StreakCount = 1
	}
	currentStreak.LastStreakDate = today
	currentStreak.TotalCount++
	if err := e.repo.TwitterPostStreak.UpsertOne(&currentStreak); err != nil {
		e.log.Fields(logger.Fields{"streak": *&currentStreak}).Error(err, "[entity.updateTwitterPostStreak] repo.TwitterPostStreak.UpsertOne() failed")
	}
	return err
}

func (e *Entity) GetTwitterLeaderboard(req request.GetTwitterLeaderboardRequest) ([]model.TwitterPostStreak, int64, error) {
	q := twitterpoststreak.ListQuery{
		GuildID: req.GuildID,
		Sort:    "total_count DESC",
		Limit:   int(req.Size),
		Offset:  int(req.Size) * int(req.Page),
	}
	list, total, err := e.repo.TwitterPostStreak.List(q)
	if err != nil {
		e.log.Fields(logger.Fields{"query": q}).Error(err, "[entity.GetTwitterLeaderboard] repo.TwitterPostStreak.List() failed")
	}
	// if streak is broken, set to 0
	startOfYesterday := util.StartOfDay(time.Now().UTC().AddDate(0, 0, -1))
	for i, s := range list {
		if s.LastStreakDate.Before(startOfYesterday) {
			list[i].StreakCount = 0
		}
	}
	return list, total, err
}

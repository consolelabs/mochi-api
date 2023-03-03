package entities

import (
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) ListCommonwealthDiscussionSubscription(discussionID *int64) ([]model.CommonwealthDiscussionSubscription, error) {
	return e.repo.CommonwealthDiscussionSubscription.List(discussionID)
}

func (e *Entity) CreateCommonwealthDiscussionSubscription(req request.CreateCommonwealthDiscussionSubscription) (*model.CommonwealthDiscussionSubscription, error) {
	sub := &model.CommonwealthDiscussionSubscription{
		DiscordThreadID: req.DiscordThreadID,
		DiscussionID:    req.DiscussionID,
	}
	if err := e.repo.CommonwealthDiscussionSubscription.Create(sub); err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "repo.CommonwealthDiscussionSubscription.Create failed")
		return nil, err
	}
	return sub, nil
}

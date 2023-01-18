package entities

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) GetUserEnvelopStreak(userID string) (*model.UserEnvelopStreak, error) {
	streak, err := e.repo.Envelop.GetUserStreak(userID)
	if err != nil {
		e.log.Fields(logger.Fields{
			"userID": userID,
		}).Error(err, "[entity.GetUserEnvelopStreak] repo.envelop.GetUserStreak failed")
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrRecordNotFound
		}
		return nil, err
	}
	return streak, nil
}

func (e *Entity) CreateEnvelop(req request.CreateEnvelop) (*model.Envelop, error) {
	envelop := &model.Envelop{
		UserID:  req.UserID,
		Command: req.Command,
	}
	if err := e.repo.Envelop.Create(envelop); err != nil {
		e.log.Fields(logger.Fields{
			"envelop": envelop,
		}).Error(err, "[entity.CreateEnvelop] repo.Envelop.Create failed")
		return nil, err
	}
	return envelop, nil
}

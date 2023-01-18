package entities

import (
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

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

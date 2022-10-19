package entities

import (
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) HandleUserFeedback(req *request.UserFeedbackRequest) error {
	err := e.svc.Discord.SendFeedback(req)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.HandleUserFeedback] e.svc.Discord.SendFeedback failed")
		return err
	}
	return nil
}

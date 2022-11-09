package entities

import (
	"fmt"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) HandleUserFeedback(req *request.UserFeedbackRequest) error {
	// store feedback
	feedback, err := e.repo.UserFeedback.CreateOne(&model.UserFeedback{
		DiscordID: req.DiscordID,
		Command:   req.Command,
		Feedback:  req.Feedback,
		MessageID: req.MessageID,
		Status:    "none",
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.HandleUserFeedback] e.repo.UserFeedback.CreateOne failed")
		return err
	}

	// send feedback to channel
	err = e.svc.Discord.SendFeedback(req, feedback.ID.UUID.String())
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.HandleUserFeedback] e.svc.Discord.SendFeedback failed")
		return err
	}
	return nil
}

func (e *Entity) UpdateUserFeedback(req *request.UpdateUserFeedbackRequest) (*model.UserFeedback, error) {
	feedback, err := e.repo.UserFeedback.UpdateStatusByID(req.ID, req.Status)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.UpdateUserFeedback] e.repo.UserFeedback.UpdateStatusByID failed")
		return nil, err
	}
	return feedback, nil
}

func (e *Entity) GetAllUserFeedback(filter string, value string) (feedbacks []model.UserFeedback, err error) {
	switch filter {
	case "command":
		feedbacks, err = e.repo.UserFeedback.GetAllByCommand(value)
		if err != nil {
			e.log.Fields(logger.Fields{"filter": filter, "value": value}).Error(err, "[entity.GetAllUserFeedback] failed to get by command")
			return nil, err
		}
	case "status":
		if value != "none" && value != "confirmed" && value != "completed" {
			err = fmt.Errorf("invalid status")
			return nil, err
		}
		feedbacks, err = e.repo.UserFeedback.GetAllByStatus(value)
		if err != nil {
			e.log.Fields(logger.Fields{"filter": filter, "value": value}).Error(err, "[entity.GetAllUserFeedback] failed to get by status")
			return nil, err
		}
	case "discord_id":
		if value == "" {
			err = fmt.Errorf("discord id empty")
			e.log.Fields(logger.Fields{"filter": filter, "value": value}).Error(err, "[entity.GetAllUserFeedback] failed to get by discord id")
			return nil, err
		}
		feedbacks, err = e.repo.UserFeedback.GetAllByDiscordID(value)
		if err != nil {
			e.log.Fields(logger.Fields{"filter": filter, "value": value}).Error(err, "[entity.GetAllUserFeedback] failed to get by status")
			return nil, err
		}
	default:
		feedbacks, err = e.repo.UserFeedback.GetAll()
		if err != nil {
			e.log.Fields(logger.Fields{"filter": filter, "value": value}).Error(err, "[entity.GetAllUserFeedback] failed to get all")
			return nil, err
		}
	}

	return feedbacks, nil

}

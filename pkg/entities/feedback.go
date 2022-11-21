package entities

import (
	"fmt"
	"strconv"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
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

func (e *Entity) GetAllUserFeedback(filter, value, pg, sz string) (res *response.UserFeedbackResponse, err error) {
	var totalRecord int64
	var feedbacks []model.UserFeedback
	if pg == "" {
		pg = "0"
	}
	if sz == "" {
		sz = "10"
	}
	page, _ := strconv.Atoi(pg)
	size, _ := strconv.Atoi(sz)

	switch filter {
	case "command":
		feedbacks, totalRecord, err = e.repo.UserFeedback.GetAllByCommand(value, page, size)
		if err != nil {
			e.log.Fields(logger.Fields{"filter": filter, "value": value}).Error(err, "[entity.GetAllUserFeedback] failed to get by command")
			return nil, err
		}
	case "status":
		if value != "none" && value != "confirmed" && value != "completed" {
			err := fmt.Errorf("invalid status")
			return nil, err
		}
		feedbacks, totalRecord, err = e.repo.UserFeedback.GetAllByStatus(value, page, size)
		if err != nil {
			e.log.Fields(logger.Fields{"filter": filter, "value": value}).Error(err, "[entity.GetAllUserFeedback] failed to get by status")
			return nil, err
		}
	case "discord_id":
		if value == "" {
			err := fmt.Errorf("discord id empty")
			e.log.Fields(logger.Fields{"filter": filter, "value": value}).Error(err, "[entity.GetAllUserFeedback] failed to get by discord id")
			return nil, err
		}
		feedbacks, totalRecord, err = e.repo.UserFeedback.GetAllByDiscordID(value, page, size)
		if err != nil {
			e.log.Fields(logger.Fields{"filter": filter, "value": value}).Error(err, "[entity.GetAllUserFeedback] failed to get by status")
			return nil, err
		}
	default:
		feedbacks, totalRecord, err = e.repo.UserFeedback.GetAll(page, size)
		if err != nil {
			e.log.Fields(logger.Fields{"filter": filter, "value": value}).Error(err, "[entity.GetAllUserFeedback] failed to get all")
			return nil, err
		}
	}

	return &response.UserFeedbackResponse{Page: page, Size: size, Total: totalRecord, Data: feedbacks}, nil

}

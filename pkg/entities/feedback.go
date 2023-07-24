package entities

import (
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	userfeedback "github.com/defipod/mochi/pkg/repo/user_feedback"
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
		ProfileId: req.ProfileID,
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

func (e *Entity) GetAllUserFeedback(req request.GetUserFeedbackRequest) (res *response.UserFeedbackResponse, err error) {
	feedbacks, total, err := e.repo.UserFeedback.List(userfeedback.FeedbackQuery{
		ProfileID: req.ProfileId,
		DiscordId: req.DiscordId,
		Sort:      "created_at DESC",
		Offset:    req.Page * req.Size,
		Limit:     req.Size,
		Command:   req.Command,
		Status:    req.Status,
	})

	return &response.UserFeedbackResponse{Page: int(req.Page), Size: int(req.Size), Total: total, Data: feedbacks}, nil

}

package response

import "github.com/defipod/mochi/pkg/model"

type UserFeedbackResponse struct {
	Data []model.UserFeedback `json:"data"`
}

type UpdateUserFeedbackResponse struct {
	Data *model.UserFeedback `json:"data"`
}

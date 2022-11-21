package response

import "github.com/defipod/mochi/pkg/model"

type UserFeedbackResponse struct {
	Page  int                  `json:"page"`
	Size  int                  `json:"size"`
	Total int64                `json:"total"`
	Data  []model.UserFeedback `json:"data"`
}

type UpdateUserFeedbackResponse struct {
	Data *model.UserFeedback `json:"data"`
}

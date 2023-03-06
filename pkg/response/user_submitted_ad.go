package response

import "github.com/defipod/mochi/pkg/model"

type GetUserSubmittedAdResponse struct {
	model.UserSubmittedAd
}

type GetAllUserSubmittedAdResponse struct {
	Data []GetAllUserSubmittedAdResponse `json:"data"`
}

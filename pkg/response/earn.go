package response

import (
	"github.com/defipod/mochi/pkg/model"
)

type EarnInfoResponse struct {
	Data *model.EarnInfo `json:"data"`
}

type EarnInfoListResponse struct {
	Data  []model.EarnInfo `json:"data"`
	Page  int              `json:"page"`
	Size  int              `json:"size"`
	Total int64            `json:"total"`
}

type UserEarnResponse struct {
	Data *model.UserEarn `json:"data"`
}

type UserEarnListResponse struct {
	Data  []model.UserEarn `json:"data"`
	Page  int              `json:"page"`
	Size  int              `json:"size"`
	Total int64            `json:"total"`
}

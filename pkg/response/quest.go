package response

import "github.com/defipod/mochi/pkg/model"

type GetUserQuestListResponse struct {
	Data []model.QuestUserList `json:"data"`
}

type ClaimQuestsRewardsResponse struct {
	Data []model.QuestUserReward `json:"data"`
}

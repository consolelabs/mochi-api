package response

import "github.com/defipod/mochi/pkg/model"

type GetUserQuestListResponse struct {
	Data []model.QuestUserList `json:"data"`
}

type ClaimQuestsRewardsResponse struct {
	Data ClaimQuestsRewardsResponseData `json:"data"`
}

type ClaimQuestsRewardsResponseData struct {
	Rewards []model.QuestUserReward `json:"rewards"`
}

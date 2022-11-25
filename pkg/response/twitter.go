package response

import "github.com/defipod/mochi/pkg/model"

type GetTwitterLeaderboardResponseData struct {
	Data       []model.TwitterPostStreak `json:"data"`
	Pagination *PaginationResponse       `json:"metadata"`
}

type GetTwitterLeaderboardResponse struct {
	Data GetTwitterLeaderboardResponseData `json:"data"`
}

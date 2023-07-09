package response

type ZSetWithScoreData struct {
	Score  float64 `json:"score"`
	Member string  `json:"member"`
}

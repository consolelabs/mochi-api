package model

type UpvoteStreakTier struct {
	ID             int `json:"id"`
	StreakRequired int `json:"streak_required"`
	XPPerInterval  int `json:"xp_per_interval"`
	VoteInterval   int ` json:"vote_interval"`
}

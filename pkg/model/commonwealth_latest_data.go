package model

import "time"

type CommonwealthLatestData struct {
	ID          int64     `json:"id"`
	CommunityID string    `json:"community_id"`
	PostCount   int64     `json:"post_count"`
	LatestAt    time.Time `json:"latest_at"`
}

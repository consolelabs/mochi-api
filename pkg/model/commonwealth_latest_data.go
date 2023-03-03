package model

import "time"

type CommonwealthLatestData struct {
	ID          int64     `json:"id"`
	CommunityID string    `json:"community_id"`
	PostCount   int64     `json:"post_count"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Website     string    `json:"website"`
	IconURL     string    `json:"icon_url"`
	LatestAt    time.Time `json:"latest_at"`
}

type CommonwealthDiscussionSubscription struct {
	ID              int64     `json:"id"`
	DiscussionID    int64     `json:"discussion_id"`
	DiscordThreadID string    `json:"discord_thread_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

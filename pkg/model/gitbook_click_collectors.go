package model

import "time"

type GitbookClickCollector struct {
	ID          int       `json:"id"`
	Command     string    `json:"command"`
	Action      string    `json:"action"`
	CountClicks int       `json:"count_clicks"`
	CreatedAt   time.Time `json:"created_at"`
}

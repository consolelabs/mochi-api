package model

import "time"

type EarnInfo struct {
	Id         int64      `json:"id"`
	Title      string     `json:"title"`
	Detail     string     `json:"detail"`
	PrevEarnId *int       `json:"prev_earn_id,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeadlineAt *time.Time `json:"deadline_at,omitempty"`
}

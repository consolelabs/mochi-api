package model

import "time"

type DaoVoteOption struct {
	Id        int64     `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (DaoVoteOption) TableName() string {
	return "dao_vote_option"
}

package request

import (
	"errors"
	"time"
)

type CreateEarnInfoRequest struct {
	Title      string    `json:"title,omitempty"`
	Detail     string    `json:"detail,omitempty"`
	PrevEarnId *int      `json:"prev_earn_id,omitempty"`
	DeadlineAt time.Time `json:"deadline_at,omitempty"`
}

type CreateUserEarnRequest struct {
	UserId     string `json:"-"`
	EarnId     int    `json:"earn_id"`
	Status     string `json:"status"`
	IsFavorite bool   `json:"is_favorite"`
}

type RemoveUserEarnRequest struct {
	UserId string `json:"-"`
	EarnId int    `json:"earn_id"`
}

type GetUserEarnListByUserIdRequest struct {
	UserId string `json:"-"`
	PaginationRequest
}

func (r GetUserEarnListByUserIdRequest) Validate() error {
	if r.UserId == "" {
		return errors.New("invalid user_id")
	}
	return nil
}

type DeleteUserEarnRequest struct {
	UserId string `uri:"id" binding:"required"`
	EarnId int    `uri:"earn_id" binding:"required"`
}

func (r DeleteUserEarnRequest) Validate() error {
	if r.UserId == "" {
		return errors.New("invalid user_id")
	}
	if r.EarnId <= 0 {
		return errors.New("invalid earn_id")
	}
	return nil
}

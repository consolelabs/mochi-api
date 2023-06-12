package request

import (
	"errors"
	"strings"
	"time"
)

const (
	StatusNew     = "new"
	StatusSkipped = "skipped"
	StatusDone    = "done"
	StatusSuccess = "success"
	StatusFailure = "failure"
)

var validStatuses = map[string]bool{
	StatusNew:     true,
	StatusSkipped: true,
	StatusDone:    true,
	StatusSuccess: true,
	StatusFailure: true,
}

type CreateEarnInfoRequest struct {
	Title      string     `json:"title,omitempty"`
	Detail     string     `json:"detail,omitempty"`
	PrevEarnId *int       `json:"prev_earn_id,omitempty"`
	DeadlineAt *time.Time `json:"deadline_at,omitempty"`
}

type CreateUserEarnRequest struct {
	UserId     string `json:"-"`
	EarnId     int    `json:"earn_id"`
	Status     string `json:"status"`
	IsFavorite bool   `json:"is_favorite"`
}

func (r *CreateUserEarnRequest) Validate() error {
	if r.UserId == "" {
		return errors.New("invalid user_id")
	}
	if _, ok := validStatuses[strings.ToLower(r.Status)]; !ok {
		r.Status = StatusNew
	}

	return nil
}

type RemoveUserEarnRequest struct {
	UserId string `json:"-"`
	EarnId int    `json:"earn_id"`
}

type GetUserEarnListByUserIdRequest struct {
	UserId string `json:"-"`
	Status string `form:"status" json:"status"`
	PaginationRequest
}

func (r *GetUserEarnListByUserIdRequest) Validate() error {
	if r.UserId == "" {
		return errors.New("invalid user_id")
	}

	if r.Status != "" {
		if _, ok := validStatuses[strings.ToLower(r.Status)]; !ok {
			return errors.New("invalid status")
		}
	}

	return nil
}

type DeleteUserEarnRequest struct {
	UserId string `uri:"id" binding:"required"`
	EarnId int    `uri:"earn_id" binding:"required"`
}

func (r *DeleteUserEarnRequest) Validate() error {
	if r.UserId == "" {
		return errors.New("invalid user_id")
	}
	if r.EarnId <= 0 {
		return errors.New("invalid earn_id")
	}
	return nil
}

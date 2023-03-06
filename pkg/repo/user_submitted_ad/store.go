package usersubmittedad

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	CreateOne(model.UserSubmittedAd) (*model.UserSubmittedAd, error)
	GetAll() ([]model.UserSubmittedAd, int64, error)
	GetById(id int) (*model.UserSubmittedAd, error)
	UpdateStatus(id int, newStatus string) error
	DeleteOne(id int) error
}

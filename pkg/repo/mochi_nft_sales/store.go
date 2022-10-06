package mochinftsales

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

type Store interface {
	CreateOne(*request.TwitterSalesMessage) error
	GetAllUnnotified() ([]model.TwitterSalesMessage, error)
	GetUnnotified(offset int, limit int) ([]model.TwitterSalesMessage, int64, error)
	DeleteOne(model.TwitterSalesMessage) error
}

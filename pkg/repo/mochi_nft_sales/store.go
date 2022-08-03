package mochinftsales

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

type Store interface {
	CreateOne(*request.TwitterSalesMessage) error
	GetAllUnnotified() ([]model.TwitterSalesMessage, error)
	DeleteOne(*model.TwitterSalesMessage) error
}

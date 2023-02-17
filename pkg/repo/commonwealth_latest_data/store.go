package commonwealth_lastest_data

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	UpsertOne(model.CommonwealthLatestData) error
	GetAll() ([]model.CommonwealthLatestData, error)
}

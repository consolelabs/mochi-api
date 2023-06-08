package earninfo

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	Create(*model.EarnInfo) (*model.EarnInfo, error)
	GetById(int64) (*model.EarnInfo, error)
	List(ListQuery) ([]model.EarnInfo, int64, error)
}

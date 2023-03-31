package treasurer

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(treasurer *model.Treasurer) (*model.Treasurer, error)
}

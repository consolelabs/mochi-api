package envelop

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(envelop *model.Envelop) error
}

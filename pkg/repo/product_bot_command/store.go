package productbotcommand

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	List(ListQuery) ([]model.ProductBotCommand, error)
}

package entities

import (
	"github.com/defipod/mochi/pkg/model"
	productbotcommand "github.com/defipod/mochi/pkg/repo/product_bot_command"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) ProductBotCommand(req request.ProductBotCommandRequest) ([]model.ProductBotCommand, error) {
	return e.repo.ProductBotCommand.List(productbotcommand.ListQuery{
		Code:  req.Code,
		Scope: req.Scope,
	})
}

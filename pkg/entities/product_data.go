package entities

import (
	"github.com/defipod/mochi/pkg/model"
	productbotcommand "github.com/defipod/mochi/pkg/repo/product_bot_command"
	productchangelogs "github.com/defipod/mochi/pkg/repo/product_changelogs"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) ProductBotCommand(req request.ProductBotCommandRequest) ([]model.ProductBotCommand, error) {
	return e.repo.ProductBotCommand.List(productbotcommand.ListQuery{
		Code:  req.Code,
		Scope: req.Scope,
	})
}

func (e *Entity) ProductChangelogs(req request.ProductChangelogsRequest) ([]model.ProductChangelogs, error) {
	return e.repo.ProductChangelogs.List(productchangelogs.ListQuery{
		Product: req.Product,
		Size:    req.Size,
	})
}

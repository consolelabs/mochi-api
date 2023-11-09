package response

import "github.com/defipod/mochi/pkg/model"

type CommandPermissions struct {
	Data []model.CommandPermission `json:"data"`
}

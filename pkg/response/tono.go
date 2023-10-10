package response

import "github.com/defipod/mochi/pkg/model"

type TonoCommandPermissions struct {
	Data []model.TonoCommandPermission `json:"data"`
}

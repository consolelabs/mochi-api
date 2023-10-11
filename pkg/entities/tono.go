package entities

import (
	"github.com/defipod/mochi/pkg/model"
	tonocommandpermission "github.com/defipod/mochi/pkg/repo/tono_command_permission"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) GetTonoCommandPermissions(req request.TonoCommandPermissionsRequest) ([]model.TonoCommandPermission, error) {
	return e.repo.TonoCommandPermission.List(tonocommandpermission.ListQuery{
		Code: req.Code,
	})
}

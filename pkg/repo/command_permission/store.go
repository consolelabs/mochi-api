package commandpermission

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	List(ListQuery) ([]model.CommandPermission, error)
	ListUniqueDiscordPermission() (pers []string, err error)
}

package entities

import (
	"fmt"
	"math/big"

	"github.com/defipod/mochi/pkg/model"
	commandpermission "github.com/defipod/mochi/pkg/repo/command_permission"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) GetCommandPermissions(req request.CommandPermissionsRequest) ([]model.CommandPermission, error) {
	return e.repo.CommandPermission.List(commandpermission.ListQuery{
		Code: req.Code,
	})
}

func (e *Entity) GetInstallBotUrl() (string, error) {
	permissions, err := e.repo.CommandPermission.ListUniqueDiscordPermission()
	if err != nil {
		return "", err
	}

	finalPermission := new(big.Int)
	for _, permission := range permissions {
		flag, ok := new(big.Int).SetString(permission, 10)
		if !ok {
			err := fmt.Errorf("failed to convert string to big int")
			e.log.Error(err, "[Entity] big.Int.SetString() failed")
			return "", err
		}
		finalPermission = finalPermission.Or(finalPermission, flag)
	}

	installBotUrl := fmt.Sprintf("https://discord.com/api/oauth2/authorize?client_id=%s&permissions=%s&scope=bot applications.commands", e.cfg.DiscordApplicationID, finalPermission.String())

	return installBotUrl, nil
}

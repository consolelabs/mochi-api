package vaultinfo

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Get() (vaultInfo *model.VaultInfo, err error)
}

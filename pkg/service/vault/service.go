package vault

import "github.com/defipod/mochi/pkg/config"

type VaultService interface {
	LoadConfig() *config.Config
}

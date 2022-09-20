package serversusagestats

import "github.com/defipod/mochi/pkg/request"

type Store interface {
	CreateOne(info *request.UsageInformation) error
}

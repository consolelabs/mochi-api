package gitbook_click_collectors

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	UpsertOne(info model.GitbookClickCollector) error
	GetByCommandAndAction(cmd, action string) (model.GitbookClickCollector, error)
}

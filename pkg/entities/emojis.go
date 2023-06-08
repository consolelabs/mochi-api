package entities

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/model/errors"
)

func (e *Entity) GetListEmojis(codes []string) ([]*model.Emojis, error) {
	emojis, err := e.repo.Emojis.ListEmojis(codes)
	if err != nil {
		e.log.Error(err, "[entity.GetListEmojis] repo.emojis.GetListEmojis failed")
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrRecordNotFound
		}

		return nil, err
	}

	return emojis, nil
}

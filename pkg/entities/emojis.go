package entities

import (
	"fmt"
	"regexp"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/model/errors"
)

func (e *Entity) GetListEmojis(codes []string) ([]*model.EmojiData, error) {
	// 1. get list emojis from db
	emojis, err := e.repo.Emojis.ListEmojis(codes)
	if err != nil {
		e.log.Error(err, "[entity.GetListEmojis] repo.emojis.GetListEmojis failed")
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrRecordNotFound
		}

		return nil, err
	}

	// 2. convert to emojiData struct
	emojiDatas := make([]*model.EmojiData, 0)
	for _, emoji := range emojis {
		if emoji == nil {
			continue
		}
		if emoji.DiscordId == nil {
			continue
		}

		var emojiUrl string
		// 2.1 get id of emoji -> regex number which has length >= 15
		re := regexp.MustCompile("[0-9]{15,}")
		matchList := re.FindAllString(*emoji.DiscordId, -1)
		if len(matchList) > 0 {
			id := matchList[0]
			emojiUrl = fmt.Sprintf("https://cdn.discordapp.com/emojis/%s.png?size=240&quality=lossless", id)
		}

		emojiDatas = append(emojiDatas, &model.EmojiData{
			Code:     emoji.Code,
			Emoji:    *emoji.DiscordId,
			EmojiUrl: emojiUrl,
		})
	}

	return emojiDatas, nil
}

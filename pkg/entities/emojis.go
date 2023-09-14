package entities

import (
	"fmt"
	"regexp"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/repo/emojis"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) GetListEmojis(req request.GetListEmojiRequest) ([]*model.EmojiData, int64, error) {
	// 1. get list emojis from db
	emojis, total, err := e.repo.Emojis.ListEmojis(emojis.Query{
		Codes: req.ListCode,
		Size:  int(req.Size),
		Page:  int(req.Page),
	})
	if err != nil {
		e.log.Error(err, "[entity.GetListEmojis] repo.emojis.GetListEmojis failed")
		if err == gorm.ErrRecordNotFound {
			return nil, 0, errors.ErrRecordNotFound
		}

		return nil, 0, err
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

	return emojiDatas, total, nil
}

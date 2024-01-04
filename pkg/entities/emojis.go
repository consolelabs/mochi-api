package entities

import (
	"fmt"
	"regexp"
	"strings"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/repo/emojis"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) GetListEmojis(req request.GetListEmojiRequest) ([]*model.EmojiData, int64, error) {
	emojisResponse := make([]*model.ProductMetadataEmojis, 0)
	var total int64
	var err error

	// 1. get list emojis from db
	if !req.IsQueryAll {
		emojisResponse, total, err = e.GetListEmojisByCode(req)
		if err != nil {
			e.log.Error(err, "[entity.GetListEmojis] e.GetListEmojisByCode failed")
			return nil, 0, err
		}
	} else {
		emojisResponse, total, err = e.repo.Emojis.ListEmojis(emojis.Query{
			// Codes: req.ListCode,
			Size: int(req.Size),
			Page: int(req.Page),
		})
		if err != nil {
			e.log.Error(err, "[entity.GetListEmojis] repo.emojis.GetListEmojis failed")
			if err == gorm.ErrRecordNotFound {
				return nil, 0, errors.ErrRecordNotFound
			}

			return nil, 0, err
		}
	}

	// 2. convert to emojiData struct
	emojiDatas := make([]*model.EmojiData, 0)
	for _, emoji := range emojisResponse {
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

func (e *Entity) GetListEmojisByCode(req request.GetListEmojiRequest) ([]*model.ProductMetadataEmojis, int64, error) {
	// TODO: get list native token from db mochi-pay
	nativeTokens := []string{"eth", "ftm", "bnb", "sol", "matic"}
	listCodeQuery := make([]string, 0)
	for _, code := range req.ListCode {
		alreadyAdded := false
		for idx, token := range nativeTokens {
			if strings.Contains(strings.ToLower(code), token) {
				listCodeQuery = append(listCodeQuery, strings.ToUpper(token))
				alreadyAdded = true
			} else {
				if idx == len(nativeTokens)-1 && !alreadyAdded {
					listCodeQuery = append(listCodeQuery, code)
				}
			}
		}
	}

	emojis, total, err := e.repo.Emojis.ListEmojis(emojis.Query{
		Codes: listCodeQuery,
	})
	if err != nil {
		e.log.Error(err, "[entity.GetListEmojis] repo.emojis.GetListEmojis failed")
		if err == gorm.ErrRecordNotFound {
			return nil, 0, errors.ErrRecordNotFound
		}

		return nil, 0, err
	}

	emojiResponse := make([]*model.ProductMetadataEmojis, 0)
	for _, code := range req.ListCode {
		alreadyAdded := false
		for idx, emoji := range emojis {
			if strings.Contains(strings.ToLower(code), strings.ToLower(emoji.Code)) {
				emojiResponse = append(emojiResponse, &model.ProductMetadataEmojis{
					ID:         emoji.ID,
					Code:       code,
					DiscordId:  emoji.DiscordId,
					TelegramId: emoji.TelegramId,
					TwitterId:  emoji.TwitterId,
				})
				alreadyAdded = true
			} else {
				if idx == len(emojis)-1 && !alreadyAdded {
					emojiResponse = append(emojiResponse, emoji)
				}
			}
		}
	}

	return emojiResponse, total, nil
}

func (e *Entity) GetEmojiByCode(code string) (*model.EmojiData, error) {
	emoji, err := e.repo.Emojis.GetByCode(code)
	if err != nil {
		e.log.Error(err, "[entity.GetEmojiByCode] repo.emojis.GetByCode failed")
		return nil, err
	}

	var emojiUrl string
	re := regexp.MustCompile("[0-9]{15,}")
	matchList := re.FindAllString(*emoji.DiscordId, -1)
	if len(matchList) > 0 {
		id := matchList[0]
		emojiUrl = fmt.Sprintf("https://cdn.discordapp.com/emojis/%s.png?size=240&quality=lossless", id)
	}

	return &model.EmojiData{
		Code:     emoji.Code,
		Emoji:    *emoji.DiscordId,
		EmojiUrl: emojiUrl,
	}, nil
}

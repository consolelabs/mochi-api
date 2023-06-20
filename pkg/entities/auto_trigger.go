package entities

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

type KafkaNotification struct {
	Type            string          `json:"type"`
	TriggerMetadata TriggerMetadata `json:"trigger_metadata"`
}

type TriggerMetadata struct {
	Content       string                 `json:"content"`
	UserProfileID string                 `json:"user_profile_id"`
	EmbedMessage  discordgo.MessageEmbed `json:"embed_message"`
}

func (e *Entity) HandleTrigger(message request.AutoTriggerRequest) error {
	// get all trigger join with condition and condition value
	autoTriggers, err := e.repo.AutoTrigger.GetAutoTriggers(message.GuildId)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"message": message}).Error(err, "[handler.HandleTrigger] - failed to query auto trigger")
		return err
	}
	// Loop all trigger
	// - Loop all condition
	// -- Loop all condition value
	// --- Loop all child condition value
	// - Loop all child condition

	// get message type, need describe use-case here
	messageType := "invalid"
	if message.Content != "" {
		messageType = "createMessage"
	} else if message.Reaction != "" {
		messageType = "reactionAdd"
	} else {
		return nil
	}

	for _, autoTrigger := range autoTriggers {
		triggerMatch := false
		// Loop all condition
		for _, condition := range autoTrigger.Conditions {
			// check message condition type
			if condition.Type.Type != messageType {
				// trigger not match type, skip to next trigger
				break
			}

			err, ok := e.AutoCheckConditions(condition, message)
			if err != nil || !ok {
				break
			}
			triggerMatch = true
		}

		if triggerMatch {
			// trigger match, execute action
			e.DoAction(autoTrigger.Actions, message)
		}
	}

	return nil
}

// parse the type and check all condition value of a condition
func (e *Entity) AutoCheckConditions(condition model.AutoCondition, message request.AutoTriggerRequest) (error, bool) {
	// check required channel id condition
	if len(condition.ChannelId) > 0 {
		channels := strings.Split(condition.ChannelId, ",")
		for _, channel := range channels {
			if channel != message.ChannelId {
				return nil, false
			}
		}
	}

	// check allowed user id condition
	if len(condition.UserIds) > 0 {
		userIds := strings.Split(condition.UserIds, ",")
		for _, userId := range userIds {
			if userId != message.UserID {
				return nil, false
			}
		}
	}

	// Loop all condition value of this condition
	err, ok := e.AutoCheckConditionValues(condition.ConditionValues, 0, message)
	if err != nil || !ok {
		return err, false
	}

	// loop all child condition
	for _, childCondition := range condition.ChildConditions {
		err, ok := e.AutoCheckConditions(childCondition, message)
		if err != nil || !ok {
			return err, false
		}
	}
	return nil, true
}

// parse the type and check call condition value of a condition
func (e *Entity) AutoCheckConditionValues(conditionValue []model.AutoConditionValue, index int, message request.AutoTriggerRequest) (error, bool) {
	valid := false
	err := error(nil)
	switch conditionValue[index].Type.Type {
	case "createMessage":
		err, valid = e.OperatorForMessage(conditionValue[index].Operator, fmt.Sprintf("%v", message.Content), conditionValue[index].Matches)
	case "reactionAdd":
		if message.ReactionCount > 0 {
			return nil, true
		}
	// user react 10 Y in channel X to post K of User Z
	case "totalReact":
		err, valid = e.OperatorNumber(conditionValue[index].Operator, fmt.Sprintf("%v", message.ReactionCount), conditionValue[index].Matches)
		// react A in channel X
	case "reactType":
		err, valid = e.OperatorString(conditionValue[index].Operator, message.Reaction, conditionValue[index].Matches)
	case "userRole":
		valid = e.OperatorRoles(conditionValue[index].Operator, message.UserRoles, conditionValue[index].Matches)
	case "authorRole":
		valid = e.OperatorRoles(conditionValue[index].Operator, message.AuthorRoles, conditionValue[index].Matches)
	default:
		valid = false
	}

	if err != nil || !valid {
		return err, false
	}

	// check next and condition value
	if index < len(conditionValue)-1 {
		err, valid = e.AutoCheckConditionValues(conditionValue, index+1, message)
	}

	return nil, valid
}

// validate all condition in type string
func (e *Entity) OperatorString(operator string, a string, b string) (error, bool) {
	result := false
	switch operator {
	case "<":
		result = reflect.ValueOf(a).String() < reflect.ValueOf(b).String()
	case ">":
		result = reflect.ValueOf(a).String() > reflect.ValueOf(b).String()
	case ">=":
		result = reflect.ValueOf(a).String() >= reflect.ValueOf(b).String()
	case "==":
		result = reflect.ValueOf(a).String() == reflect.ValueOf(b).String()
	case "!=":
		result = reflect.ValueOf(a).String() != reflect.ValueOf(b).String()
	case "in": // TODO
		// parse b to array and match
		result = reflect.ValueOf(a).String() == reflect.ValueOf(b).String()
	case "not in": // TODO
		// parse b to array and match
		result = reflect.ValueOf(a).String() != reflect.ValueOf(b).String()

	default:
		e.log.Debug("Invalid operator")
	}
	return nil, result
}

// validate all condition in type message content
func (e *Entity) OperatorForMessage(operator string, a string, b string) (error, bool) {
	result := false
	if len(a) == 0 {
		return nil, false
	}

	switch operator {
	case "":
		result = len(a) > 0
	case "==":
		result = a == b
	case "!=":
		result = a != b
	case "in":
		// parse b to array and match
		bItems := strings.Split(b, ",")
		for _, bItem := range bItems {
			if reflect.ValueOf(a).String() == reflect.ValueOf(bItem).String() {
				result = true
				break
			}
		}
	case "not in":
		// parse json array to array and match
		bItems := strings.Split(b, ",")
		for _, bItem := range bItems {
			if reflect.ValueOf(a).String() != reflect.ValueOf(bItem).String() {
				result = true
				break
			}
		}
	default:
		e.log.Debug("Invalid operator")
	}
	return nil, result
}

func (e *Entity) OperatorNumber(operator string, a string, b string) (error, bool) {
	result := false
	numA, err := strconv.Atoi(a)
	if err != nil {
		return err, false
	}

	numB, err := strconv.Atoi(b)
	if err != nil {
		return err, false
	}

	switch operator {
	case "==":
		result = numA == numB
	case "!=":
		result = numA != numB
	case ">=":
		result = numA >= numB
	case "<=":
		result = numA >= numB
	case ">":
		result = numA > numB
	case "<":
		result = numA < numB
	default:
		e.log.Debug("Invalid operator")
	}
	return nil, result
}

func (e *Entity) OperatorRoles(operator string, userRoles []string, requiredRoles string) bool {
	result := false
	if len(userRoles) == 0 {
		return false
	}

	// split requiredRoles to array
	redRoles := strings.Split(requiredRoles, ",")

	// check user role
	switch operator {
	case "in":
		for _, userRole := range userRoles {
			for _, requiredRole := range redRoles {
				if strings.EqualFold(userRole, requiredRole) {
					result = true
					break
				}
			}
		}
	case "not in":
		result = true
		for _, userRole := range userRoles {
			for _, requiredRole := range userRoles {
				if strings.EqualFold(userRole, requiredRole) {
					result = false
					break
				}
			}
		}
	default:
		e.log.Debug("Invalid operator")
	}
	return result
}

func (e *Entity) DoAction(action []model.AutoAction, message request.AutoTriggerRequest) error {
	for _, act := range action {
		actionCount, err := e.repo.AutoActionHistory.CountByTriggerActionUserMessage(act.TriggerId, act.Id, message.AuthorId, message.MessageId)
		if err != nil {
			e.log.Fields(logger.Fields{"TriggerId": act.TriggerId, "ActionId": act.Id, "UserId": message.AuthorId}).Error(err, "Do action error: action was existed")
			return err
		}
		if actionCount >= int64(act.LimitPerUser) {
			continue
		}

		switch act.Type.Type {
		case "sendMessage":
			err = e.actionSendMessage(act.Content, &act.Embed, message.UserID)
		case "sendDM":
			fmt.Println("sendDM " + act.Content)
		case "addRole":
			fmt.Println("addRole " + act.Content)
		case "removeRole":
			fmt.Println("removeRole " + act.Content)
		case "kick":
			fmt.Println("kick " + act.Content)
		case "ban":
			fmt.Println("ban " + act.Content)
		case "vaultTransfer":
			err = e.actionVaultTransfer(act.ActionData, message)
		default:
			e.log.Debug("Invalid action")
		}
		if err != nil {
			e.log.Fields(logger.Fields{"TriggerId": act.TriggerId, "ActionId": act.Id, "UserId": message.AuthorId}).Error(err, "Do action error")
			return err
		}

		// save action history, TODO: should save action result for debug
		err = e.repo.AutoActionHistory.Create(&model.AutoActionHistory{
			TriggerId:     act.TriggerId,
			ActionId:      act.Id,
			UserDiscordId: message.AuthorId,
			MessageId:     message.MessageId,
			Total:         int(actionCount) + 1,
		})
		if err != nil {
			e.log.Fields(logger.Fields{"TriggerId": act.TriggerId, "ActionId": act.Id, "UserId": message.AuthorId}).Error(err, "Do action error")
			return err
		}

		// if then action is settled then do then action
		if act.ThenAction != nil {
			err = e.DoAction([]model.AutoAction{*act.ThenAction}, message)
			if err != nil {
				e.log.Fields(logger.Fields{"TriggerId": act.TriggerId, "ActionId": act.Id, "UserId": message.AuthorId}).Error(err, "Do action error")
				return err
			}
		}
	}
	return nil
}

func (e *Entity) actionSendMessage(content string, embed *model.AutoEmbed, discordId string) error {
	var message = KafkaNotification{
		Type: "trigger",
		TriggerMetadata: TriggerMetadata{
			UserProfileID: discordId,
			Content:       content,
		},
	}

	if embed != nil {
		var discordEmbed discordgo.MessageEmbed
		discordEmbed.Title = embed.Title
		discordEmbed.Description = embed.Description
		discordEmbed.URL = embed.Url

		// colorInt, _ := strconv.ParseInt(embed.Color[1:], 16, 32)
		// discordEmbed.Color = int(colorInt)

		if len(embed.Thumbnail) > 0 {
			discordEmbed.Thumbnail = &discordgo.MessageEmbedThumbnail{
				URL: embed.Thumbnail,
			}
		}
		if embed.Image != nil {
			discordEmbed.Image = &discordgo.MessageEmbedImage{
				URL:    embed.Image.Url,
				Width:  embed.Image.Width,
				Height: embed.Image.Height,
			}
		}

		if embed.Video != nil {
			discordEmbed.Video = &discordgo.MessageEmbedVideo{
				URL:    embed.Video.Url,
				Width:  embed.Video.Width,
				Height: embed.Video.Height,
			}
		}
		if embed.Footer != nil {
			discordEmbed.Footer = &discordgo.MessageEmbedFooter{
				Text:    embed.Footer.Content,
				IconURL: embed.Footer.ImageUrl,
			}
		}
		message.TriggerMetadata.EmbedMessage = discordEmbed
	}
	bytes, _ := json.Marshal(message)
	key := strconv.Itoa(rand.Intn(100000))
	err := e.kafka.Produce("mochiNotification.local", key, bytes) // TODO move to env
	if err != nil {
		e.log.Error(err, "[actionSendMessage] Produce error")
	}
	return err
}

func (e *Entity) actionVaultTransfer(actionData string, message request.AutoTriggerRequest) error {
	if actionData == "" {
		return nil
	}
	var req model.AutoTransferVaultTokenRequest
	json.Unmarshal([]byte(actionData), &req)
	// transfer vault to author
	req.Target = message.AuthorId

	// validate guild
	if req.GuildId != message.GuildId {
		return errors.New("[e.actionVaultTransfer] guild id is required")
	}

	err := e.AutoTransferVaultToken(&req)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[e.actionVaultTransfer] - failed to transfer vault token")
	}
	return err
}

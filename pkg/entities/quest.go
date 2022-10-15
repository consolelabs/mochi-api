package entities

import (
	"math/rand"
	"time"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/repo/quest"
	questuserlist "github.com/defipod/mochi/pkg/repo/quest_user_list"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/util"
	"github.com/google/uuid"
)

func (e *Entity) GetUserQuestList(req request.GetUserQuestListRequest) ([]model.QuestUserList, error) {
	now := time.Now().UTC()
	startTime := util.StartOfDay(now)
	listQ := questuserlist.ListQuery{
		UserID:    &req.UserID,
		StartTime: &startTime,
		Routine:   &req.Routine,
	}
	list, err := e.repo.QuestUserList.List(listQ)
	if err != nil {
		e.log.Fields(logger.Fields{"listQ": listQ}).Error(err, "[entity.GetUserQuestList] repo.QuestUserList.List() failed")
		return nil, err
	}
	if len(list) == 0 {
		return e.generateUserQuestList(request.GenerateUserQuestListRequest{
			UserID:    req.UserID,
			Routine:   req.Routine,
			StartTime: startTime,
			Quantity:  req.Quantity,
		})
	}
	return list, nil
}

func (e *Entity) getBonusQuest(routine model.QuestRoutine) (*model.Quest, error) {
	bonus := model.BONUS
	listQ := quest.ListQuery{Routine: &routine, Action: &bonus}
	quests, err := e.repo.Quest.List(listQ)
	if err != nil {
		e.log.Fields(logger.Fields{"listQ": listQ}).Error(err, "[entity.getBonusQuest] repo.Quest.List() failed")
		return nil, err
	}
	if len(quests) == 0 {
		e.log.Info("[entity.getBonusQuest] Bonus quest not found")
		return nil, nil
	}
	return &quests[0], nil
}

func (e *Entity) generateUserQuestList(req request.GenerateUserQuestListRequest) ([]model.QuestUserList, error) {
	bonus := model.BONUS
	quests, err := e.repo.Quest.List(quest.ListQuery{Routine: &req.Routine, NotAction: &bonus})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.generateUserQuestList] repo.Quest.List() failed")
		return nil, err
	}
	size := util.MinInt(req.Quantity, len(quests))
	userQuests := make([]model.QuestUserList, 0, size)
	for i := 0; i < size; i++ {
		rand.Seed(time.Now().UnixNano())
		idx := rand.Intn(size - i)
		item := quests[idx]
		q := model.QuestUserList{
			UserID:    req.UserID,
			QuestID:   item.ID,
			Action:    item.Action,
			Routine:   req.Routine,
			Target:    item.Frequency,
			StartTime: req.StartTime,
			Quest:     &item,
		}
		// currently only daily quest supported
		switch req.Routine {
		case model.DAILY:
			q.EndTime = q.StartTime.Add(24 * time.Hour)
		default:
			q.EndTime = q.StartTime.Add(24 * time.Hour)
		}
		userQuests = append(userQuests, q)
		quests = append(quests[:idx], quests[idx+1:]...)
		util.Shuffle(quests)
	}
	// append bonus quest
	bonusQuest, err := e.getBonusQuest(req.Routine)
	if err != nil {
		e.log.Fields(logger.Fields{"routine": req.Routine}).Error(err, "[entity.generateUserQuestList] entity.getBonusQuest() failed")
		return nil, err
	}
	if bonusQuest != nil {
		userQuests = append(userQuests, model.QuestUserList{
			UserID:    req.UserID,
			QuestID:   bonusQuest.ID,
			Action:    bonusQuest.Action,
			Routine:   bonusQuest.Routine,
			Target:    bonusQuest.Frequency,
			StartTime: req.StartTime,
			Quest:     bonusQuest,
		})
	}
	// save user quests list
	err = e.repo.QuestUserList.UpsertMany(userQuests)
	if err != nil {
		e.log.Fields(logger.Fields{"list": userQuests}).Error(err, "[entity.generateUserQuestList] repo.QuestUserList.UpsertMany() failed")
		return nil, err
	}
	return userQuests, nil
}

func (e *Entity) UpdateUserQuestProgress(log *model.QuestUserLog) error {
	// bonus quest progress cannot be updated from outside
	if log.Action == model.BONUS || log.UserID == "" {
		return nil
	}
	startTime := util.StartOfDay(time.Now().UTC())
	// check if user has the quest
	uQuestQ := questuserlist.ListQuery{Action: &log.Action, StartTime: &startTime, UserID: &log.UserID}
	uQuests, err := e.repo.QuestUserList.List(uQuestQ)
	if err != nil {
		e.log.Fields(logger.Fields{"uQuestQ": uQuestQ}).Error(err, "[entity.UpdateUserQuestProgress] repo.QuestUserList.List() failed")
		return err
	}
	if len(uQuests) == 0 {
		e.log.Fields(logger.Fields{"uQuestQ": uQuestQ}).Info("[entity.UpdateUserQuestProgress] repo.QuestUserList.List() returns empty")
		return nil
	}
	for _, uQuest := range uQuests {
		log.QuestID = uQuest.QuestID
		log.Target = uQuest.Quest.Frequency
		// create quest log
		if err := e.repo.QuestUserLog.CreateOne(log); err != nil {
			e.log.Fields(logger.Fields{"log": log}).Error(err, "[entity.UpdateUserQuestProgress] repo.QuestUserLog.CreateOne() failed")
			return err
		}
		// update quest progress
		uQuest.Current++
		if uQuest.Current >= uQuest.Target {
			uQuest.IsCompleted = true
		}
		err = e.repo.QuestUserList.UpsertMany([]model.QuestUserList{uQuest})
		if err != nil {
			e.log.Fields(logger.Fields{"uQuest": uQuest}).Error(err, "[entity.UpdateUserQuestProgress] repo.QuestUserList.UpsertMany() failed")
			return err
		}
		// if the quest have not been completed yet, no need to update bonus quest ...
		if !uQuest.IsCompleted {
			return nil
		}
		// ... else update bonus quest progress
		err = e.updateUserBonusQuest(uQuest.Routine, log.UserID, startTime)
		if err != nil {
			e.log.Fields(logger.Fields{
				"routine":   uQuest.Routine,
				"userID":    log.UserID,
				"startTime": startTime,
			}).Error(err, "[entity.UpdateUserQuestProgress] entity.updateUserBonusQuest() failed")
			return err
		}
	}
	return nil
}

func (e *Entity) updateUserBonusQuest(routine model.QuestRoutine, userID string, startTime time.Time) error {
	bonusQuest, err := e.getBonusQuest(routine)
	if err != nil {
		e.log.Fields(logger.Fields{"routine": routine}).Error(err, "[entity.updateUserBonusQuest] entity.getBonusQuest() failed")
		return err
	}
	if bonusQuest == nil {
		return nil
	}
	listQ := questuserlist.ListQuery{
		UserID:    &userID,
		QuestID:   &bonusQuest.ID,
		StartTime: &startTime,
	}
	userBonusQuest, err := e.repo.QuestUserList.List(listQ)
	if err != nil {
		e.log.Fields(logger.Fields{"listQ": listQ}).Error(err, "[entity.updateUserBonusQuest] repo.QuestUserList.List() failed")
		return err
	}
	if len(userBonusQuest) == 0 {
		return nil
	}
	completed := true
	// count completed quests (excluding bonus quest)
	completedQuestsQ := questuserlist.ListQuery{
		UserID:      &userID,
		StartTime:   &startTime,
		IsCompleted: &completed,
		NotAction:   &bonusQuest.Action,
	}
	completedQuests, err := e.repo.QuestUserList.List(completedQuestsQ)
	if err != nil {
		e.log.Fields(logger.Fields{"completedQuestsQ": completedQuestsQ}).Error(err, "[entity.updateUserBonusQuest] repo.QuestUserList.List() failed")
		return err
	}
	userBonusQuest[0].Current = len(completedQuests)
	if userBonusQuest[0].Current >= userBonusQuest[0].Target {
		userBonusQuest[0].IsCompleted = true
	}
	err = e.repo.QuestUserList.UpsertMany(userBonusQuest)
	if err != nil {
		e.log.Fields(logger.Fields{"userBonusQuest": userBonusQuest}).Error(err, "[entity.updateUserBonusQuest] repo.QuestUserList.UpsertMany() failed")
		return err
	}
	return nil
}

func (e *Entity) ClaimQuestsRewards(req request.ClaimQuestsRewardsRequest) ([]model.QuestUserReward, error) {
	now := time.Now().UTC()
	startTime := util.StartOfDay(now)
	completed := true
	claimed := false
	listQ := questuserlist.ListQuery{
		UserID:      &req.UserID,
		StartTime:   &startTime,
		Routine:     &req.Routine,
		IsCompleted: &completed,
		IsClaimed:   &claimed,
	}
	list, err := e.repo.QuestUserList.List(listQ)
	if err != nil {
		e.log.Fields(logger.Fields{"listQ": listQ}).Error(err, "[entity.ClaimQuestsRewards] repo.QuestUserList.List() failed")
		return nil, err
	}
	if len(list) == 0 {
		e.log.Fields(logger.Fields{"listQ": listQ}).Error(err, "[entity.ClaimQuestsRewards] repo.QuestUserList.List() returns empty")
		return nil, nil
	}
	questIDs := make([]uuid.UUID, len(list))
	for i := range list {
		questIDs = append(questIDs, list[i].QuestID)
		list[i].IsClaimed = true
	}
	rewards, err := e.repo.QuestReward.GetQuestRewards(questIDs)
	if err != nil {
		e.log.Fields(logger.Fields{"questIDs": questIDs}).Error(err, "[entity.ClaimQuestsRewards] repo.QuestReward.GetQuestRewards() failed")
		return nil, err
	}
	uRewards := make([]model.QuestUserReward, 0, len(rewards))
	for i, r := range rewards {
		uRewards = append(uRewards, model.QuestUserReward{
			UserID:       req.UserID,
			QuestID:      r.QuestID,
			RewardID:     r.ID,
			RewardTypeID: r.RewardTypeID,
			RewardAmount: r.RewardAmount,
			PassID:       r.PassID,
			StartTime:    startTime,
			Reward:       &rewards[i],
		})
	}
	err = e.repo.QuestUserReward.CreateMany(uRewards)
	if err != nil {
		e.log.Fields(logger.Fields{"uRewards": uRewards}).Error(err, "[entity.ClaimQuestsRewards] repo.QuestUserReward.CreateMany() failed")
		return nil, err
	}
	// update claim status
	err = e.repo.QuestUserList.UpsertMany(list)
	if err != nil {
		e.log.Fields(logger.Fields{"list": list}).Error(err, "[entity.ClaimQuestsRewards] repo.QuestUserList.UpsertMany() failed")
		return nil, err
	}
	return uRewards, nil
}

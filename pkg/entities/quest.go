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
			Quantity:  5,
		})
	}
	return list, nil
}

func (e *Entity) generateUserQuestList(req request.GenerateUserQuestListRequest) ([]model.QuestUserList, error) {
	quests, err := e.repo.Quest.List(quest.ListQuery{Routine: &req.Routine})
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
	// save user quests list
	err = e.repo.QuestUserList.UpsertMany(userQuests)
	if err != nil {
		e.log.Fields(logger.Fields{"list": userQuests}).Error(err, "[entity.generateUserQuestList] repo.QuestUserList.UpsertMany() failed")
		return nil, err
	}
	return userQuests, nil
}

func (e *Entity) UpdateUserQuestProgress(log *model.QuestUserLog) error {
	// get quest by action (e.g. GM)
	questQ := quest.ListQuery{Action: &log.Action}
	quests, err := e.repo.Quest.List(questQ)
	if err != nil {
		e.log.Fields(logger.Fields{"questQ": questQ}).Error(err, "[entity.UpdateUserQuestProgress] repo.Quest.List() failed")
		return err
	}
	if len(quests) == 0 {
		e.log.Fields(logger.Fields{"questQ": questQ}).Info("[entity.UpdateUserQuestProgress] repo.Quest.List() returns empty")
		return nil
	}
	log.QuestID = quests[0].ID
	log.Target = quests[0].Frequency
	// create quest log
	if err := e.repo.QuestUserLog.CreateOne(log); err != nil {
		e.log.Fields(logger.Fields{"log": log}).Error(err, "[entity.UpdateUserQuestProgress] repo.QuestUserLog.CreateOne() failed")
		return err
	}
	// get user's quests list ...
	startTime := util.StartOfDay(time.Now().UTC())
	listQ := questuserlist.ListQuery{
		UserID:    &log.UserID,
		QuestID:   &log.QuestID,
		StartTime: &startTime,
	}
	userQuest, err := e.repo.QuestUserList.List(listQ)
	if err != nil {
		e.log.Fields(logger.Fields{"listQ": listQ}).Error(err, "[entity.UpdateUserQuestProgress] repo.QuestUserList.List() failed")
		return err
	}
	if len(userQuest) == 0 {
		return nil
	}
	// ... and update progress
	userQuest[0].Current++
	if userQuest[0].Current >= userQuest[0].Target {
		userQuest[0].IsCompleted = true
	}
	err = e.repo.QuestUserList.UpsertMany(userQuest)
	if err != nil {
		e.log.Fields(logger.Fields{"userQuest": userQuest}).Error(err, "[entity.UpdateUserQuestProgress] repo.QuestUserList.UpsertMany() failed")
	}
	return err
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
	for _, r := range rewards {
		uRewards = append(uRewards, model.QuestUserReward{
			UserID:       req.UserID,
			QuestID:      r.QuestID,
			RewardID:     r.ID,
			RewardTypeID: r.RewardTypeID,
			RewardAmount: r.RewardAmount,
			PassID:       r.PassID,
			StartTime:    startTime,
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

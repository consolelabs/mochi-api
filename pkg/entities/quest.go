package entities

import (
	"math"
	"math/rand"
	"sort"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/repo/quest"
	questreward "github.com/defipod/mochi/pkg/repo/quest_reward"
	queststreak "github.com/defipod/mochi/pkg/repo/quest_streak"
	questuserlist "github.com/defipod/mochi/pkg/repo/quest_user_list"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
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
		generateReq := request.GenerateUserQuestListRequest{
			UserID:    req.UserID,
			Routine:   req.Routine,
			StartTime: startTime,
		}
		list, err = e.generateUserQuestList(generateReq)
		if err != nil {
			e.log.Fields(logger.Fields{"generateReq": generateReq}).Error(err, "[entity.GetUserQuestList] entity.generateUserQuestList() failed")
			return nil, err
		}
	} else {
		list, err = e.finalizeUserQuestList(req.UserID, string(model.VOTE), list, &list[0].Multiplier)
		if err != nil {
			e.log.Fields(logger.Fields{"userID": req.UserID, "list": list}).Error(err, "[entity.GetUserQuestList] entity.finalizeUserQuestList() failed")
			return nil, err
		}
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Quest.Title < list[j].Quest.Title
	})
	return list, nil
}

func (e *Entity) getBonusQuest(routine model.QuestRoutine) (*model.Quest, error) {
	listQ := quest.ListQuery{Routine: string(routine), Action: string(model.BONUS)}
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
	if req.Quantity == 0 {
		req.Quantity = 5
	}
	e.log.Fields(logger.Fields{"req": req}).Info("[entity.generateUserQuestList] start generating quests list ...")
	quests, err := e.repo.Quest.List(quest.ListQuery{Routine: string(req.Routine), NotActions: []model.QuestAction{model.BONUS}})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.generateUserQuestList] repo.Quest.List() failed")
		return nil, err
	}
	e.log.Fields(logger.Fields{"req": req}).Infof("[entity.generateUserQuestList] %d %s quests found", len(quests), req.Routine)
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
			Routine:   item.Routine,
			Target:    item.Frequency,
			StartTime: req.StartTime,
			Quest:     &item,
		}
		// currently only daily quest supported
		switch item.Routine {
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
	userQuests, err = e.finalizeUserQuestList(req.UserID, string(model.VOTE), userQuests, nil)
	if err != nil {
		e.log.Fields(logger.Fields{"userID": req.UserID, "userQuests": userQuests}).Error(err, "[entity.generateUserQuestList] entity.finalizeUserQuestList() failed")
		return nil, err
	}
	e.log.Fields(logger.Fields{"userQuests": userQuests}).Infof("[entity.generateUserQuestList] creating %d quests for user %s", len(userQuests), req.UserID)
	// save user quests list
	err = e.repo.QuestUserList.UpsertMany(userQuests)
	if err != nil {
		e.log.Fields(logger.Fields{"list": userQuests}).Error(err, "[entity.generateUserQuestList] repo.QuestUserList.UpsertMany() failed")
		return nil, err
	}
	e.log.Fields(logger.Fields{"userQuests": userQuests}).Infof("[entity.generateUserQuestList] created %d quests for user %s", len(userQuests), req.UserID)
	return userQuests, nil
}

func (e *Entity) getOrGenerateUserQuests(startTime time.Time, userID string, routine model.QuestRoutine) ([]model.QuestUserList, error) {
	uQuestQ := questuserlist.ListQuery{StartTime: &startTime, UserID: &userID, Routine: &routine}
	uQuests, err := e.repo.QuestUserList.List(uQuestQ)
	if err != nil {
		e.log.Fields(logger.Fields{"uQuestQ": uQuestQ}).Error(err, "[entity.getOrGenerateUserQuests] repo.QuestUserList.List() failed")
		return nil, err
	}
	if len(uQuests) > 0 {
		e.log.Fields(logger.Fields{"uQuestQ": uQuestQ}).Errorf(err, "[entity.getOrGenerateUserQuests] found %d quests for user %s", len(uQuests), userID)
		return uQuests, nil
	}
	generateReq := request.GenerateUserQuestListRequest{
		UserID:    userID,
		Routine:   routine,
		StartTime: startTime,
	}
	uQuests, err = e.generateUserQuestList(generateReq)
	if err != nil {
		e.log.Fields(logger.Fields{"generateReq": generateReq}).Error(err, "[entity.getOrGenerateUserQuests] entity.generateUserQuestList() failed")
		return nil, err
	}
	if len(uQuests) == 0 {
		e.log.Fields(logger.Fields{"req": generateReq}).Info("[entity.getOrGenerateUserQuests] no quests were generated ")
	}
	sort.Slice(uQuests, func(i, j int) bool {
		return uQuests[i].Quest.Title < uQuests[j].Quest.Title
	})
	return uQuests, nil
}

func (e *Entity) UpdateUserQuestProgress(log *model.QuestUserLog) error {
	// bonus quest progress cannot be updated from outside
	if log.Action == model.BONUS || log.UserID == "" {
		return nil
	}
	startTime := util.StartOfDay(time.Now().UTC())
	routines, err := e.repo.Quest.GetAvailableRoutines()
	if err != nil {
		e.log.Error(err, "[entity.UpdateUserQuestProgress] repo.Quest.GetAvailableRoutines() failed")
		return err
	}
	if len(routines) == 0 {
		e.log.Error(err, "[entity.UpdateUserQuestProgress] no available routines")
		return nil
	}
	for _, routine := range routines {
		// check if user quests have been generated yet ...
		_, err := e.getOrGenerateUserQuests(startTime, log.UserID, routine)
		if err != nil {
			e.log.Fields(logger.Fields{
				"startTime": startTime,
				"userID":    log.UserID,
				"routine":   routine,
			}).Error(err, "[entity.UpdateUserQuestProgress] entity.getOrGenerateUserQuests() failed")
			return err
		}
	}
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
		if uQuest.IsCompleted {
			continue
		}
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
		NotActions:  []model.QuestAction{bonusQuest.Action},
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

func (e *Entity) ClaimQuestsRewards(req request.ClaimQuestsRewardsRequest) (*response.ClaimQuestsRewardsResponseData, error) {
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
		QuestID:     req.QuestID,
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
	rewards, err := e.repo.QuestReward.List(questreward.ListQuery{QuestIDs: questIDs})
	if err != nil {
		e.log.Fields(logger.Fields{"questIDs": questIDs}).Error(err, "[entity.ClaimQuestsRewards] repo.QuestReward.GetQuestRewards() failed")
		return nil, err
	}
	var uRewards []model.QuestUserReward
	for i, r := range rewards {
		uRewards = append(uRewards, model.QuestUserReward{
			UserID:       req.UserID,
			QuestID:      r.QuestID,
			RewardID:     r.ID,
			RewardTypeID: r.RewardTypeID,
			RewardAmount: r.RewardAmount,
			PassID:       r.PassID,
			StartTime:    &startTime,
			Reward:       &rewards[i],
		})
	}
	uRewards, err = e.finalizeUserRewards(req.UserID, string(model.VOTE), uRewards, list[0].Multiplier)
	if err != nil {
		e.log.Fields(logger.Fields{"userID": req.UserID, "uRewards": uRewards}).Error(err, "[entity.ClaimQuestsRewards] entity.finalizeUserRewards() failed")
		return nil, err
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
	return &response.ClaimQuestsRewardsResponseData{
		Rewards: uRewards,
	}, nil
}

func (e *Entity) finalizeUserRewards(userID string, action string, uRewards []model.QuestUserReward, multiplier float64) ([]model.QuestUserReward, error) {
	for i, ur := range uRewards {
		rewardAmount := float64(ur.RewardAmount)
		rewardAmount *= multiplier
		uRewards[i].RewardAmount = int(math.Ceil(rewardAmount))
	}
	return uRewards, nil
}

func (e *Entity) finalizeUserQuestList(userID string, action string, list []model.QuestUserList, multiplier *float64) ([]model.QuestUserList, error) {
	// if multiplier == nil then we have to find it ...
	var streakCount int
	var m float64
	if multiplier == nil {
		switch action {
		case string(model.VOTE):
			voteStreak, err := e.repo.DiscordUserUpvoteStreak.GetByDiscordID(userID)
			if err == gorm.ErrRecordNotFound {
				m = 1
			} else if err != nil {
				e.log.Fields(logger.Fields{"userID": userID}).Error(err, "[entity.finalizeUserQuestList] repo.DiscordUserUpvoteStreak.GetByDiscordID() failed")
				return nil, err
			}
			startOfYesterday := util.StartOfDay(time.Now().UTC().AddDate(0, 0, -1))
			streakCount = voteStreak.StreakCount
			// if user didn't vote yesterday, streak is broken
			if util.StartOfDay(voteStreak.LastStreakDate).Before(startOfYesterday) {
				// then assign streakCount = 0 will guarantee no streak quest is found => multiplier = 1
				streakCount = 0
			}
		default:
			m = 1
		}
		listQ := queststreak.ListQuery{StreakCount: streakCount, Action: action, Sort: "streak_from DESC", Limit: 1}
		streakQuests, err := e.repo.QuestStreak.List(listQ)
		if err != nil {
			e.log.Fields(logger.Fields{"listQ": listQ}).Error(err, "[entity.finalizeUserQuestList] repo.QuestStreak.List() failed")
			return nil, err
		}
		if len(streakQuests) == 0 {
			m = 1
		}
		if m == 0 {
			m = streakQuests[0].Multiplier
		}
	} else {
		m = *multiplier
	}

	// ... else we can use it
	for i, item := range list {
		if item.Quest == nil || item.Quest.Rewards == nil {
			continue
		}
		for j, r := range item.Quest.Rewards {
			rewardAmount := float64(r.RewardAmount)
			rewardAmount *= m
			item.Quest.Rewards[j].RewardAmount = int(math.Ceil(rewardAmount))
		}
		list[i].Multiplier = m
	}
	return list, nil
}

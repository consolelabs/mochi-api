package entities

import (
	"encoding/json"
	"github.com/defipod/mochi/pkg/kafka/message"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"time"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/model/errors"
	"gorm.io/gorm"
)

func (e *Entity) GetUserFriendTechKeyWatchlist(profileID string) ([]model.FriendTechKeyWatchlistItem, error) {
	// 1. get list tracking keys from db
	trackingKeys, err := e.repo.FriendTechKeyWatchlistItem.ListByProfileId(profileID)
	if err != nil {
		e.log.Error(err, "[entity.GetUserFriendTechKeyWatchlist] repo.FriendTechKeyWatchlistItem.ListByProfileId failed")
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrRecordNotFound
		}

		return nil, err
	}

	return trackingKeys, nil
}

func (e *Entity) GetFriendTechKeyWatchlist() ([]model.FriendTechKeyWatchlistItem, error) {
	trackingKeys, err := e.repo.FriendTechKeyWatchlistItem.List()
	if err != nil {
		e.log.Error(err, "[entity.GetUserFriendTechKeyWatchlist] repo.FriendTechKeyWatchlistItem.List failed")
		return nil, err
	}

	return trackingKeys, nil
}

func (e *Entity) TrackFriendTechKey(profileId, keyAddress string, increaseAlertAt, decreaseAlertAt int) (*model.FriendTechKeyWatchlistItem, error) {
	exist, err := e.repo.FriendTechKeyWatchlistItem.Exist(0, keyAddress, profileId)
	if err != nil {
		e.log.Error(err, "[entity.TrackFriendTechKey] repo.FriendTechKeyWatchlistItem.Exist failed")
		return nil, err
	}
	if exist {
		e.log.Error(err, "[entity.TrackFriendTechKey] Friend tech key already tracked")
		return nil, errors.ErrFriendTechKeyAlreadyTracked
	}

	trackItem := model.FriendTechKeyWatchlistItem{
		ProfileId:       profileId,
		KeyAddress:      keyAddress,
		IncreaseAlertAt: increaseAlertAt,
		DecreaseAlertAt: decreaseAlertAt,
	}

	newTrackItem, err := e.repo.FriendTechKeyWatchlistItem.Create(trackItem)
	if err != nil {
		e.log.Error(err, "[entity.TrackFriendTechKey] repo.FriendTechKeyWatchlistItem.Create failed")
		return nil, err
	}

	return newTrackItem, nil
}

func (e *Entity) UnTrackFriendTechKey(id int) error {
	err := e.repo.FriendTechKeyWatchlistItem.Delete(id)
	if err != nil {
		e.log.Error(err, "[entity.UnTrackFriendTechKey] repo.FriendTechKeyWatchlistItem.DeleteByAddressAndProfileId failed")
		return err
	}

	return nil
}

func (e *Entity) UpdateFriendTechKeyTrack(id int, increaseAlertAt, decreaseAlertAt int) (*model.FriendTechKeyWatchlistItem, error) {
	trackItem, err := e.repo.FriendTechKeyWatchlistItem.Get(id)
	if err != nil {
		e.log.Error(err, "[entity.UpdateFriendTechKeyTrack] repo.FriendTechKeyWatchlistItem.Get failed")

		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrFriendTechKeyNotTrackedYet
		}
		return nil, err
	}

	trackItem.IncreaseAlertAt = increaseAlertAt
	trackItem.DecreaseAlertAt = decreaseAlertAt
	trackItem.UpdatedAt = time.Now()

	if err := e.repo.FriendTechKeyWatchlistItem.Update(
		*trackItem,
	); err != nil {
		e.log.Error(err, "[entity.UpdateFriendTechKeyTrack] repo.FriendTechKeyWatchlistItem.Update failed")
		return nil, err
	}

	return trackItem, nil
}

// SearchFriendTechKeys search friend scan account
func (e *Entity) SearchFriendTechKeys(req request.SearchFriendTechKeysRequest) (*response.FriendTechKeysResponse, error) {
	data, err := e.svc.FriendTech.Search(req.Query, req.Limit)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.SearchFriendTechKeys] svc.FriendTech.Search() failed")
		return nil, err
	}
	return data, nil
}

// GetFriendTechKeyPriceHistory search friend scan account
func (e *Entity) GetFriendTechKeyPriceHistory(keyAddressID, interval string) (*response.FriendTechKeyPriceHistoryResponse, error) {
	data, err := e.svc.FriendTech.GetHistory(keyAddressID, interval)
	if err != nil {
		e.log.Fields(logger.Fields{"keyAddressID": keyAddressID, "interval": interval}).Error(err, "[entity.GetFriendTechKeyPriceHistory] svc.FriendTech.GetHistory() failed")
		return nil, err
	}

	return data, nil
}

func (e *Entity) PublishKeyPriceChangeMessage(messages []message.KeyPriceChangeAlertMessage) {
	for _, msg := range messages {
		byteNotification, _ := json.Marshal(msg)
		err := e.kafka.ProduceNotification(e.cfg.Kafka.NotificationTopic, byteNotification)
		if err != nil {
			e.log.Errorf(err, "[watchKeyPriceChanges.publishMessage] - e.kafka.ProduceNotification failed")
			return
		}
	}
}

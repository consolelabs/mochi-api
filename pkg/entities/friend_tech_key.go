package entities

import (
	"encoding/json"
	"math"
	"time"

	queuetypes "github.com/consolelabs/mochi-typeset/queue/notification/typeset"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/kafka/message"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) GetUserFriendTechKeyWatchlist(profileId string) ([]model.FriendTechKeyWatchlistItem, error) {
	// 1. get list tracking keys from db
	trackingKeys, err := e.repo.FriendTechKeyWatchlistItem.List(
		model.ListFriendTechKeysFilter{
			ProfileId: profileId,
		},
	)
	if err != nil {
		e.log.Error(err, "[entity.GetUserFriendTechKeyWatchlist] repo.FriendTechKeyWatchlistItem.List failed")
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrRecordNotFound
		}

		return nil, err
	}

	return trackingKeys, nil
}

func (e *Entity) GetFriendTechKeyWatchlist() ([]model.FriendTechKeyWatchlistItem, error) {
	trackingKeys, err := e.repo.FriendTechKeyWatchlistItem.List(
		model.ListFriendTechKeysFilter{},
	)
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

func (e *Entity) GetFriendTechKeyTransactions(keyAddress string, limit int) (*response.FriendTechKeyTransactionsResponse, error) {
	data, err := e.svc.FriendTech.GetTransactions(keyAddress, limit)
	if err != nil {
		e.log.Fields(logger.Fields{"keyAddress": keyAddress, "limit": limit}).Error(err, "[entity.GetFriendTechKeyTransactions] svc.FriendTech.GetTransactions() failed")
		return nil, err
	}

	return data, nil
}

func (e *Entity) PublishKeyRelatedNotification(messages []queuetypes.NotifierMessage) {
	for _, msg := range messages {
		byteNotification, _ := json.Marshal(msg)
		err := e.kafka.ProduceNotification(e.cfg.Kafka.NotificationTopic, byteNotification)
		if err != nil {
			e.log.Errorf(err, "[watchKeyPriceChanges.publishMessage] - e.kafka.ProduceNotification failed")
			return
		}
	}
}

func (e *Entity) CalculateFriendTechKeyPriceChangePercentage(keyAddress string) (float64, error) {
	priceHistories, err := e.svc.FriendTech.GetHistory(keyAddress, "hour")
	if err != nil {
		e.log.Fields(logger.Fields{"keyAddressID": keyAddress, "interval": "hour"}).Error(err, "[entity.CalculateFriendTechKeyPriceChangePercentage] svc.FriendTech.GetHistory() failed")
		return 0, err
	}

	currentDay := time.Now().UTC().Day()
	yesterdayClosedPrice := decimal.NewFromInt(0)
	currentPrice := priceHistories.Data[len(priceHistories.Data)-1].Price

	for i := len(priceHistories.Data) - 1; i >= 0; i-- {
		if priceHistories.Data[i].Time.Day() != currentDay {
			yesterdayClosedPrice = priceHistories.Data[i].Price
			break
		}
	}

	for i := len(priceHistories.Data) - 1; i >= 0; i-- {
		if priceHistories.Data[i].Time.Day() != currentDay {
			yesterdayClosedPrice = priceHistories.Data[i].Price
			break
		}
	}

	priceChangePercentage := calculatePercentageChange(yesterdayClosedPrice, currentPrice)

	priceChangeInFloat, _ := priceChangePercentage.Float64()

	// round to nearest
	return math.Round(priceChangeInFloat*100) / 100, nil
}

func calculatePercentageChange(yesterdayPrice, todayPrice decimal.Decimal) decimal.Decimal {
	if yesterdayPrice.IsZero() {
		return decimal.Zero
	}
	return ((todayPrice.Sub(yesterdayPrice)).Div(yesterdayPrice)).Mul(decimal.NewFromInt(100))
}

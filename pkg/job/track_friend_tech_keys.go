package job

import (
	"time"

	queuetypes "github.com/consolelabs/mochi-typeset/queue/notification/typeset"
	typeset "github.com/consolelabs/mochi-typeset/typeset"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
)

const (
	jobInterval = 60 * 2 // 2 minutes
)

type trackFriendTechKeys struct {
	entity *entities.Entity
	log    logger.Logger
}

func NewTrackFriendTechKeysJob(e *entities.Entity, l logger.Logger) Job {
	return &trackFriendTechKeys{
		entity: e,
		log:    l,
	}
}

func (c *trackFriendTechKeys) Run() error {
	c.log.Info("start tracking friend tech keys")

	// fetch list tracking friend tech keys
	trackingItems, err := c.entity.GetFriendTechKeyWatchlist()
	if err != nil {
		c.log.Error(err, "[GetFriendTechKeyWatchlist] failed to get list tracking friend tech keys")
		return err
	}

	// retrieve list unique key address
	trackingKeys := make(map[string]response.FriendTechKeyTransactionsResponse)
	for _, item := range trackingItems {
		trackingKeys[item.KeyAddress] = response.FriendTechKeyTransactionsResponse{}
	}

	// fetch
	for keyAddress := range trackingKeys {
		txs, err := c.entity.GetFriendTechKeyTransactions(keyAddress, 50)
		if err != nil {
			c.log.Error(err, "[GetFriendTechKeyTransactions] failed to get transactions")
			return err
		}

		if len(txs.Data) == 0 {
			c.log.Fields(map[string]interface{}{"key_address": keyAddress}).Info("no transaction found")
			continue
		}

		trackingKeys[keyAddress] = *txs
	}

	// For each tracking item, check if it has any new transactions
	// Check if it is new transaction
	// Condition: time.Now.Unix() - interval < txs.timestamp
	// TODO: recommend that we should use queue types NotificationMessage for both producer (any service) and consumer (notification service), in this way can ignore re-define the same struct in every service that need to publish the same message
	messages := make([]queuetypes.NotifierMessage, 0)
	for _, trackingItem := range trackingItems {
		txs, ok := trackingKeys[trackingItem.KeyAddress]
		if !ok || len(txs.Data) == 0 {
			continue
		}

		// check if there is any new transaction
		for _, tx := range txs.Data {
			if time.Now().Unix()-jobInterval <= int64(tx.Timestamp) {
				messages = append(messages, queuetypes.NotifierMessage{
					Type: typeset.NOTIFICATION_KEY_NEW_TRANSACTION,
					KeyNewTransactionMetadata: &queuetypes.KeyNewTransactionMetadata{
						ProfileId:   trackingItem.ProfileId,
						Transaction: &tx,
					},
				})
			}
		}
	}

	// send notification
	if len(messages) > 0 {
		c.entity.PublishKeyRelatedNotification(messages)
	}

	return nil
}

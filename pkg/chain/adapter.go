package chain

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/logger"
)

func (s *Solana) doCacheBalance(address string) (string, error) {
	return s.cache.GetString(fmt.Sprintf("%s-%s", solanaBalanceKey, strings.ToLower(address)))
}

func (s *Solana) doNetworkBalance(address string) (float64, error) {
	balance, err := s.client.GetBalance(
		context.TODO(),
		address,
	)
	if err != nil {
		s.logger.Fields(logger.Fields{"address": address}).Error(err, "[solana.Balance] client.GetBalance() failed")
	}
	res := float64(balance) / 1e9

	// cache solana-balance-token-data
	// if error occurs -> ignore
	bytes, _ := json.Marshal(&res)
	s.logger.Infof("cache data solscan-service, key: %s", solanaBalanceKey)
	s.cache.Set(fmt.Sprintf("%s-%s", solanaBalanceKey, strings.ToLower(address)), string(bytes), 7*24*time.Hour)

	return res, err
}

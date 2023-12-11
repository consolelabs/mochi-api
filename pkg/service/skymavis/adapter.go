package skymavis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/sentrygo"
	"github.com/defipod/mochi/pkg/util"
)

var (
	sentryTags = map[string]string{
		"type": "system",
	}
)

func (s *skymavis) doCacheNft(address string) (string, error) {
	return s.cache.GetString(fmt.Sprintf("%s-%s", nftKey, strings.ToLower(address)))
}

func (s *skymavis) doNetworkNfts(address string) (*response.AxieMarketNftResponse, error) {
	q := fmt.Sprintf(`
	{
		axies(owner: "%s", from: 0, size: 10) {
			total
			results {
				id
				image
				minPrice
				name
				owner
			}
		}
		equipments(
			owner: "%s"
			from: 0
			size: 10
		) {
			total
			results {
				total
				name
				minPrice
				collections
				alias
				rarity
			}
		}
		items(owner: "%s", from: 0, size: 10) {
			results {
				tokenId
				minPrice
				figureURL
				name
				itemId
				itemAlias
				rarity
			}
			total
		}
		lands(
			from: 0
			size: 10
			owner: {ownerships: Owned, address: "%s"}
		) {
			results {
				tokenId
				minPrice
				landType
				col
				row
			}
			total
		}
	}
	`, address, address, address, address)
	q = strings.ReplaceAll(q, "\n", " ")
	q = strings.ReplaceAll(q, "\t", " ")

	req := GraphqlRequest{Query: q}
	v, err := json.Marshal(req)
	if err != nil {
		s.logger.Fields(logger.Fields{"address": address}).Error(err, "[skymavis.GetOwnedAxies] json.Marshal() failed")
		return nil, err
	}
	body := bytes.NewBuffer(v)

	res := &response.AxieMarketNftResponse{}
	status, err := util.SendRequest(util.SendRequestQuery{
		URL:      fmt.Sprintf("%s/graphql/marketplace", s.cfg.SkyMavisApiBaseUrl),
		Method:   "POST",
		Headers:  map[string]string{"Content-Type": "application/json", "X-API-Key": s.cfg.SkyMavisApiKey},
		Body:     body,
		Response: res,
	})
	if err != nil {
		s.logger.Fields(logger.Fields{"status": status}).Error(err, "[skymavis.GetOwnedAxies] util.SendRequest() failed")
		s.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Skymavis - GetOwnedAxies failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"address": address,
			},
		})
		return nil, err
	}
	if status != 200 {
		s.logger.Fields(logger.Fields{"status": status}).Error(err, "[skymavis.GetOwnedAxies] failed to query")
		s.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Skymavis - GetOwnedAxies failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"address": address,
			},
		})
		return nil, err
	}

	// if error occurs -> ignore
	bytes, _ := json.Marshal(&res)
	s.logger.Infof("cache data skymavis-service, key: %s", nftKey)
	s.cache.Set(nftKey+"-"+strings.ToLower(address), string(bytes), 7*24*time.Hour)

	return res, nil
}

func (s *skymavis) doCacheFarming(address string) (string, error) {
	return s.cache.GetString(fmt.Sprintf("%s-%s", farmingKey, strings.ToLower(address)))
}

func (s *skymavis) doNetworkFarming(address string) (*response.WalletFarmingResponse, error) {
	q := fmt.Sprintf(`
	{
		liquidityPositions(where: {user: "%s"}) {
			id
			liquidityTokenBalance
			pair {
				id
				totalSupply
				reserveUSD
				token0Price
				token1Price
				token0 {
					id
					name
					symbol
					tokenDayData(orderBy: date, orderDirection: desc, first: 1) {
						priceUSD
					}
				}
				token1 {
					id
					name
					symbol
					tokenDayData(orderBy: date, orderDirection: desc, first: 1) {
						priceUSD
					}
				}
			}
		}
	}
	`, address)
	q = strings.ReplaceAll(q, "\n", " ")
	q = strings.ReplaceAll(q, "\t", " ")

	req := GraphqlRequest{Query: q}
	v, err := json.Marshal(req)
	if err != nil {
		s.logger.Fields(logger.Fields{"address": address}).Error(err, "[skymavis.GetAddressFarming] json.Marshal() failed")
		return nil, err
	}
	body := bytes.NewBuffer(v)

	res := &response.WalletFarmingResponse{}
	status, err := util.SendRequest(util.SendRequestQuery{
		URL:      fmt.Sprintf("%s/graphql/katana", s.cfg.SkyMavisApiBaseUrl),
		Method:   "POST",
		Headers:  map[string]string{"Content-Type": "application/json", "X-API-Key": s.cfg.SkyMavisApiKey},
		Body:     body,
		Response: res,
	})
	if err != nil {
		s.logger.Fields(logger.Fields{"status": status}).Error(err, "[skymavis.GetAddressFarming] util.SendRequest() failed")
		s.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Skymavis - GetAddressFarming failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"address": address,
			},
		})
		return nil, err
	}
	if status != 200 {
		if status == 429 {
			s.logger.Fields(logger.Fields{"address": address, "status": status}).Error(err, "[skymavis.GetAddressFarming] reach Skymavis API rate limit, retrying")
			time.Sleep(3 * time.Second)
			return s.doNetworkFarming(address)
		} else {
			s.logger.Fields(logger.Fields{"status": status}).Error(err, "[skymavis.GetAddressFarming] failed to query")
			s.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
				Message: fmt.Sprintf("[API mochi] - Skymavis - GetAddressFarming failed %v", err),
				Tags:    sentryTags,
				Extra: map[string]interface{}{
					"address": address,
				},
			})
		}
		return nil, err
	}

	// if error occurs -> ignore
	bytes, _ := json.Marshal(&res)
	s.logger.Infof("cache data skymavis-service, key: %s", farmingKey)
	s.cache.Set(farmingKey+"-"+strings.ToLower(address), string(bytes), 7*24*time.Hour)

	return res, nil
}

func (s *skymavis) doCacheInternalTxns(hash string) (string, error) {
	return s.cache.GetString(fmt.Sprintf("%s-%s", internalTxsKey, strings.ToLower(hash)))
}

func (s *skymavis) doNetworkInternalTxs(hash string) (*response.SkymavisTransactionsResponse, error) {
	res := &response.SkymavisTransactionsResponse{}
	status, err := util.SendRequest(util.SendRequestQuery{
		URL:      fmt.Sprintf("%s/explorer/tx/%s/internal?from=0", s.cfg.SkyMavisApiBaseUrl, hash),
		Method:   "GET",
		Headers:  map[string]string{"Content-Type": "application/json", "X-API-Key": s.cfg.SkyMavisApiKey},
		Response: res,
	})
	if err != nil {
		s.logger.Fields(logger.Fields{"status": status}).Error(err, "[skymavis.doNetworkInternalTxs] util.SendRequest() failed")
		s.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Skymavis - doNetWorkInternalTxs failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"hash": hash,
			},
		})
		return nil, err
	}
	if status != 200 {
		s.logger.Fields(logger.Fields{"status": status}).Error(err, "[skymavis.doNetworkInternalTxs] failed to get internal txs")
		return nil, err
	}

	// if error occurs -> ignore
	bytes, _ := json.Marshal(&res)
	s.logger.Infof("cache data skymavis-service, key: %s", internalTxsKey)
	s.cache.Set(internalTxsKey+"-"+strings.ToLower(hash), string(bytes), 7*24*time.Hour)

	return res, nil
}

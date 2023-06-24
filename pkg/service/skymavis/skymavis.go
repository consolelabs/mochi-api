package skymavis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type skymavis struct {
	cfg    *config.Config
	logger logger.Logger
}

func New(cfg *config.Config) Service {
	return &skymavis{
		cfg:    cfg,
		logger: logger.NewLogrusLogger(),
	}
}

func (s *skymavis) GetAddressFarming(address string) (*response.WalletFarmingResponse, error) {
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
		URL:       fmt.Sprintf("%s/graphql/katana", s.cfg.SkyMavisApiBaseUrl),
		Method:    "POST",
		Headers:   map[string]string{"Content-Type": "application/json", "X-API-Key": s.cfg.SkyMavisApiKey},
		Body:      body,
		ParseForm: res,
	})
	if err != nil {
		s.logger.Fields(logger.Fields{"status": status}).Error(err, "[skymavis.GetAddressFarming] util.SendRequest() failed")
		return nil, err
	}
	if status != 200 {
		s.logger.Fields(logger.Fields{"status": status}).Error(err, "[skymavis.GetAddressFarming] failed to query")
		return nil, err
	}

	return res, nil
}

func (s *skymavis) GetOwnedNfts(address string) (*response.NftResponse, error) {
	q := fmt.Sprintf(`
	{
		axies(owner: "%s", size: 10, from: 0) {
			total
			results {
				id
				image
				level
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
		items(owner: "%s" from: 0, size: 10) {
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

	res := &response.NftResponse{}
	status, err := util.SendRequest(util.SendRequestQuery{
		URL:       fmt.Sprintf("%s/graphql/marketplace", s.cfg.SkyMavisApiBaseUrl),
		Method:    "POST",
		Headers:   map[string]string{"Content-Type": "application/json", "X-API-Key": s.cfg.SkyMavisApiKey},
		Body:      body,
		ParseForm: res,
	})
	if err != nil {
		s.logger.Fields(logger.Fields{"status": status}).Error(err, "[skymavis.GetOwnedAxies] util.SendRequest() failed")
		return nil, err
	}
	if status != 200 {
		s.logger.Fields(logger.Fields{"status": status}).Error(err, "[skymavis.GetOwnedAxies] failed to query")
		return nil, err
	}

	return res, nil
}

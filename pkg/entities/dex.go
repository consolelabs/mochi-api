package entities

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/dexscreener"
	"github.com/defipod/mochi/pkg/service/geckoterminal"
	"github.com/gammazero/workerpool"
)

type dexSearch struct {
	geckoterminal *geckoterminal.SearchPoolElement
	dexscreener   *dexscreener.Pair
	// TODO: ethplorer
}

func (e *Entity) SearchDexPair(req request.SearchDexPairRequest) (*response.SearchDexPairResponse, error) {
	pairs := []response.DexPair{}
	pairsMu := sync.Mutex{}
	query := req.Query

	// search geckoterminal
	geckoterminalSearch, err := e.svc.GeckoTerminal.Search(query)
	if err != nil {
		return nil, fmt.Errorf("search geckoterminal failed: %w", err)
	}

	// search dexscreener
	dexscreenerSearch, err := e.svc.DexScreener.Search(query)
	if err != nil {
		return nil, fmt.Errorf("search dexscreener failed: %w", err)
	}

	// merge by pair address
	dexSearchs := []dexSearch{}

	for i := range geckoterminalSearch.Data.Attributes.Pools {
		dexSearchs = append(dexSearchs, dexSearch{
			geckoterminal: &geckoterminalSearch.Data.Attributes.Pools[i],
		})
	}

	for i := range dexSearchs {
		for ii := range dexscreenerSearch {
			addr1 := strings.ToLower(dexSearchs[i].geckoterminal.Address)
			addr2 := strings.ToLower(dexscreenerSearch[ii].PairAddress)
			if addr1 == addr2 {
				dexSearchs[i].dexscreener = &dexscreenerSearch[ii]
			}
		}
	}

	wp := workerpool.New(10)

	for i := range dexSearchs {
		i := i
		wp.Submit(func() {
			s := dexSearchs[i]
			if s.geckoterminal == nil {
				return
			}

			if len(s.geckoterminal.Tokens) != 2 {
				return
			}

			marketCap, err := strconv.ParseFloat(s.geckoterminal.ReserveInUsd, 64)
			if err != nil {
				marketCap = 0
			}

			p := response.DexPair{
				Id:      fmt.Sprintf("%s-%s", s.geckoterminal.Network.Identifier, s.geckoterminal.Address),
				Address: s.geckoterminal.Address,
				ChainId: s.geckoterminal.Network.Identifier,
				DexId:   s.geckoterminal.Dex.Identifier,
				Url: map[string]string{
					"geckoterminal": fmt.Sprintf("https://www.geckoterminal.com/%s/pools/%s", s.geckoterminal.Network.Identifier, s.geckoterminal.Address),
				},
				MarketCapUsd: marketCap,
			}

			for _, token := range s.geckoterminal.Tokens {
				if token.IsBaseToken {
					p.BaseToken = response.DexToken{
						Name:   token.Name,
						Symbol: token.Symbol,
					}

					continue
				}
				p.QuoteToken = response.DexToken{
					Name:   token.Name,
					Symbol: token.Symbol,
				}
			}

			p.Name = fmt.Sprintf("%s/%s", p.BaseToken.Symbol, p.QuoteToken.Symbol)

			// use dexscreener as secondary source
			if s.dexscreener != nil {
				p.CreatedAt = s.dexscreener.PairCreatedAt
				p.Url["dexscreener"] = s.dexscreener.URL
				p.Txn24hBuy = int(s.dexscreener.Txns.H24.Buys)
				p.Txn24hSell = int(s.dexscreener.Txns.H24.Sells)
				p.VolumeUsd24h = s.dexscreener.Volume.H24
				p.LiquidityUsd = s.dexscreener.Liquidity.Usd
				p.Fdv = s.dexscreener.Fdv

				p.BaseToken.Address = s.dexscreener.BaseToken.Address
				p.QuoteToken.Address = s.dexscreener.QuoteToken.Address

				price, err := strconv.ParseFloat(s.dexscreener.PriceNative, 64)
				if err == nil {
					p.Price = price
				}

				priceUsd, err := strconv.ParseFloat(s.dexscreener.PriceUsd, 64)
				if err == nil {
					p.PriceUsd = priceUsd
				}

				p.PricePercentChange24H = s.dexscreener.PriceChange.H24
			}

			// if chain is eth, get info from ethplorer
			if p.IsEth() {
				holders, err := e.svc.Ethplorer.GetTopTokenHolders(p.BaseToken.Address, 5)
				if err != nil {
					e.log.Error(err, "get top token holders failed")
				}

				if holders != nil {
					for _, holder := range holders.Holders {
						p.Holders = append(p.Holders, response.DexTokenHolder{
							Address: holder.Address,
							Balance: holder.Balance,
							Percent: holder.Share,
						})
					}
				}

				info, err := e.svc.Ethplorer.GetTokenInfo(p.BaseToken.Address)
				if err != nil {
					e.log.Error(err, "get token info failed")
				}

				if info != nil {
					p.Owner = info.Owner
				}
			}
			pairsMu.Lock()
			defer pairsMu.Unlock()
			pairs = append(pairs, p)
		})
	}

	wp.StopWait()

	// sort by liquidity
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].LiquidityUsd > pairs[j].LiquidityUsd
	})

	return &response.SearchDexPairResponse{
		Pairs: pairs,
	}, nil
}

package coingecko

import (
	"reflect"
	"testing"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/response"
)

func TestCoinGecko_GetCoinBRC20(t *testing.T) {
	type fields struct {
		getMarketChartURL                 string
		searchCoinURL                     string
		getCoinURL                        string
		getPriceURL                       string
		getCoinOhlc                       string
		getCoinsMarketData                string
		getSupportedCoins                 string
		getAssetPlatforms                 string
		getCoinByContract                 string
		getTrendingSearch                 string
		getTopGainerLoser                 string
		getHistoricalGlobalMarketChartURL string
		brc20Cache                        cache.Cache
		brc20KeyPrefix                    string
	}
	type args struct {
		coinName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.GetCoinResponse
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CoinGecko{
				getMarketChartURL:                 tt.fields.getMarketChartURL,
				searchCoinURL:                     tt.fields.searchCoinURL,
				getCoinURL:                        tt.fields.getCoinURL,
				getPriceURL:                       tt.fields.getPriceURL,
				getCoinOhlc:                       tt.fields.getCoinOhlc,
				getCoinsMarketData:                tt.fields.getCoinsMarketData,
				getSupportedCoins:                 tt.fields.getSupportedCoins,
				getAssetPlatforms:                 tt.fields.getAssetPlatforms,
				getCoinByContract:                 tt.fields.getCoinByContract,
				getTrendingSearch:                 tt.fields.getTrendingSearch,
				getTopGainerLoser:                 tt.fields.getTopGainerLoser,
				getHistoricalGlobalMarketChartURL: tt.fields.getHistoricalGlobalMarketChartURL,
				brc20Cache:                        tt.fields.brc20Cache,
				brc20KeyPrefix:                    tt.fields.brc20KeyPrefix,
			}
			got, err, got1 := c.GetCoinBRC20(tt.args.coinName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CoinGecko.GetCoinBRC20() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CoinGecko.GetCoinBRC20() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CoinGecko.GetCoinBRC20() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

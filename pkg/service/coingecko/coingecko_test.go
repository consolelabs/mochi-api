package coingecko

import (
	"reflect"
	"testing"

	"github.com/defipod/mochi/pkg/response"
)

func TestCoinGecko_GetMarketData(t *testing.T) {
	type fields struct {
		getMarketChartURL string
		getMarketDataURL  string
		searchCoinURL     string
		getCoinURL        string
		getPriceURL       string
	}
	type args struct {
		coinID   string
		currency string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.MarketDataResponse
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "get data successfully",
			fields: fields{
				getMarketDataURL: "https://api.coingecko.com/api/v3/coins/markets?vs_currency=%s&ids=%s",
			},
			args: args{
				coinID:   "bitcoin",
				currency: "usd",
			},
			want:    &response.MarketDataResponse{},
			want1:   200,
			wantErr: false,
		},
		{
			name: "coin id invalid",
			fields: fields{
				getMarketDataURL: "https://api.coingecko.com/api/v3/coins/markets?vs_currency=%s&ids=%s",
			},
			args: args{
				coinID:   "abcdcoin",
				currency: "usd",
			},
			want1:   400,
			wantErr: true,
		},
		{
			name: "currency invalid",
			fields: fields{
				getMarketDataURL: "https://api.coingecko.com/api/v3/coins/markets?vs_currency=%s&ids=%s",
			},
			args: args{
				coinID:   "bitcoin",
				currency: "usdabc",
			},
			want1:   400,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CoinGecko{
				getMarketChartURL: tt.fields.getMarketChartURL,
				getMarketDataURL:  tt.fields.getMarketDataURL,
				searchCoinURL:     tt.fields.searchCoinURL,
				getCoinURL:        tt.fields.getCoinURL,
				getPriceURL:       tt.fields.getPriceURL,
			}
			got, err, got1 := c.GetMarketData(tt.args.coinID, tt.args.currency)
			if (err != nil) != tt.wantErr {
				t.Errorf("CoinGecko.GetMarketData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// returned value is not MarketDataResponse struct OR returned struct is empty
			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) || got == (&response.MarketDataResponse{}) {
				t.Errorf("CoinGecko.GetMarketData() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CoinGecko.GetMarketData() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

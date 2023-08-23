package dexscreener

import (
	"testing"
)

func Test_dexscreener_Get(t *testing.T) {
	type args struct {
		network string
		address string
	}
	tests := []struct {
		name    string
		d       *dexscreener
		args    args
		want    *Pair
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			d:    &dexscreener{},
			args: args{
				network: "ethereum",
				address: "0x5201196516E375FAbe6A22Ed7eD61b3DA3DDbe4d",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dexscreener{}
			got, err := d.Get(tt.args.network, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("dexscreener.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Logf("got: %+v", got)
		})
	}
}

func Test_dexscreener_Search(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		d       *dexscreener
		args    args
		want    []Pair
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			d:    &dexscreener{},
			args: args{
				query: "0x5201196516E375FAbe6A22Ed7eD61b3DA3DDbe4d",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dexscreener{}
			got, err := d.Search(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("dexscreener.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Logf("got: %+v", got)
		})
	}
}

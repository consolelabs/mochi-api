package entities

import (
	"reflect"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/discordwallet"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service"
	"github.com/defipod/mochi/pkg/service/indexer"
	mock_indexer "github.com/defipod/mochi/pkg/service/indexer/mocks"
	"github.com/golang/mock/gomock"
)

func TestEntity_GetNftSales(t *testing.T) {
	type fields struct {
		repo     *repo.Repo
		store    repo.Store
		log      logger.Logger
		dcwallet discordwallet.IDiscordWallet
		discord  *discordgo.Session
		cache    cache.Cache
		svc      *service.Service
		cfg      config.Config
		indexer  indexer.Service
	}
	type args struct {
		addr     string
		platform string
	}
	ctrl := gomock.NewController(t)
	mockIndexer := mock_indexer.NewMockService(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.NftSales
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "get all successfully",
			fields: fields{
				indexer: mockIndexer,
			},
			args: args{
				addr:     "0xb54FF1EBc9950fce19Ee9E055A382B1219f862f0",
				platform: "paintswap",
			},
			want: &response.NftSales{
				Platform:             "paintswap",
				NftName:              "Cyber Neko 4",
				NftStatus:            "sold",
				NftCollectionAddress: "0xb54FF1EBc9950fce19Ee9E055A382B1219f862f0",
				NftPrice:             12.12,
				NftPriceToken:        "eth",
				Buyer:                "0x9f1420cd1a1bbef2240de9d8a005ec2dba9c58c5",
				Seller:               "0x9dce416892c8a38c187016c16355443ccae3aae4",
			},
			wantErr: false,
		},
		{
			name: "collection not exist",
			fields: fields{
				indexer: mockIndexer,
			},
			args: args{
				addr:     "0xb54FF1EBc9950fce19Ee9E055A382B1abc",
				platform: "paintswap",
			},
			want:    nil,
			wantErr: true,
		},
	}
	indexerResponse := &response.NftSalesResponse{
		Data: []response.NftSales{
			{
				Platform:             "paintswap",
				NftName:              "Light",
				NftStatus:            "sold",
				NftCollectionAddress: "0xb54FF1EBc9950fce19Ee9E055A382B1219f862f0",
				NftPrice:             12.12,
				NftPriceToken:        "eth",
				Buyer:                "0x9f1420cd1a1bbef2240de9d8a005ec2dba9c58c5",
				Seller:               "0x9dce416892c8a38c187016c16355443ccae3aae4",
			},
			{
				Platform:             "paintswap",
				NftName:              "Cyber Neko 3",
				NftStatus:            "sold",
				NftCollectionAddress: "0xb54FF1EBc9950fce19Ee9E055A382B1219f862f0",
				NftPrice:             12.12,
				NftPriceToken:        "eth",
				Buyer:                "0x9f1420cd1a1bbef2240de9d8a005ec2dba9c58c5",
				Seller:               "0x9dce416892c8a38c187016c16355443ccae3aae4",
			},
			{
				Platform:             "paintswap",
				NftName:              "Cyber Neko 4",
				NftStatus:            "sold",
				NftCollectionAddress: "0xb54FF1EBc9950fce19Ee9E055A382B1219f862f0",
				NftPrice:             12.12,
				NftPriceToken:        "eth",
				Buyer:                "0x9f1420cd1a1bbef2240de9d8a005ec2dba9c58c5",
				Seller:               "0x9dce416892c8a38c187016c16355443ccae3aae4",
			},
		},
	}
	mockIndexer.EXPECT().GetNftSales().Return(indexerResponse, nil).AnyTimes()
	mockIndexer.EXPECT().GetNftSales().Return(indexerResponse, nil).AnyTimes()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				repo:     tt.fields.repo,
				store:    tt.fields.store,
				log:      tt.fields.log,
				dcwallet: tt.fields.dcwallet,
				discord:  tt.fields.discord,
				cache:    tt.fields.cache,
				svc:      tt.fields.svc,
				cfg:      tt.fields.cfg,
				indexer:  tt.fields.indexer,
			}
			got, err := e.GetNftSales(tt.args.addr, tt.args.platform)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetNftSales() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetNftSales() = %v, want %v", got, tt.want)
			}
		})
	}
}
